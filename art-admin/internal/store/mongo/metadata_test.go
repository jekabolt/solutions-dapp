package mongo

import (
	"context"
	"testing"

	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	"github.com/matryer/is"
)

func getMetadataUnits() []*pb_metadata.MetadataUnit {
	return []*pb_metadata.MetadataUnit{
		{
			Name:               "MetadataUnit 1 ",
			Description:        "description 1 ",
			MintSequenceNumber: 1,
			ExternalUrl:        "external-url",
			Image:              "image",
			Attributes: []*pb_metadata.Attributes{
				{
					TraitType:   "trait-type",
					Value:       "value",
					DisplayType: "display-type",
				},
			},
		},
		{
			Name:               "MetadataUnit 2 ",
			Description:        "description 2 ",
			MintSequenceNumber: 2,
			ExternalUrl:        "external-url 2",
			Image:              "image 2",
			Attributes: []*pb_metadata.Attributes{
				{
					TraitType:   "trait-type",
					Value:       "value",
					DisplayType: "display-type",
				},
			},
		},
	}
}

func TestMetadata(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()

	c := Config{
		DSN: getMongoDSN(),
		DB:  "test",
	}
	mdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := mdb.MetadataStore()
	is.NoErr(err)

	ids := []string{}
	defer func() {
		err := mdb.metadata.Drop(ctx)
		is.NoErr(err)
	}()

	for i := 0; i < 100; i++ {
		mu, err := ms.AddMetadata(ctx, getMetadataUnits())
		is.NoErr(err)
		ids = append(ids, mu.Id)
	}

	all, err := ms.GetAllMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(all), 100)

	err = ms.SetProcessing(ctx, all[0].MetaInfo.Id, true)
	is.NoErr(err)

	md, err := ms.GetMetadataById(ctx, all[0].MetaInfo.Id)
	is.NoErr(err)
	is.Equal(md.MetaInfo.Processing, true)

	err = ms.SetIPFSUrl(ctx, all[0].MetaInfo.Id, "ipfs-url")
	is.NoErr(err)

	md, err = ms.GetMetadataById(ctx, all[0].MetaInfo.Id)
	is.NoErr(err)
	is.Equal(md.MetaInfo.IpfsUrl, "ipfs-url")

	err = ms.DeleteById(ctx, all[0].MetaInfo.Id)
	is.NoErr(err)

	md, err = ms.GetMetadataById(ctx, all[0].MetaInfo.Id)
	is.True(err != nil)

	all, err = ms.GetAllMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(all), 99)

	// test offchain
	emptyMeta, err := ms.GetOffchainMetadata(ctx)
	is.NoErr(err)
	is.Equal(emptyMeta, nil)

	err = ms.SetOffchain(ctx, all[2].MetaInfo.Id)
	is.NoErr(err)

	omd, err := ms.GetOffchainMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(omd.Metadata), 2)

	newAttr := []*pb_metadata.Attributes{
		{
			TraitType:   "test",
			Value:       "test",
			DisplayType: "test",
		}, {
			TraitType:   "test",
			Value:       "test",
			DisplayType: "test",
		},
	}
	testImg := "test-img"

	_, err = ms.UpdateOffchainMetadataAttributes(ctx, omd.Metadata[0].MintSequenceNumber, newAttr)
	is.NoErr(err)

	_, err = ms.UpdateOffchainMetadataImage(ctx, omd.Metadata[0].MintSequenceNumber, testImg)
	is.NoErr(err)

	updOmd, err := ms.GetOffchainMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(omd.Metadata), 2)
	for _, md := range updOmd.Metadata {
		if md.MintSequenceNumber == omd.Metadata[0].MintSequenceNumber {
			is.Equal(len(md.Attributes), 2)
			is.Equal(md.Image, testImg)
		}
	}

}
