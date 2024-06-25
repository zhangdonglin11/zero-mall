// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: goods.proto

// proto 包名

package goods

import (
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

// 商品
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{0}
}

type GoodsFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
}

func (x *GoodsFilterRequest) Reset() {
	*x = GoodsFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsFilterRequest) ProtoMessage() {}

func (x *GoodsFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsFilterRequest.ProtoReflect.Descriptor instead.
func (*GoodsFilterRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsFilterRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GoodsFilterRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type BatchGoodsIdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *BatchGoodsIdInfo) Reset() {
	*x = BatchGoodsIdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGoodsIdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGoodsIdInfo) ProtoMessage() {}

func (x *BatchGoodsIdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGoodsIdInfo.ProtoReflect.Descriptor instead.
func (*BatchGoodsIdInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGoodsIdInfo) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type CreateGoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GoodsSn    string `protobuf:"bytes,3,opt,name=goodsSn,proto3" json:"goodsSn,omitempty"`
	GoodsBrief string `protobuf:"bytes,4,opt,name=goodsBrief,proto3" json:"goodsBrief,omitempty"`
	ShopPrice  int64  `protobuf:"varint,5,opt,name=shopPrice,proto3" json:"shopPrice,omitempty"`
	Status     int64  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateGoodsInfo) Reset() {
	*x = CreateGoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsInfo) ProtoMessage() {}

func (x *CreateGoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsInfo.ProtoReflect.Descriptor instead.
func (*CreateGoodsInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGoodsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateGoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoodsInfo) GetGoodsSn() string {
	if x != nil {
		return x.GoodsSn
	}
	return ""
}

func (x *CreateGoodsInfo) GetGoodsBrief() string {
	if x != nil {
		return x.GoodsBrief
	}
	return ""
}

func (x *CreateGoodsInfo) GetShopPrice() int64 {
	if x != nil {
		return x.ShopPrice
	}
	return 0
}

func (x *CreateGoodsInfo) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DeleteGoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGoodsInfo) Reset() {
	*x = DeleteGoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsInfo) ProtoMessage() {}

func (x *DeleteGoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsInfo.ProtoReflect.Descriptor instead.
func (*DeleteGoodsInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGoodsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GoodInfoRequest) Reset() {
	*x = GoodInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodInfoRequest) ProtoMessage() {}

func (x *GoodInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodInfoRequest.ProtoReflect.Descriptor instead.
func (*GoodInfoRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GoodInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*GoodsInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GoodsListResponse) Reset() {
	*x = GoodsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListResponse) ProtoMessage() {}

func (x *GoodsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListResponse.ProtoReflect.Descriptor instead.
func (*GoodsListResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{6}
}

func (x *GoodsListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GoodsListResponse) GetData() []*GoodsInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type GoodsInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GoodsSn    string `protobuf:"bytes,3,opt,name=goodsSn,proto3" json:"goodsSn,omitempty"`
	GoodsBrief string `protobuf:"bytes,4,opt,name=goodsBrief,proto3" json:"goodsBrief,omitempty"`
	ShopPrice  int64  `protobuf:"varint,5,opt,name=shopPrice,proto3" json:"shopPrice,omitempty"`
	Status     int64  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GoodsInfoResponse) Reset() {
	*x = GoodsInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfoResponse) ProtoMessage() {}

func (x *GoodsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfoResponse.ProtoReflect.Descriptor instead.
func (*GoodsInfoResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{7}
}

func (x *GoodsInfoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfoResponse) GetGoodsSn() string {
	if x != nil {
		return x.GoodsSn
	}
	return ""
}

func (x *GoodsInfoResponse) GetGoodsBrief() string {
	if x != nil {
		return x.GoodsBrief
	}
	return ""
}

func (x *GoodsInfoResponse) GetShopPrice() int64 {
	if x != nil {
		return x.ShopPrice
	}
	return 0
}

func (x *GoodsInfoResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_goods_proto protoreflect.FileDescriptor

var file_goods_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x44, 0x0a,
	0x12, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x42, 0x72, 0x69, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa7,
	0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x53, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x53, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x72, 0x69, 0x65, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xfc, 0x02, 0x0a, 0x05, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_proto_rawDescOnce sync.Once
	file_goods_proto_rawDescData = file_goods_proto_rawDesc
)

func file_goods_proto_rawDescGZIP() []byte {
	file_goods_proto_rawDescOnce.Do(func() {
		file_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_proto_rawDescData)
	})
	return file_goods_proto_rawDescData
}

var file_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goods_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: goods.Empty
	(*GoodsFilterRequest)(nil), // 1: goods.GoodsFilterRequest
	(*BatchGoodsIdInfo)(nil),   // 2: goods.BatchGoodsIdInfo
	(*CreateGoodsInfo)(nil),    // 3: goods.CreateGoodsInfo
	(*DeleteGoodsInfo)(nil),    // 4: goods.DeleteGoodsInfo
	(*GoodInfoRequest)(nil),    // 5: goods.GoodInfoRequest
	(*GoodsListResponse)(nil),  // 6: goods.GoodsListResponse
	(*GoodsInfoResponse)(nil),  // 7: goods.GoodsInfoResponse
}
var file_goods_proto_depIdxs = []int32{
	7, // 0: goods.GoodsListResponse.data:type_name -> goods.GoodsInfoResponse
	1, // 1: goods.Goods.GoodsList:input_type -> goods.GoodsFilterRequest
	2, // 2: goods.Goods.BatchGetGoods:input_type -> goods.BatchGoodsIdInfo
	3, // 3: goods.Goods.CreateGoods:input_type -> goods.CreateGoodsInfo
	4, // 4: goods.Goods.DeleteGoods:input_type -> goods.DeleteGoodsInfo
	3, // 5: goods.Goods.UpdateGoods:input_type -> goods.CreateGoodsInfo
	5, // 6: goods.Goods.GetGoodsDetail:input_type -> goods.GoodInfoRequest
	6, // 7: goods.Goods.GoodsList:output_type -> goods.GoodsListResponse
	6, // 8: goods.Goods.BatchGetGoods:output_type -> goods.GoodsListResponse
	7, // 9: goods.Goods.CreateGoods:output_type -> goods.GoodsInfoResponse
	0, // 10: goods.Goods.DeleteGoods:output_type -> goods.Empty
	0, // 11: goods.Goods.UpdateGoods:output_type -> goods.Empty
	7, // 12: goods.Goods.GetGoodsDetail:output_type -> goods.GoodsInfoResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goods_proto_init() }
func file_goods_proto_init() {
	if File_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsFilterRequest); i {
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
		file_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGoodsIdInfo); i {
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
		file_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsInfo); i {
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
		file_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsInfo); i {
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
		file_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodInfoRequest); i {
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
		file_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListResponse); i {
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
		file_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfoResponse); i {
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
			RawDescriptor: file_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_proto_goTypes,
		DependencyIndexes: file_goods_proto_depIdxs,
		MessageInfos:      file_goods_proto_msgTypes,
	}.Build()
	File_goods_proto = out.File
	file_goods_proto_rawDesc = nil
	file_goods_proto_goTypes = nil
	file_goods_proto_depIdxs = nil
}