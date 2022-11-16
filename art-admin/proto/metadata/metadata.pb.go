// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: metadata/metadata.proto

package metadata

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UploadOffchainMetadata response
type UploadOffchainMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal id in db
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UploadOffchainMetadataResponse) Reset() {
	*x = UploadOffchainMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadOffchainMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadOffchainMetadataResponse) ProtoMessage() {}

func (x *UploadOffchainMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadOffchainMetadataResponse.ProtoReflect.Descriptor instead.
func (*UploadOffchainMetadataResponse) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *UploadOffchainMetadataResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request for UploadIPFSMetadataRequest
type UploadIPFSMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db metadata id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UploadIPFSMetadataRequest) Reset() {
	*x = UploadIPFSMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadIPFSMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadIPFSMetadataRequest) ProtoMessage() {}

func (x *UploadIPFSMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadIPFSMetadataRequest.ProtoReflect.Descriptor instead.
func (*UploadIPFSMetadataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *UploadIPFSMetadataRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request for UploadIPFSMetadataRequest
type DeleteIPFSMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db metadata id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteIPFSMetadataRequest) Reset() {
	*x = DeleteIPFSMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIPFSMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIPFSMetadataRequest) ProtoMessage() {}

func (x *DeleteIPFSMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIPFSMetadataRequest.ProtoReflect.Descriptor instead.
func (*DeleteIPFSMetadataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteIPFSMetadataRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// description of metadata stored in db
type MetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpfsUrl    string `protobuf:"bytes,1,opt,name=ipfsUrl,proto3" json:"ipfsUrl,omitempty"`
	Processing bool   `protobuf:"varint,3,opt,name=processing,proto3" json:"processing,omitempty"`
	Ts         int64  `protobuf:"varint,4,opt,name=ts,proto3" json:"ts,omitempty"`
	Key        string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *MetaInfo) Reset() {
	*x = MetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaInfo) ProtoMessage() {}

func (x *MetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaInfo.ProtoReflect.Descriptor instead.
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *MetaInfo) GetIpfsUrl() string {
	if x != nil {
		return x.IpfsUrl
	}
	return ""
}

func (x *MetaInfo) GetProcessing() bool {
	if x != nil {
		return x.Processing
	}
	return false
}

func (x *MetaInfo) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *MetaInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response for GetAllMetadata
type GetAllMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// array of metadata
	MetaInfo []*MetaInfo `protobuf:"bytes,1,rep,name=metaInfo,proto3" json:"metaInfo,omitempty"`
}

func (x *GetAllMetadataResponse) Reset() {
	*x = GetAllMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMetadataResponse) ProtoMessage() {}

func (x *GetAllMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetAllMetadataResponse) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllMetadataResponse) GetMetaInfo() []*MetaInfo {
	if x != nil {
		return x.MetaInfo
	}
	return nil
}

// Single Unit of metadata _metadata.json
type MetadataUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of nft
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// description of nft
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// image of nft
	OffchainImage string `protobuf:"bytes,3,opt,name=offchainImage,proto3" json:"offchainImage,omitempty"`
	// image of nft
	OnchainImage string `protobuf:"bytes,4,opt,name=onchainImage,proto3" json:"onchainImage,omitempty"`
	// edition of nft
	Edition int32 `protobuf:"varint,5,opt,name=edition,proto3" json:"edition,omitempty"`
	// sequence number of nft
	MintSequenceNumber int32 `protobuf:"varint,6,opt,name=mintSequenceNumber,proto3" json:"mintSequenceNumber,omitempty"`
	// date of submit
	Date int32 `protobuf:"varint,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *MetadataUnit) Reset() {
	*x = MetadataUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataUnit) ProtoMessage() {}

func (x *MetadataUnit) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataUnit.ProtoReflect.Descriptor instead.
func (*MetadataUnit) Descriptor() ([]byte, []int) {
	return file_metadata_metadata_proto_rawDescGZIP(), []int{5}
}

func (x *MetadataUnit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetadataUnit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MetadataUnit) GetOffchainImage() string {
	if x != nil {
		return x.OffchainImage
	}
	return ""
}

func (x *MetadataUnit) GetOnchainImage() string {
	if x != nil {
		return x.OnchainImage
	}
	return ""
}

func (x *MetadataUnit) GetEdition() int32 {
	if x != nil {
		return x.Edition
	}
	return 0
}

func (x *MetadataUnit) GetMintSequenceNumber() int32 {
	if x != nil {
		return x.MintSequenceNumber
	}
	return 0
}

func (x *MetadataUnit) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_metadata_metadata_proto protoreflect.FileDescriptor

var file_metadata_metadata_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x1e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x2d, 0x0a, 0x19, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x50, 0x46, 0x53,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x2d, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x53, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x66, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x70, 0x66, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x70, 0x66, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x66, 0x66,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x12, 0x6d, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x74, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x32, 0xd3, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7d,
	0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a,
	0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x50, 0x46, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x50, 0x46, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x70, 0x66, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x6f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x53, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x65, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6c, 0x74, 0x2f, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x64, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x72,
	0x74, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_metadata_proto_rawDescOnce sync.Once
	file_metadata_metadata_proto_rawDescData = file_metadata_metadata_proto_rawDesc
)

func file_metadata_metadata_proto_rawDescGZIP() []byte {
	file_metadata_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_metadata_proto_rawDescData)
	})
	return file_metadata_metadata_proto_rawDescData
}

var file_metadata_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metadata_metadata_proto_goTypes = []interface{}{
	(*UploadOffchainMetadataResponse)(nil), // 0: metadata.UploadOffchainMetadataResponse
	(*UploadIPFSMetadataRequest)(nil),      // 1: metadata.UploadIPFSMetadataRequest
	(*DeleteIPFSMetadataRequest)(nil),      // 2: metadata.DeleteIPFSMetadataRequest
	(*MetaInfo)(nil),                       // 3: metadata.MetaInfo
	(*GetAllMetadataResponse)(nil),         // 4: metadata.GetAllMetadataResponse
	(*MetadataUnit)(nil),                   // 5: metadata.MetadataUnit
	(*emptypb.Empty)(nil),                  // 6: google.protobuf.Empty
}
var file_metadata_metadata_proto_depIdxs = []int32{
	3, // 0: metadata.GetAllMetadataResponse.metaInfo:type_name -> metadata.MetaInfo
	6, // 1: metadata.Metadata.UploadOffchainMetadata:input_type -> google.protobuf.Empty
	1, // 2: metadata.Metadata.UploadIPFSMetadata:input_type -> metadata.UploadIPFSMetadataRequest
	2, // 3: metadata.Metadata.DeleteIPFSMetadata:input_type -> metadata.DeleteIPFSMetadataRequest
	6, // 4: metadata.Metadata.GetAllMetadata:input_type -> google.protobuf.Empty
	0, // 5: metadata.Metadata.UploadOffchainMetadata:output_type -> metadata.UploadOffchainMetadataResponse
	6, // 6: metadata.Metadata.UploadIPFSMetadata:output_type -> google.protobuf.Empty
	6, // 7: metadata.Metadata.DeleteIPFSMetadata:output_type -> google.protobuf.Empty
	4, // 8: metadata.Metadata.GetAllMetadata:output_type -> metadata.GetAllMetadataResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metadata_metadata_proto_init() }
func file_metadata_metadata_proto_init() {
	if File_metadata_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadOffchainMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metadata_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadIPFSMetadataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metadata_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIPFSMetadataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metadata_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metadata_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metadata_metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataUnit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metadata_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_metadata_proto_depIdxs,
		MessageInfos:      file_metadata_metadata_proto_msgTypes,
	}.Build()
	File_metadata_metadata_proto = out.File
	file_metadata_metadata_proto_rawDesc = nil
	file_metadata_metadata_proto_goTypes = nil
	file_metadata_metadata_proto_depIdxs = nil
}
