// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: bid.proto

package genprotos

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

type SubmitBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenderId     string  `protobuf:"bytes,1,opt,name=tender_id,json=tenderId,proto3" json:"tender_id,omitempty"`
	ContractorId string  `protobuf:"bytes,2,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
	Price        float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryTime int32   `protobuf:"varint,4,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"` // in days
	Comments     string  `protobuf:"bytes,5,opt,name=comments,proto3" json:"comments,omitempty"`                              // Optional
}

func (x *SubmitBidRequest) Reset() {
	*x = SubmitBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitBidRequest) ProtoMessage() {}

func (x *SubmitBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitBidRequest.ProtoReflect.Descriptor instead.
func (*SubmitBidRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitBidRequest) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *SubmitBidRequest) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

func (x *SubmitBidRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SubmitBidRequest) GetDeliveryTime() int32 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

func (x *SubmitBidRequest) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllByid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAllByid) Reset() {
	*x = GetAllByid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByid) ProtoMessage() {}

func (x *GetAllByid) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByid.ProtoReflect.Descriptor instead.
func (*GetAllByid) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllByid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenderId     string             `protobuf:"bytes,1,opt,name=tender_id,json=tenderId,proto3" json:"tender_id,omitempty"`
	Price        float32            `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryTime int32              `protobuf:"varint,3,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"` // in days
	Comments     string             `protobuf:"bytes,4,opt,name=comments,proto3" json:"comments,omitempty"`
	Status       string             `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"` // "pending", "awarded"
	CreatedAt    string             `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Tenders      *GetTenderResponse `protobuf:"bytes,7,opt,name=Tenders,proto3" json:"Tenders,omitempty"`
}

func (x *GetAllBidResponse) Reset() {
	*x = GetAllBidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBidResponse) ProtoMessage() {}

func (x *GetAllBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBidResponse.ProtoReflect.Descriptor instead.
func (*GetAllBidResponse) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllBidResponse) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *GetAllBidResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetAllBidResponse) GetDeliveryTime() int32 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

func (x *GetAllBidResponse) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *GetAllBidResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetAllBidResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetAllBidResponse) GetTenders() *GetTenderResponse {
	if x != nil {
		return x.Tenders
	}
	return nil
}

type ListBidsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price        float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryTime int32   `protobuf:"varint,2,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"`
	Limit        int32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int32   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListBidsRequest) Reset() {
	*x = ListBidsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBidsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBidsRequest) ProtoMessage() {}

func (x *ListBidsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBidsRequest.ProtoReflect.Descriptor instead.
func (*ListBidsRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{4}
}

func (x *ListBidsRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListBidsRequest) GetDeliveryTime() int32 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

func (x *ListBidsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBidsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListBidsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bids []*GetAllBidResponse `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids,omitempty"`
}

func (x *ListBidsResponse) Reset() {
	*x = ListBidsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBidsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBidsResponse) ProtoMessage() {}

func (x *ListBidsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBidsResponse.ProtoReflect.Descriptor instead.
func (*ListBidsResponse) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{5}
}

func (x *ListBidsResponse) GetBids() []*GetAllBidResponse {
	if x != nil {
		return x.Bids
	}
	return nil
}

type GetAllBidsByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractorId string             `protobuf:"bytes,1,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
	Price        float32            `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	DeliveryTime int32              `protobuf:"varint,3,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"`
	Comments     string             `protobuf:"bytes,4,opt,name=comments,proto3" json:"comments,omitempty"`
	Tenders      *GetTenderResponse `protobuf:"bytes,5,opt,name=Tenders,proto3" json:"Tenders,omitempty"`
}

func (x *GetAllBidsByUser) Reset() {
	*x = GetAllBidsByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBidsByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBidsByUser) ProtoMessage() {}

func (x *GetAllBidsByUser) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBidsByUser.ProtoReflect.Descriptor instead.
func (*GetAllBidsByUser) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllBidsByUser) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

func (x *GetAllBidsByUser) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetAllBidsByUser) GetDeliveryTime() int32 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

func (x *GetAllBidsByUser) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *GetAllBidsByUser) GetTenders() *GetTenderResponse {
	if x != nil {
		return x.Tenders
	}
	return nil
}

type GetAllBidsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bids []*GetAllBidsByUser `protobuf:"bytes,1,rep,name=Bids,proto3" json:"Bids,omitempty"`
}

func (x *GetAllBidsByUserIdRequest) Reset() {
	*x = GetAllBidsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBidsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBidsByUserIdRequest) ProtoMessage() {}

func (x *GetAllBidsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBidsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetAllBidsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllBidsByUserIdRequest) GetBids() []*GetAllBidsByUser {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_bid_proto protoreflect.FileDescriptor

var file_bid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x1a, 0x0c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x27, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x79, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf3, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x54, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x07, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x7a, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x41, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x62, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07,
	0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x49, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x04, 0x42, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x42, 0x69, 0x64, 0x73, 0x32, 0x9a, 0x02, 0x0a,
	0x0a, 0x42, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x69, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x69, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x69, 0x64, 0x73, 0x42, 0x79, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79,
	0x69, 0x64, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x69,
	0x64, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bid_proto_rawDescOnce sync.Once
	file_bid_proto_rawDescData = file_bid_proto_rawDesc
)

func file_bid_proto_rawDescGZIP() []byte {
	file_bid_proto_rawDescOnce.Do(func() {
		file_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bid_proto_rawDescData)
	})
	return file_bid_proto_rawDescData
}

var file_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bid_proto_goTypes = []interface{}{
	(*SubmitBidRequest)(nil),          // 0: protos.SubmitBidRequest
	(*BidResponse)(nil),               // 1: protos.BidResponse
	(*GetAllByid)(nil),                // 2: protos.GetAllByid
	(*GetAllBidResponse)(nil),         // 3: protos.GetAllBidResponse
	(*ListBidsRequest)(nil),           // 4: protos.ListBidsRequest
	(*ListBidsResponse)(nil),          // 5: protos.ListBidsResponse
	(*GetAllBidsByUser)(nil),          // 6: protos.GetAllBidsByUser
	(*GetAllBidsByUserIdRequest)(nil), // 7: protos.GetAllBidsByUserIdRequest
	(*GetTenderResponse)(nil),         // 8: protos.GetTenderResponse
}
var file_bid_proto_depIdxs = []int32{
	8, // 0: protos.GetAllBidResponse.Tenders:type_name -> protos.GetTenderResponse
	3, // 1: protos.ListBidsResponse.bids:type_name -> protos.GetAllBidResponse
	8, // 2: protos.GetAllBidsByUser.Tenders:type_name -> protos.GetTenderResponse
	6, // 3: protos.GetAllBidsByUserIdRequest.Bids:type_name -> protos.GetAllBidsByUser
	0, // 4: protos.BidService.SubmitBid:input_type -> protos.SubmitBidRequest
	4, // 5: protos.BidService.ListBids:input_type -> protos.ListBidsRequest
	2, // 6: protos.BidService.GetAllBidsByTenderId:input_type -> protos.GetAllByid
	2, // 7: protos.BidService.ListContractorBids:input_type -> protos.GetAllByid
	1, // 8: protos.BidService.SubmitBid:output_type -> protos.BidResponse
	5, // 9: protos.BidService.ListBids:output_type -> protos.ListBidsResponse
	5, // 10: protos.BidService.GetAllBidsByTenderId:output_type -> protos.ListBidsResponse
	7, // 11: protos.BidService.ListContractorBids:output_type -> protos.GetAllBidsByUserIdRequest
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bid_proto_init() }
func file_bid_proto_init() {
	if File_bid_proto != nil {
		return
	}
	file_tender_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitBidRequest); i {
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
		file_bid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_bid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByid); i {
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
		file_bid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBidResponse); i {
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
		file_bid_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBidsRequest); i {
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
		file_bid_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBidsResponse); i {
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
		file_bid_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBidsByUser); i {
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
		file_bid_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBidsByUserIdRequest); i {
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
			RawDescriptor: file_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bid_proto_goTypes,
		DependencyIndexes: file_bid_proto_depIdxs,
		MessageInfos:      file_bid_proto_msgTypes,
	}.Build()
	File_bid_proto = out.File
	file_bid_proto_rawDesc = nil
	file_bid_proto_goTypes = nil
	file_bid_proto_depIdxs = nil
}
