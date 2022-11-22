// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: collection/collection.proto

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

// Request for CreateNewCollection
type CreateNewCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// collcection name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// collection capacity
	Capacity int32 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *CreateNewCollectionRequest) Reset() {
	*x = CreateNewCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewCollectionRequest) ProtoMessage() {}

func (x *CreateNewCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewCollectionRequest.ProtoReflect.Descriptor instead.
func (*CreateNewCollectionRequest) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNewCollectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNewCollectionRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// Response for CreateNewCollection
type CreateNewCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateNewCollectionResponse) Reset() {
	*x = CreateNewCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewCollectionResponse) ProtoMessage() {}

func (x *CreateNewCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewCollectionResponse.ProtoReflect.Descriptor instead.
func (*CreateNewCollectionResponse) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNewCollectionResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request for DeleteCollection
type DeleteCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteCollectionRequest) Reset() {
	*x = DeleteCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCollectionRequest) ProtoMessage() {}

func (x *DeleteCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCollectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteCollectionRequest) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCollectionRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response for DeleteCollection
type DeleteCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteCollectionResponse) Reset() {
	*x = DeleteCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCollectionResponse) ProtoMessage() {}

func (x *DeleteCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCollectionResponse.ProtoReflect.Descriptor instead.
func (*DeleteCollectionResponse) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCollectionResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request for UpdateCollectionCapacity
type UpdateCollectionCapacityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// new capacity
	Capacity int32 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *UpdateCollectionCapacityRequest) Reset() {
	*x = UpdateCollectionCapacityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionCapacityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionCapacityRequest) ProtoMessage() {}

func (x *UpdateCollectionCapacityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionCapacityRequest.ProtoReflect.Descriptor instead.
func (*UpdateCollectionCapacityRequest) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCollectionCapacityRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateCollectionCapacityRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// Response for UpdateCollectionCapacity
type UpdateCollectionCapacityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UpdateCollectionCapacityResponse) Reset() {
	*x = UpdateCollectionCapacityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionCapacityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionCapacityResponse) ProtoMessage() {}

func (x *UpdateCollectionCapacityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionCapacityResponse.ProtoReflect.Descriptor instead.
func (*UpdateCollectionCapacityResponse) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCollectionCapacityResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request for UpdateCollectionName
type UpdateCollectionNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// new collection name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateCollectionNameRequest) Reset() {
	*x = UpdateCollectionNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionNameRequest) ProtoMessage() {}

func (x *UpdateCollectionNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateCollectionNameRequest) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCollectionNameRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateCollectionNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response for UpdateCollectionName
type UpdateCollectionNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UpdateCollectionNameResponse) Reset() {
	*x = UpdateCollectionNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionNameResponse) ProtoMessage() {}

func (x *UpdateCollectionNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionNameResponse.ProtoReflect.Descriptor instead.
func (*UpdateCollectionNameResponse) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCollectionNameResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetAllCollectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *GetAllCollectionsResponse) Reset() {
	*x = GetAllCollectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCollectionsResponse) ProtoMessage() {}

func (x *GetAllCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCollectionsResponse.ProtoReflect.Descriptor instead.
func (*GetAllCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

// Request for GetCollectionByKey
type GetCollectionByKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetCollectionByKeyRequest) Reset() {
	*x = GetCollectionByKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionByKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionByKeyRequest) ProtoMessage() {}

func (x *GetCollectionByKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionByKeyRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionByKeyRequest) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{9}
}

func (x *GetCollectionByKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal db collections id
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// collection name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// collection capacity
	Capacity int32 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// collection remained capacity
	Used int32 `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{10}
}

func (x *Collection) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Collection) GetUsed() int32 {
	if x != nil {
		return x.Used
	}
	return 0
}

var File_collection_collection_proto protoreflect.FileDescriptor

var file_collection_collection_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x22, 0x2f, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x2b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x2c, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x4f,
	0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22,
	0x34, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x43, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x1c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x55, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x2d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x62, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x32, 0xb2, 0x06, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x3a, 0x01, 0x2a, 0x12,
	0x7d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0xa1,
	0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x3a,
	0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x6c, 0x6c, 0x12, 0x72, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x6b, 0x61, 0x62, 0x6f,
	0x6c, 0x74, 0x2f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x64, 0x61, 0x70,
	0x70, 0x2f, 0x61, 0x72, 0x74, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_collection_proto_rawDescOnce sync.Once
	file_collection_collection_proto_rawDescData = file_collection_collection_proto_rawDesc
)

func file_collection_collection_proto_rawDescGZIP() []byte {
	file_collection_collection_proto_rawDescOnce.Do(func() {
		file_collection_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_collection_proto_rawDescData)
	})
	return file_collection_collection_proto_rawDescData
}

var file_collection_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_collection_collection_proto_goTypes = []interface{}{
	(*CreateNewCollectionRequest)(nil),       // 0: collection.CreateNewCollectionRequest
	(*CreateNewCollectionResponse)(nil),      // 1: collection.CreateNewCollectionResponse
	(*DeleteCollectionRequest)(nil),          // 2: collection.DeleteCollectionRequest
	(*DeleteCollectionResponse)(nil),         // 3: collection.DeleteCollectionResponse
	(*UpdateCollectionCapacityRequest)(nil),  // 4: collection.UpdateCollectionCapacityRequest
	(*UpdateCollectionCapacityResponse)(nil), // 5: collection.UpdateCollectionCapacityResponse
	(*UpdateCollectionNameRequest)(nil),      // 6: collection.UpdateCollectionNameRequest
	(*UpdateCollectionNameResponse)(nil),     // 7: collection.UpdateCollectionNameResponse
	(*GetAllCollectionsResponse)(nil),        // 8: collection.GetAllCollectionsResponse
	(*GetCollectionByKeyRequest)(nil),        // 9: collection.GetCollectionByKeyRequest
	(*Collection)(nil),                       // 10: collection.Collection
	(*emptypb.Empty)(nil),                    // 11: google.protobuf.Empty
}
var file_collection_collection_proto_depIdxs = []int32{
	10, // 0: collection.GetAllCollectionsResponse.collections:type_name -> collection.Collection
	0,  // 1: collection.Collections.CreateNewCollection:input_type -> collection.CreateNewCollectionRequest
	2,  // 2: collection.Collections.DeleteCollection:input_type -> collection.DeleteCollectionRequest
	4,  // 3: collection.Collections.UpdateCollectionCapacity:input_type -> collection.UpdateCollectionCapacityRequest
	6,  // 4: collection.Collections.UpdateCollectionName:input_type -> collection.UpdateCollectionNameRequest
	11, // 5: collection.Collections.GetAllCollections:input_type -> google.protobuf.Empty
	9,  // 6: collection.Collections.GetCollectionByKey:input_type -> collection.GetCollectionByKeyRequest
	1,  // 7: collection.Collections.CreateNewCollection:output_type -> collection.CreateNewCollectionResponse
	3,  // 8: collection.Collections.DeleteCollection:output_type -> collection.DeleteCollectionResponse
	5,  // 9: collection.Collections.UpdateCollectionCapacity:output_type -> collection.UpdateCollectionCapacityResponse
	7,  // 10: collection.Collections.UpdateCollectionName:output_type -> collection.UpdateCollectionNameResponse
	8,  // 11: collection.Collections.GetAllCollections:output_type -> collection.GetAllCollectionsResponse
	10, // 12: collection.Collections.GetCollectionByKey:output_type -> collection.Collection
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_collection_collection_proto_init() }
func file_collection_collection_proto_init() {
	if File_collection_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewCollectionRequest); i {
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
		file_collection_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewCollectionResponse); i {
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
		file_collection_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCollectionRequest); i {
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
		file_collection_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCollectionResponse); i {
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
		file_collection_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionCapacityRequest); i {
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
		file_collection_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionCapacityResponse); i {
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
		file_collection_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionNameRequest); i {
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
		file_collection_collection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionNameResponse); i {
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
		file_collection_collection_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCollectionsResponse); i {
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
		file_collection_collection_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionByKeyRequest); i {
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
		file_collection_collection_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
			RawDescriptor: file_collection_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_collection_proto_goTypes,
		DependencyIndexes: file_collection_collection_proto_depIdxs,
		MessageInfos:      file_collection_collection_proto_msgTypes,
	}.Build()
	File_collection_collection_proto = out.File
	file_collection_collection_proto_rawDesc = nil
	file_collection_collection_proto_goTypes = nil
	file_collection_collection_proto_depIdxs = nil
}
