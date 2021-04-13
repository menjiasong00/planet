// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: bas.proto

package pb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//产品通用
type BaslProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//bool data = 3;
	Code    int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Error   string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	Details bool   `protobuf:"varint,6,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *BaslProductsResponse) Reset() {
	*x = BaslProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsResponse) ProtoMessage() {}

func (x *BaslProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsResponse.ProtoReflect.Descriptor instead.
func (*BaslProductsResponse) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{0}
}

func (x *BaslProductsResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BaslProductsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaslProductsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaslProductsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *BaslProductsResponse) GetDetails() bool {
	if x != nil {
		return x.Details
	}
	return false
}

//产品列表
type BaslProductsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//每页多少个
	PageSize int32 `protobuf:"varint,100,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	//页码
	PageNumber int32 `protobuf:"varint,101,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	//排序字段
	OrderKey string `protobuf:"bytes,102,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	//排序方式
	OrderSort string `protobuf:"bytes,103,opt,name=orderSort,proto3" json:"orderSort,omitempty"`
}

func (x *BaslProductsListRequest) Reset() {
	*x = BaslProductsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsListRequest) ProtoMessage() {}

func (x *BaslProductsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsListRequest.ProtoReflect.Descriptor instead.
func (*BaslProductsListRequest) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{1}
}

func (x *BaslProductsListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BaslProductsListRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *BaslProductsListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *BaslProductsListRequest) GetOrderSort() string {
	if x != nil {
		return x.OrderSort
	}
	return ""
}

type BaslProductsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//BaslProductsList data = 3;
	Code    int32             `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Error   string            `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	Details *BaslProductsList `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *BaslProductsListResponse) Reset() {
	*x = BaslProductsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsListResponse) ProtoMessage() {}

func (x *BaslProductsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsListResponse.ProtoReflect.Descriptor instead.
func (*BaslProductsListResponse) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{2}
}

func (x *BaslProductsListResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BaslProductsListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaslProductsListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaslProductsListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *BaslProductsListResponse) GetDetails() *BaslProductsList {
	if x != nil {
		return x.Details
	}
	return nil
}

type BaslProductsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	//列表
	List []*BaslProductsOneRequest `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *BaslProductsList) Reset() {
	*x = BaslProductsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsList) ProtoMessage() {}

func (x *BaslProductsList) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsList.ProtoReflect.Descriptor instead.
func (*BaslProductsList) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{3}
}

func (x *BaslProductsList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BaslProductsList) GetList() []*BaslProductsOneRequest {
	if x != nil {
		return x.List
	}
	return nil
}

//产品结构
type BaslProductsOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	//
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	//
	DeletedAt string `protobuf:"bytes,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	//
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	//
	Price int32 `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *BaslProductsOneRequest) Reset() {
	*x = BaslProductsOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsOneRequest) ProtoMessage() {}

func (x *BaslProductsOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsOneRequest.ProtoReflect.Descriptor instead.
func (*BaslProductsOneRequest) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{4}
}

func (x *BaslProductsOneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaslProductsOneRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BaslProductsOneRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BaslProductsOneRequest) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *BaslProductsOneRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BaslProductsOneRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

//产品详情
type BaslProductsIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BaslProductsIdRequest) Reset() {
	*x = BaslProductsIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsIdRequest) ProtoMessage() {}

func (x *BaslProductsIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsIdRequest.ProtoReflect.Descriptor instead.
func (*BaslProductsIdRequest) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{5}
}

func (x *BaslProductsIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BaslProductsDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//BaslProductsOneRequest data = 3;
	Code    int32                   `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Error   string                  `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	Details *BaslProductsOneRequest `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *BaslProductsDetailResponse) Reset() {
	*x = BaslProductsDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaslProductsDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaslProductsDetailResponse) ProtoMessage() {}

func (x *BaslProductsDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaslProductsDetailResponse.ProtoReflect.Descriptor instead.
func (*BaslProductsDetailResponse) Descriptor() ([]byte, []int) {
	return file_bas_proto_rawDescGZIP(), []int{6}
}

func (x *BaslProductsDetailResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BaslProductsDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaslProductsDetailResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaslProductsDetailResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *BaslProductsDetailResponse) GetDetails() *BaslProductsOneRequest {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_bas_proto protoreflect.FileDescriptor

var file_bas_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01,
	0x0a, 0x14, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x8f, 0x01, 0x0a,
	0x17, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6f, 0x72, 0x74, 0x22, 0xa6,
	0x01, 0x0a, 0x18, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61,
	0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x58, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0xac, 0x01, 0x0a, 0x16, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x27, 0x0a, 0x15, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x1a, 0x42, 0x61,
	0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x9f, 0x04, 0x0a, 0x03, 0x42,
	0x61, 0x73, 0x12, 0x69, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x62, 0x61, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x70, 0x0a,
	0x12, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x62, 0x61, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x69, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x62, 0x61, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x12, 0x42, 0x61,
	0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12,
	0x2f, 0x62, 0x61, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x42, 0x61, 0x73, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x62, 0x61, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bas_proto_rawDescOnce sync.Once
	file_bas_proto_rawDescData = file_bas_proto_rawDesc
)

func file_bas_proto_rawDescGZIP() []byte {
	file_bas_proto_rawDescOnce.Do(func() {
		file_bas_proto_rawDescData = protoimpl.X.CompressGZIP(file_bas_proto_rawDescData)
	})
	return file_bas_proto_rawDescData
}

var file_bas_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bas_proto_goTypes = []interface{}{
	(*BaslProductsResponse)(nil),       // 0: pb.BaslProductsResponse
	(*BaslProductsListRequest)(nil),    // 1: pb.BaslProductsListRequest
	(*BaslProductsListResponse)(nil),   // 2: pb.BaslProductsListResponse
	(*BaslProductsList)(nil),           // 3: pb.BaslProductsList
	(*BaslProductsOneRequest)(nil),     // 4: pb.BaslProductsOneRequest
	(*BaslProductsIdRequest)(nil),      // 5: pb.BaslProductsIdRequest
	(*BaslProductsDetailResponse)(nil), // 6: pb.BaslProductsDetailResponse
}
var file_bas_proto_depIdxs = []int32{
	3, // 0: pb.BaslProductsListResponse.details:type_name -> pb.BaslProductsList
	4, // 1: pb.BaslProductsList.list:type_name -> pb.BaslProductsOneRequest
	4, // 2: pb.BaslProductsDetailResponse.details:type_name -> pb.BaslProductsOneRequest
	1, // 3: pb.Bas.BaslProductsList:input_type -> pb.BaslProductsListRequest
	5, // 4: pb.Bas.BaslProductsDetail:input_type -> pb.BaslProductsIdRequest
	4, // 5: pb.Bas.BaslProductsCreate:input_type -> pb.BaslProductsOneRequest
	4, // 6: pb.Bas.BaslProductsMotify:input_type -> pb.BaslProductsOneRequest
	5, // 7: pb.Bas.BaslProductsDelete:input_type -> pb.BaslProductsIdRequest
	2, // 8: pb.Bas.BaslProductsList:output_type -> pb.BaslProductsListResponse
	6, // 9: pb.Bas.BaslProductsDetail:output_type -> pb.BaslProductsDetailResponse
	0, // 10: pb.Bas.BaslProductsCreate:output_type -> pb.BaslProductsResponse
	0, // 11: pb.Bas.BaslProductsMotify:output_type -> pb.BaslProductsResponse
	0, // 12: pb.Bas.BaslProductsDelete:output_type -> pb.BaslProductsResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bas_proto_init() }
func file_bas_proto_init() {
	if File_bas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsResponse); i {
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
		file_bas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsListRequest); i {
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
		file_bas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsListResponse); i {
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
		file_bas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsList); i {
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
		file_bas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsOneRequest); i {
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
		file_bas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsIdRequest); i {
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
		file_bas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaslProductsDetailResponse); i {
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
			RawDescriptor: file_bas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bas_proto_goTypes,
		DependencyIndexes: file_bas_proto_depIdxs,
		MessageInfos:      file_bas_proto_msgTypes,
	}.Build()
	File_bas_proto = out.File
	file_bas_proto_rawDesc = nil
	file_bas_proto_goTypes = nil
	file_bas_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BasClient is the client API for Bas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasClient interface {
	//产品列表
	BaslProductsList(ctx context.Context, in *BaslProductsListRequest, opts ...grpc.CallOption) (*BaslProductsListResponse, error)
	//产品详情
	BaslProductsDetail(ctx context.Context, in *BaslProductsIdRequest, opts ...grpc.CallOption) (*BaslProductsDetailResponse, error)
	//产品创建
	BaslProductsCreate(ctx context.Context, in *BaslProductsOneRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error)
	//产品修改
	BaslProductsMotify(ctx context.Context, in *BaslProductsOneRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error)
	//产品删除
	BaslProductsDelete(ctx context.Context, in *BaslProductsIdRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error)
}

type basClient struct {
	cc grpc.ClientConnInterface
}

func NewBasClient(cc grpc.ClientConnInterface) BasClient {
	return &basClient{cc}
}

func (c *basClient) BaslProductsList(ctx context.Context, in *BaslProductsListRequest, opts ...grpc.CallOption) (*BaslProductsListResponse, error) {
	out := new(BaslProductsListResponse)
	err := c.cc.Invoke(ctx, "/pb.Bas/BaslProductsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basClient) BaslProductsDetail(ctx context.Context, in *BaslProductsIdRequest, opts ...grpc.CallOption) (*BaslProductsDetailResponse, error) {
	out := new(BaslProductsDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.Bas/BaslProductsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basClient) BaslProductsCreate(ctx context.Context, in *BaslProductsOneRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error) {
	out := new(BaslProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.Bas/BaslProductsCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basClient) BaslProductsMotify(ctx context.Context, in *BaslProductsOneRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error) {
	out := new(BaslProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.Bas/BaslProductsMotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basClient) BaslProductsDelete(ctx context.Context, in *BaslProductsIdRequest, opts ...grpc.CallOption) (*BaslProductsResponse, error) {
	out := new(BaslProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.Bas/BaslProductsDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasServer is the server API for Bas service.
type BasServer interface {
	//产品列表
	BaslProductsList(context.Context, *BaslProductsListRequest) (*BaslProductsListResponse, error)
	//产品详情
	BaslProductsDetail(context.Context, *BaslProductsIdRequest) (*BaslProductsDetailResponse, error)
	//产品创建
	BaslProductsCreate(context.Context, *BaslProductsOneRequest) (*BaslProductsResponse, error)
	//产品修改
	BaslProductsMotify(context.Context, *BaslProductsOneRequest) (*BaslProductsResponse, error)
	//产品删除
	BaslProductsDelete(context.Context, *BaslProductsIdRequest) (*BaslProductsResponse, error)
}

// UnimplementedBasServer can be embedded to have forward compatible implementations.
type UnimplementedBasServer struct {
}

func (*UnimplementedBasServer) BaslProductsList(context.Context, *BaslProductsListRequest) (*BaslProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaslProductsList not implemented")
}
func (*UnimplementedBasServer) BaslProductsDetail(context.Context, *BaslProductsIdRequest) (*BaslProductsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaslProductsDetail not implemented")
}
func (*UnimplementedBasServer) BaslProductsCreate(context.Context, *BaslProductsOneRequest) (*BaslProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaslProductsCreate not implemented")
}
func (*UnimplementedBasServer) BaslProductsMotify(context.Context, *BaslProductsOneRequest) (*BaslProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaslProductsMotify not implemented")
}
func (*UnimplementedBasServer) BaslProductsDelete(context.Context, *BaslProductsIdRequest) (*BaslProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaslProductsDelete not implemented")
}

func RegisterBasServer(s *grpc.Server, srv BasServer) {
	s.RegisterService(&_Bas_serviceDesc, srv)
}

func _Bas_BaslProductsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaslProductsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasServer).BaslProductsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bas/BaslProductsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasServer).BaslProductsList(ctx, req.(*BaslProductsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bas_BaslProductsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaslProductsIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasServer).BaslProductsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bas/BaslProductsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasServer).BaslProductsDetail(ctx, req.(*BaslProductsIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bas_BaslProductsCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaslProductsOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasServer).BaslProductsCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bas/BaslProductsCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasServer).BaslProductsCreate(ctx, req.(*BaslProductsOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bas_BaslProductsMotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaslProductsOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasServer).BaslProductsMotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bas/BaslProductsMotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasServer).BaslProductsMotify(ctx, req.(*BaslProductsOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bas_BaslProductsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaslProductsIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasServer).BaslProductsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Bas/BaslProductsDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasServer).BaslProductsDelete(ctx, req.(*BaslProductsIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Bas",
	HandlerType: (*BasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BaslProductsList",
			Handler:    _Bas_BaslProductsList_Handler,
		},
		{
			MethodName: "BaslProductsDetail",
			Handler:    _Bas_BaslProductsDetail_Handler,
		},
		{
			MethodName: "BaslProductsCreate",
			Handler:    _Bas_BaslProductsCreate_Handler,
		},
		{
			MethodName: "BaslProductsMotify",
			Handler:    _Bas_BaslProductsMotify_Handler,
		},
		{
			MethodName: "BaslProductsDelete",
			Handler:    _Bas_BaslProductsDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bas.proto",
}
