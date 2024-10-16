// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: depth_by_order.proto

package rti

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

type DepthByOrder_TransactionType int32

const (
	DepthByOrder_BUY  DepthByOrder_TransactionType = 1
	DepthByOrder_SELL DepthByOrder_TransactionType = 2
)

// Enum value maps for DepthByOrder_TransactionType.
var (
	DepthByOrder_TransactionType_name = map[int32]string{
		1: "BUY",
		2: "SELL",
	}
	DepthByOrder_TransactionType_value = map[string]int32{
		"BUY":  1,
		"SELL": 2,
	}
)

func (x DepthByOrder_TransactionType) Enum() *DepthByOrder_TransactionType {
	p := new(DepthByOrder_TransactionType)
	*p = x
	return p
}

func (x DepthByOrder_TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DepthByOrder_TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_depth_by_order_proto_enumTypes[0].Descriptor()
}

func (DepthByOrder_TransactionType) Type() protoreflect.EnumType {
	return &file_depth_by_order_proto_enumTypes[0]
}

func (x DepthByOrder_TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DepthByOrder_TransactionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DepthByOrder_TransactionType(num)
	return nil
}

// Deprecated: Use DepthByOrder_TransactionType.Descriptor instead.
func (DepthByOrder_TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_depth_by_order_proto_rawDescGZIP(), []int{0, 0}
}

type DepthByOrder_UpdateType int32

const (
	DepthByOrder_NEW    DepthByOrder_UpdateType = 1
	DepthByOrder_CHANGE DepthByOrder_UpdateType = 2
	DepthByOrder_DELETE DepthByOrder_UpdateType = 3
)

// Enum value maps for DepthByOrder_UpdateType.
var (
	DepthByOrder_UpdateType_name = map[int32]string{
		1: "NEW",
		2: "CHANGE",
		3: "DELETE",
	}
	DepthByOrder_UpdateType_value = map[string]int32{
		"NEW":    1,
		"CHANGE": 2,
		"DELETE": 3,
	}
)

func (x DepthByOrder_UpdateType) Enum() *DepthByOrder_UpdateType {
	p := new(DepthByOrder_UpdateType)
	*p = x
	return p
}

func (x DepthByOrder_UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DepthByOrder_UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_depth_by_order_proto_enumTypes[1].Descriptor()
}

func (DepthByOrder_UpdateType) Type() protoreflect.EnumType {
	return &file_depth_by_order_proto_enumTypes[1]
}

func (x DepthByOrder_UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DepthByOrder_UpdateType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DepthByOrder_UpdateType(num)
	return nil
}

// Deprecated: Use DepthByOrder_UpdateType.Descriptor instead.
func (DepthByOrder_UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_depth_by_order_proto_rawDescGZIP(), []int{0, 1}
}

type DepthByOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId         *int32                         `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                                      // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol             *string                        `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                                                 // PB_OFFSET + MNM_SYMBOL
	Exchange           *string                        `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                                             // PB_OFFSET + MNM_EXCHANGE
	SequenceNumber     *uint64                        `protobuf:"varint,112002,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`                                          // PB_OFFSET + MNM_SEQUENCE_NUMBER
	UpdateType         []DepthByOrder_UpdateType      `protobuf:"varint,110121,rep,name=update_type,json=updateType,enum=rti.DepthByOrder_UpdateType" json:"update_type,omitempty"`                     // PB_OFFSET + MNM_UPDATE_TYPE
	TransactionType    []DepthByOrder_TransactionType `protobuf:"varint,153612,rep,name=transaction_type,json=transactionType,enum=rti.DepthByOrder_TransactionType" json:"transaction_type,omitempty"` // PB_OFFSET + MNM_MARKET_DEPTH_SIDE
	DepthPrice         []float64                      `protobuf:"fixed64,154405,rep,name=depth_price,json=depthPrice" json:"depth_price,omitempty"`                                                     // PB_OFFSET + MNM_MARKET_DEPTH_PRICE
	PrevDepthPrice     []float64                      `protobuf:"fixed64,154906,rep,name=prev_depth_price,json=prevDepthPrice" json:"prev_depth_price,omitempty"`                                       // PB_OFFSET + MNM_PREVIOUS_MARKET_DEPTH_PRICE
	PrevDepthPriceFlag []bool                         `protobuf:"varint,154930,rep,name=prev_depth_price_flag,json=prevDepthPriceFlag" json:"prev_depth_price_flag,omitempty"`                          // PB_OFFSET + MNM_PREVIOUS_MARKET_DEPTH_PRICE_FLAG
	DepthSize          []int32                        `protobuf:"varint,154406,rep,name=depth_size,json=depthSize" json:"depth_size,omitempty"`                                                         // PB_OFFSET + MNM_MARKET_DEPTH_SIZE
	DepthOrderPriority []uint64                       `protobuf:"varint,153613,rep,name=depth_order_priority,json=depthOrderPriority" json:"depth_order_priority,omitempty"`                            // PB_OFFSET + MNM_MARKET_DEPTH_ORDER_PRIORITY
	ExchangeOrderId    []string                       `protobuf:"bytes,149238,rep,name=exchange_order_id,json=exchangeOrderId" json:"exchange_order_id,omitempty"`                                      // PB_OFFSET + MNM_EXCH_ORD_ID
	Ssboe              *int32                         `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                                                                  // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs              *int32                         `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                                                                  // PB_OFFSET + MNM_USECS
	SourceSsboe        *int32                         `protobuf:"varint,150400,opt,name=source_ssboe,json=sourceSsboe" json:"source_ssboe,omitempty"`                                                   // PB_OFFSET + MNM_SOURCE_SSBOE
	SourceUsecs        *int32                         `protobuf:"varint,150401,opt,name=source_usecs,json=sourceUsecs" json:"source_usecs,omitempty"`                                                   // PB_OFFSET + MNM_SOURCE_USECS
	SourceNsecs        *int32                         `protobuf:"varint,150404,opt,name=source_nsecs,json=sourceNsecs" json:"source_nsecs,omitempty"`                                                   // PB_OFFSET + MNM_SOURCE_NSECS
	JopSsboe           *int32                         `protobuf:"varint,150600,opt,name=jop_ssboe,json=jopSsboe" json:"jop_ssboe,omitempty"`                                                            // PB_OFFSET + MNM_JOP_SSBOE
	JopNsecs           *int32                         `protobuf:"varint,150604,opt,name=jop_nsecs,json=jopNsecs" json:"jop_nsecs,omitempty"`                                                            // PB_OFFSET + MNM_JOP_NSECS
}

func (x *DepthByOrder) Reset() {
	*x = DepthByOrder{}
	mi := &file_depth_by_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepthByOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepthByOrder) ProtoMessage() {}

func (x *DepthByOrder) ProtoReflect() protoreflect.Message {
	mi := &file_depth_by_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepthByOrder.ProtoReflect.Descriptor instead.
func (*DepthByOrder) Descriptor() ([]byte, []int) {
	return file_depth_by_order_proto_rawDescGZIP(), []int{0}
}

func (x *DepthByOrder) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *DepthByOrder) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *DepthByOrder) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *DepthByOrder) GetSequenceNumber() uint64 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

func (x *DepthByOrder) GetUpdateType() []DepthByOrder_UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return nil
}

func (x *DepthByOrder) GetTransactionType() []DepthByOrder_TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return nil
}

func (x *DepthByOrder) GetDepthPrice() []float64 {
	if x != nil {
		return x.DepthPrice
	}
	return nil
}

func (x *DepthByOrder) GetPrevDepthPrice() []float64 {
	if x != nil {
		return x.PrevDepthPrice
	}
	return nil
}

func (x *DepthByOrder) GetPrevDepthPriceFlag() []bool {
	if x != nil {
		return x.PrevDepthPriceFlag
	}
	return nil
}

func (x *DepthByOrder) GetDepthSize() []int32 {
	if x != nil {
		return x.DepthSize
	}
	return nil
}

func (x *DepthByOrder) GetDepthOrderPriority() []uint64 {
	if x != nil {
		return x.DepthOrderPriority
	}
	return nil
}

func (x *DepthByOrder) GetExchangeOrderId() []string {
	if x != nil {
		return x.ExchangeOrderId
	}
	return nil
}

func (x *DepthByOrder) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *DepthByOrder) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

func (x *DepthByOrder) GetSourceSsboe() int32 {
	if x != nil && x.SourceSsboe != nil {
		return *x.SourceSsboe
	}
	return 0
}

func (x *DepthByOrder) GetSourceUsecs() int32 {
	if x != nil && x.SourceUsecs != nil {
		return *x.SourceUsecs
	}
	return 0
}

func (x *DepthByOrder) GetSourceNsecs() int32 {
	if x != nil && x.SourceNsecs != nil {
		return *x.SourceNsecs
	}
	return 0
}

func (x *DepthByOrder) GetJopSsboe() int32 {
	if x != nil && x.JopSsboe != nil {
		return *x.JopSsboe
	}
	return 0
}

func (x *DepthByOrder) GetJopNsecs() int32 {
	if x != nil && x.JopNsecs != nil {
		return *x.JopNsecs
	}
	return 0
}

var File_depth_by_order_proto protoreflect.FileDescriptor

var file_depth_by_order_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xde, 0x06, 0x0a, 0x0c,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x82, 0xeb, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0xa9, 0xdc, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x74, 0x69, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x8c, 0xb0, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0xa5, 0xb6, 0x09, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x74,
	0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x9a, 0xba, 0x09, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x44, 0x65, 0x70, 0x74, 0x68, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0xb2, 0xba, 0x09, 0x20,
	0x03, 0x28, 0x08, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76, 0x44, 0x65, 0x70, 0x74, 0x68, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xa6, 0xb6, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x70, 0x74,
	0x68, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x8d, 0xb0, 0x09, 0x20, 0x03, 0x28, 0x04, 0x52, 0x12, 0x64, 0x65, 0x70, 0x74, 0x68, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0xf6, 0x8d, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73,
	0x62, 0x6f, 0x65, 0x18, 0xd4, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62,
	0x6f, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0x80, 0x97, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12,
	0x23, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18,
	0x81, 0x97, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x73, 0x65, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x73, 0x65, 0x63, 0x73, 0x18, 0x84, 0x97, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x73, 0x65, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x6f, 0x70,
	0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xc8, 0x98, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6a, 0x6f, 0x70, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x6f, 0x70, 0x5f,
	0x6e, 0x73, 0x65, 0x63, 0x73, 0x18, 0xcc, 0x98, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a,
	0x6f, 0x70, 0x4e, 0x73, 0x65, 0x63, 0x73, 0x22, 0x24, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55,
	0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x02, 0x22, 0x2d, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e,
	0x45, 0x57, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_depth_by_order_proto_rawDescOnce sync.Once
	file_depth_by_order_proto_rawDescData = file_depth_by_order_proto_rawDesc
)

func file_depth_by_order_proto_rawDescGZIP() []byte {
	file_depth_by_order_proto_rawDescOnce.Do(func() {
		file_depth_by_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_depth_by_order_proto_rawDescData)
	})
	return file_depth_by_order_proto_rawDescData
}

var file_depth_by_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_depth_by_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_depth_by_order_proto_goTypes = []any{
	(DepthByOrder_TransactionType)(0), // 0: rti.DepthByOrder.TransactionType
	(DepthByOrder_UpdateType)(0),      // 1: rti.DepthByOrder.UpdateType
	(*DepthByOrder)(nil),              // 2: rti.DepthByOrder
}
var file_depth_by_order_proto_depIdxs = []int32{
	1, // 0: rti.DepthByOrder.update_type:type_name -> rti.DepthByOrder.UpdateType
	0, // 1: rti.DepthByOrder.transaction_type:type_name -> rti.DepthByOrder.TransactionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_depth_by_order_proto_init() }
func file_depth_by_order_proto_init() {
	if File_depth_by_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_depth_by_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_depth_by_order_proto_goTypes,
		DependencyIndexes: file_depth_by_order_proto_depIdxs,
		EnumInfos:         file_depth_by_order_proto_enumTypes,
		MessageInfos:      file_depth_by_order_proto_msgTypes,
	}.Build()
	File_depth_by_order_proto = out.File
	file_depth_by_order_proto_rawDesc = nil
	file_depth_by_order_proto_goTypes = nil
	file_depth_by_order_proto_depIdxs = nil
}
