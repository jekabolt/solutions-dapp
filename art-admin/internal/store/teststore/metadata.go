package teststore

import (
	"context"
	"fmt"
	"time"

	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
)

func (ts *testStore) GetOffchainMetadata(ctx context.Context) (*pb_metadata.Meta, error) {
	if ts.offchainMetadata.MetaInfo == nil {
		return nil, nil
	}
	return ts.offchainMetadata, nil
}

func (ts *testStore) UpdateOffchainMetadataAttributes(ctx context.Context, mintSeq int32, attr []*pb_metadata.Attributes) (*pb_metadata.Meta, error) {
	if mintSeq >= int32(len(ts.offchainMetadata.Metadata)) {
		return nil, fmt.Errorf("invalid mint seq")
	}
	for k, md := range ts.offchainMetadata.Metadata {
		if md.MintSequenceNumber == mintSeq {
			md.Attributes = attr
			ts.offchainMetadata.Metadata[k] = md
		}
	}
	return ts.offchainMetadata, nil
}

func (ts *testStore) UpdateOffchainMetadataImage(ctx context.Context, mintSeq int32, imageUrl string) (*pb_metadata.Meta, error) {
	if mintSeq >= int32(len(ts.offchainMetadata.Metadata)) {
		return nil, fmt.Errorf("invalid mint seq")
	}
	for k, md := range ts.offchainMetadata.Metadata {
		if md.MintSequenceNumber == mintSeq {
			md.Image = imageUrl
			ts.offchainMetadata.Metadata[k] = md
		}
	}
	return ts.offchainMetadata, nil
}

func (ts *testStore) AddMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (*pb_metadata.MetaInfo, error) {
	id := getId()
	meta := &pb_metadata.Meta{
		Metadata: md,
		MetaInfo: &pb_metadata.MetaInfo{
			Id:         id,
			Processing: true,
			IpfsUrl:    "",
			Ts:         time.Now().Unix(),
		},
	}
	ts.metadata[id] = meta
	return meta.MetaInfo, nil
}

func (ts *testStore) SetIPFSUrl(ctx context.Context, key string, IPFSUrl string) error {

	if ts.offchainMetadata.MetaInfo.Id == key {
		ts.offchainMetadata.MetaInfo.IpfsUrl = IPFSUrl
		return nil
	}

	md, ok := ts.metadata[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	md.MetaInfo.Processing = false
	md.MetaInfo.IpfsUrl = IPFSUrl
	ts.metadata[key] = md
	return nil
}

func (ts *testStore) SetProcessing(ctx context.Context, key string, processing bool) error {

	if ts.offchainMetadata.MetaInfo.Id == key {
		ts.offchainMetadata.MetaInfo.Processing = processing
		return nil
	}

	md, ok := ts.metadata[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	md.MetaInfo.Processing = processing
	ts.metadata[key] = md

	return nil
}

func (ts *testStore) SetOffchain(ctx context.Context, id string) error {
	if ts.offchainMetadata.MetaInfo != nil {
		return fmt.Errorf("offchain metadata already exists")
	}

	md, ok := ts.metadata[id]
	if !ok {
		return fmt.Errorf("no such key")
	}
	ts.offchainMetadata = md
	delete(ts.metadata, id)
	return nil
}

func (ts *testStore) GetAllMetadata(ctx context.Context) ([]*pb_metadata.Meta, error) {
	md := []*pb_metadata.Meta{}
	for _, v := range ts.metadata {
		md = append(md, v)
	}
	md = append(md, ts.offchainMetadata)
	return md, nil
}

func (ts *testStore) GetMetadataById(ctx context.Context, id string) (*pb_metadata.Meta, error) {
	md, ok := ts.metadata[id]
	if !ok {
		if ts.offchainMetadata.MetaInfo.Id == id {
			return ts.offchainMetadata, nil
		}
		return nil, fmt.Errorf("no such key")
	}
	return md, nil
}

func (ts *testStore) DeleteById(ctx context.Context, id string) error {
	delete(ts.metadata, id)
	delete(ts.metadata, id)
	return nil
}
