package metadata

import (
	"context"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/mongo"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	"github.com/rs/zerolog/log"
)

type Uploader interface {
	UploadInitial(ctx context.Context) (*pb_metadata.MetaInfo, error)
	UploadIpfs(ctx context.Context) (*pb_metadata.MetaInfo, error)
}

type Config struct {
	BaseExternalUrl string `env:"METADATA_BASE_EXTERNAL_URL" envDefault:"https://nft.sys.solutions/nft/%d"`
	BaseImageUrl    string `env:"METADATA_BASE_IMAGE_URL" envDefault:"https://nft.sys.solutions/nft/%d/image"`
	TotalQuantity   int    `env:"METADATA_TOTAL_QUANTITY" envDefault:"100000"`
	NamePrefix      string `env:"METADATA_NAME_PREFIX" envDefault:"Solutions art #%d"`
	DefaultAuthor   string `env:"METADATA_DEFAULT_AUTHOR" envDefault:"solutions"`
	UploadRetries   int    `env:"METADATA_UPLOAD_RETRIES" envDefault:"3"`
}

func (c *Config) New(
	desc descriptions.MintDescription,
	mints mongo.GetToUpload,
	metaStore mongo.MetadataStore,
	ipfs ipfs.Uploader,
) (Uploader, error) {
	return &MetaManager{
		desc:      desc,
		mints:     mints,
		metaStore: metaStore,
		ipfs:      ipfs,
		C:         c,
	}, nil
}

type MetaManager struct {
	desc      descriptions.MintDescription
	mints     mongo.GetToUpload
	metaStore mongo.MetadataStore
	ipfs      ipfs.Uploader
	C         *Config
}

// GetInitialMetadata returns initial metadata with all offchain images and no attributes
func (mm *MetaManager) getInitialMetadata() ([]*pb_metadata.MetadataUnit, error) {
	ms := []*pb_metadata.MetadataUnit{}
	for i := 1; i <= mm.C.TotalQuantity; i++ {
		d, err := mm.desc.GetDescriptionOn(i)
		if err != nil {
			return nil, fmt.Errorf("g:mm.desc.GetDescriptionOn: [%s]", err.Error())
		}
		m := &pb_metadata.MetadataUnit{
			Name:        fmt.Sprintf(mm.C.NamePrefix, i),
			Description: d.String(),
			ExternalUrl: fmt.Sprintf(mm.C.BaseExternalUrl, i),
			Image:       fmt.Sprintf(mm.C.BaseImageUrl, i),
			Attributes:  newMetadataAttributes("", "", mm.C.DefaultAuthor, "0s"),
		}
		ms = append(ms, m)
	}
	return ms, nil
}

func (mm *MetaManager) UploadInitial(ctx context.Context) (*pb_metadata.MetaInfo, error) {

	omd, err := mm.metaStore.GetOffchainMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("UploadInitial:mm.metaStore.GetOffchainMetadata: [%s]", err.Error())
	}
	if omd != nil {
		// already uploaded can be uploaded only once
		return omd.MetaInfo, nil
	}

	// upload to ipfs and store in db
	mds, err := mm.getInitialMetadata()
	if err != nil {
		return nil, fmt.Errorf("UploadInitial:mm.getInitialMetadata: [%s]", err.Error())
	}

	mi, err := mm.metaStore.AddMetadata(ctx, mds)
	if err != nil {
		return nil, fmt.Errorf("UploadInitial:mm.metaStore.AddMetadata: [%s]", err.Error())
	}

	err = mm.metaStore.SetOffchain(ctx, mi.Id)
	if err != nil {
		return nil, fmt.Errorf("UploadInitial:mm.metaStore.SetOffchain: [%s]", err.Error())
	}

	bs, err := marshalMetadata(mds)
	if err != nil {
		return nil, fmt.Errorf("UploadInitial:marshalMetadata: [%s]", err.Error())
	}

	// upload to ipfs in background and store url in db retrying on error mm.C.UploadRetries times
	go func() {
		// mark as processing false if upload fails or succeeds
		defer func() {
			mm.metaStore.SetProcessing(ctx, mi.Id, false)
		}()
		// retry on error
		for i := 0; i < mm.C.UploadRetries; i++ {
			url, err := mm.ipfs.UploadData(bs)
			if err != nil {
				log.Ctx(ctx).Error().Err(err).Msgf("UploadInitial:mm.moralis.Upload: retrying[%v...]", i)
			}
			if err == nil {
				mm.metaStore.SetIPFSUrl(ctx, mi.Id, url)
				log.Ctx(ctx).Info().Msgf("UploadInitial:mm.moralis.Upload: success[%v]", i)
				return
			}
		}
		log.Ctx(ctx).Error().Msgf("UploadInitial:mm.moralis.Upload: failed after %v retries", mm.C.UploadRetries)
	}()

	return mi, nil
}

func (mm *MetaManager) UploadIpfs(ctx context.Context) (*pb_metadata.MetaInfo, error) {

	omd, err := mm.metaStore.GetOffchainMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("UploadIpfs:mm.metaStore.GetOffchainMetadata: [%s]", err.Error())
	}
	if omd == nil {
		return nil, fmt.Errorf("UploadIpfs:mm.metaStore.GetOffchainMetadata: make initial upload first")
	}

	toUpl, err := mm.mints.GetAllToUpload(context.Background())
	if err != nil {
		return nil, fmt.Errorf("UploadIpfs:mm.mints.GetAllToUpload: [%s]", err.Error())
	}
	if len(toUpl) == 0 {
		return nil, fmt.Errorf("UploadIpfs:mm.mints.GetAllToUpload: no mints to upload")
	}

	mergedMeta := mergeMeta(omd.Metadata, toUpl)

	mi, err := mm.metaStore.AddMetadata(ctx, mergedMeta)
	if err != nil {
		return nil, fmt.Errorf("UploadIpfs:mm.metaStore.AddMetadata: [%s]", err.Error())
	}
	bs, err := marshalMetadata(mergedMeta)
	if err != nil {
		return nil, fmt.Errorf("UploadIpfs:marshalMetadata: [%s]", err.Error())
	}

	// upload to ipfs in background and store url in db retrying on error mm.C.UploadRetries times
	go func() {
		// mark as processing false if upload fails or succeeds
		defer func() {
			mm.metaStore.SetProcessing(ctx, mi.Id, false)
		}()
		// retry on error
		for i := 0; i < mm.C.UploadRetries; i++ {
			url, err := mm.ipfs.UploadData(bs)
			if err != nil {
				log.Ctx(ctx).Error().Err(err).Msgf("UploadIpfs:mm.moralis.Upload: retrying[%v...]", i)
			}
			if err == nil {
				mm.metaStore.SetIPFSUrl(ctx, mi.Id, url)
				log.Ctx(ctx).Info().Msgf("UploadIpfs:mm.moralis.Upload: success[%v]", i)
				return
			}
		}
		log.Ctx(ctx).Error().Msgf("UploadIpfs:mm.moralis.Upload: failed after %v retries", mm.C.UploadRetries)
	}()

	return mi, nil

}
