// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: response_depth_by_order_snapshot.proto

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

type ResponseDepthByOrderSnapshot_TransactionType int32

const (
	ResponseDepthByOrderSnapshot_BUY  ResponseDepthByOrderSnapshot_TransactionType = 1
	ResponseDepthByOrderSnapshot_SELL ResponseDepthByOrderSnapshot_TransactionType = 2
)

// Enum value maps for ResponseDepthByOrderSnapshot_TransactionType.
var (
	ResponseDepthByOrderSnapshot_TransactionType_name = map[int32]string{
		1: "BUY",
		2: "SELL",
	}
	ResponseDepthByOrderSnapshot_TransactionType_value = map[string]int32{
		"BUY":  1,
		"SELL": 2,
	}
)

func (x ResponseDepthByOrderSnapshot_TransactionType) Enum() *ResponseDepthByOrderSnapshot_TransactionType {
	p := new(ResponseDepthByOrderSnapshot_TransactionType)
	*p = x
	return p
}

func (x ResponseDepthByOrderSnapshot_TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseDepthByOrderSnapshot_TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_response_depth_by_order_snapshot_proto_enumTypes[0].Descriptor()
}

func (ResponseDepthByOrderSnapshot_TransactionType) Type() protoreflect.EnumType {
	return &file_response_depth_by_order_snapshot_proto_enumTypes[0]
}

func (x ResponseDepthByOrderSnapshot_TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseDepthByOrderSnapshot_TransactionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseDepthByOrderSnapshot_TransactionType(num)
	return nil
}

// Deprecated: Use ResponseDepthByOrderSnapshot_TransactionType.Descriptor instead.
func (ResponseDepthByOrderSnapshot_TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_response_depth_by_order_snapshot_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseDepthByOrderSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId         *int32                                        `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                                    // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg            []string                                      `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                                              // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode    []string                                      `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`                                  // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode             []string                                      `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                                                 // PB_OFFSET + MNM_RESPONSE_CODE
	Exchange           *string                                       `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                                           // PB_OFFSET + MNM_EXCHANGE
	Symbol             *string                                       `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                                               // PB_OFFSET + MNM_SYMBOL
	SequenceNumber     *uint64                                       `protobuf:"varint,112002,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`                                        // PB_OFFSET + MNM_SEQUENCE_NUMBER
	DepthSide          *ResponseDepthByOrderSnapshot_TransactionType `protobuf:"varint,153612,opt,name=depth_side,json=depthSide,enum=rti.ResponseDepthByOrderSnapshot_TransactionType" json:"depth_side,omitempty"` // PB_OFFSET + MNM_MARKET_DEPTH_SIDE
	DepthPrice         *float64                                      `protobuf:"fixed64,154405,opt,name=depth_price,json=depthPrice" json:"depth_price,omitempty"`                                                   // PB_OFFSET + MNM_MARKET_DEPTH_PRICE
	DepthSize          []int32                                       `protobuf:"varint,154406,rep,name=depth_size,json=depthSize" json:"depth_size,omitempty"`                                                       // PB_OFFSET + MNM_MARKET_DEPTH_SIZE
	DepthOrderPriority []uint64                                      `protobuf:"varint,153613,rep,name=depth_order_priority,json=depthOrderPriority" json:"depth_order_priority,omitempty"`                          // PB_OFFSET + MNM_MARKET_DEPTH_ORDER_PRIORITY
	ExchangeOrderId    []string                                      `protobuf:"bytes,149238,rep,name=exchange_order_id,json=exchangeOrderId" json:"exchange_order_id,omitempty"`                                    // PB_OFFSET + MNM_EXCH_ORD_ID
}

func (x *ResponseDepthByOrderSnapshot) Reset() {
	*x = ResponseDepthByOrderSnapshot{}
	mi := &file_response_depth_by_order_snapshot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseDepthByOrderSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDepthByOrderSnapshot) ProtoMessage() {}

func (x *ResponseDepthByOrderSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_response_depth_by_order_snapshot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDepthByOrderSnapshot.ProtoReflect.Descriptor instead.
func (*ResponseDepthByOrderSnapshot) Descriptor() ([]byte, []int) {
	return file_response_depth_by_order_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseDepthByOrderSnapshot) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseDepthByOrderSnapshot) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseDepthByOrderSnapshot) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseDepthByOrderSnapshot) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseDepthByOrderSnapshot) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *ResponseDepthByOrderSnapshot) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *ResponseDepthByOrderSnapshot) GetSequenceNumber() uint64 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

func (x *ResponseDepthByOrderSnapshot) GetDepthSide() ResponseDepthByOrderSnapshot_TransactionType {
	if x != nil && x.DepthSide != nil {
		return *x.DepthSide
	}
	return ResponseDepthByOrderSnapshot_BUY
}

func (x *ResponseDepthByOrderSnapshot) GetDepthPrice() float64 {
	if x != nil && x.DepthPrice != nil {
		return *x.DepthPrice
	}
	return 0
}

func (x *ResponseDepthByOrderSnapshot) GetDepthSize() []int32 {
	if x != nil {
		return x.DepthSize
	}
	return nil
}

func (x *ResponseDepthByOrderSnapshot) GetDepthOrderPriority() []uint64 {
	if x != nil {
		return x.DepthOrderPriority
	}
	return nil
}

func (x *ResponseDepthByOrderSnapshot) GetExchangeOrderId() []string {
	if x != nil {
		return x.ExchangeOrderId
	}
	return nil
}

var File_response_depth_by_order_snapshot_proto protoreflect.FileDescriptor

var file_response_depth_by_order_snapshot_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xab, 0x04,
	0x0a, 0x1c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x21,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6,
	0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2d,
	0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x71,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a,
	0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x82, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x74, 0x68, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x8c, 0xb0, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x64, 0x65, 0x70, 0x74, 0x68, 0x53, 0x69, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xa5,
	0xb6, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x74, 0x68, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0xa6, 0xb6, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x70, 0x74, 0x68, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x8d, 0xb0, 0x09, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x12, 0x64, 0x65, 0x70, 0x74, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0xf6, 0x8d, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x72, 0x74, 0x69,
}

var (
	file_response_depth_by_order_snapshot_proto_rawDescOnce sync.Once
	file_response_depth_by_order_snapshot_proto_rawDescData = file_response_depth_by_order_snapshot_proto_rawDesc
)

func file_response_depth_by_order_snapshot_proto_rawDescGZIP() []byte {
	file_response_depth_by_order_snapshot_proto_rawDescOnce.Do(func() {
		file_response_depth_by_order_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_depth_by_order_snapshot_proto_rawDescData)
	})
	return file_response_depth_by_order_snapshot_proto_rawDescData
}

var file_response_depth_by_order_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_depth_by_order_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_depth_by_order_snapshot_proto_goTypes = []any{
	(ResponseDepthByOrderSnapshot_TransactionType)(0), // 0: rti.ResponseDepthByOrderSnapshot.TransactionType
	(*ResponseDepthByOrderSnapshot)(nil),              // 1: rti.ResponseDepthByOrderSnapshot
}
var file_response_depth_by_order_snapshot_proto_depIdxs = []int32{
	0, // 0: rti.ResponseDepthByOrderSnapshot.depth_side:type_name -> rti.ResponseDepthByOrderSnapshot.TransactionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_depth_by_order_snapshot_proto_init() }
func file_response_depth_by_order_snapshot_proto_init() {
	if File_response_depth_by_order_snapshot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_depth_by_order_snapshot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_depth_by_order_snapshot_proto_goTypes,
		DependencyIndexes: file_response_depth_by_order_snapshot_proto_depIdxs,
		EnumInfos:         file_response_depth_by_order_snapshot_proto_enumTypes,
		MessageInfos:      file_response_depth_by_order_snapshot_proto_msgTypes,
	}.Build()
	File_response_depth_by_order_snapshot_proto = out.File
	file_response_depth_by_order_snapshot_proto_rawDesc = nil
	file_response_depth_by_order_snapshot_proto_goTypes = nil
	file_response_depth_by_order_snapshot_proto_depIdxs = nil
}
