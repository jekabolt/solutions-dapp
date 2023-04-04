package metadata

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

func newMetadataAttributes(
	userDescription, collection,
	author, timeSpent string) []*pb_metadata.Attributes {

	a := []*pb_metadata.Attributes{}
	if userDescription != "" {
		a = append(a, &pb_metadata.Attributes{
			TraitType: "Description",
			Value:     userDescription,
		})
	}

	if collection != "" {
		a = append(a, &pb_metadata.Attributes{
			TraitType: "Collection",
			Value:     collection,
		})
	}
	if author != "" {
		a = append(a, &pb_metadata.Attributes{
			TraitType: "Author",
			Value:     author,
		})
	}
	if timeSpent != "" {
		a = append(a, &pb_metadata.Attributes{
			TraitType: "Time Spent",
			Value:     timeSpent,
		})
	}
	return a
}

func marshalMetadata(mds []*pb_metadata.MetadataUnit) ([]byte, error) {
	all, err := json.Marshal(mds)
	if err != nil {
		return nil, fmt.Errorf("marshalMetadata:json.Marshal: [%s]", err.Error())
	}
	allMeta := base64.StdEncoding.EncodeToString(all)

	uf := []ipfs.UploadFolder{
		{
			Path:    "_metadata.json",
			Content: allMeta,
		},
	}

	for i, md := range mds {
		mdB64, err := toB64(md)
		if err != nil {
			return nil, fmt.Errorf("marshalMetadata:mr.toB64: [%s]", err.Error())
		}
		uf = append(uf, ipfs.UploadFolder{
			Path:    fmt.Sprintf("%d.%s", i+1, "json"),
			Content: mdB64,
		})
	}
	return json.Marshal(uf)
}

func toB64(m *pb_metadata.MetadataUnit) (string, error) {
	all, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("marshalMetadata:json.Marshal: [%s]", err.Error())
	}
	return base64.StdEncoding.EncodeToString(all), nil
}

// mergeMeta merges metadata with mint request
func mergeMeta(mds []*pb_metadata.MetadataUnit, mrs map[int32]*pb_nft.NFTMintRequestWithStatus) []*pb_metadata.MetadataUnit {
	for i, md := range mds {
		mr, ok := mrs[md.MintSequenceNumber]
		if ok && mr.GetOnchainUrl() != "" {
			mds[i].Image = mr.GetOnchainUrl()
			mds[i].Attributes = append(
				mds[i].Attributes,
				newMetadataAttributes(
					mr.NftMintRequest.Description,
					mr.Collection,
					mr.Author,
					mr.Duration,
				)...,
			)
		}
	}
	return mds
}
