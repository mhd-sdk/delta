// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: request_get_instrument_by_underlying.proto

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

type RequestGetInstrumentByUnderlying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId       *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                  // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg          []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                            // PB_OFFSET + MNM_USER_MSG
	UnderlyingSymbol *string  `protobuf:"bytes,101026,opt,name=underlying_symbol,json=underlyingSymbol" json:"underlying_symbol,omitempty"` // PB_OFFSET + MNM_UNDERLYING_SYMBOL
	Exchange         *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                         // PB_OFFSET + MNM_EXCHANGE
	ExpirationDate   *string  `protobuf:"bytes,100067,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`       // PB_OFFSET + MNM_EXPIRATION_DATE
}

func (x *RequestGetInstrumentByUnderlying) Reset() {
	*x = RequestGetInstrumentByUnderlying{}
	mi := &file_request_get_instrument_by_underlying_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestGetInstrumentByUnderlying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetInstrumentByUnderlying) ProtoMessage() {}

func (x *RequestGetInstrumentByUnderlying) ProtoReflect() protoreflect.Message {
	mi := &file_request_get_instrument_by_underlying_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetInstrumentByUnderlying.ProtoReflect.Descriptor instead.
func (*RequestGetInstrumentByUnderlying) Descriptor() ([]byte, []int) {
	return file_request_get_instrument_by_underlying_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGetInstrumentByUnderlying) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestGetInstrumentByUnderlying) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestGetInstrumentByUnderlying) GetUnderlyingSymbol() string {
	if x != nil && x.UnderlyingSymbol != nil {
		return *x.UnderlyingSymbol
	}
	return ""
}

func (x *RequestGetInstrumentByUnderlying) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestGetInstrumentByUnderlying) GetExpirationDate() string {
	if x != nil && x.ExpirationDate != nil {
		return *x.ExpirationDate
	}
	return ""
}

var File_request_get_instrument_by_underlying_proto protoreflect.FileDescriptor

var file_request_get_instrument_by_underlying_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74,
	0x69, 0x22, 0xda, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x64, 0x65,
	0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x11, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c,
	0x79, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0xa2, 0x95, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xe3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_request_get_instrument_by_underlying_proto_rawDescOnce sync.Once
	file_request_get_instrument_by_underlying_proto_rawDescData = file_request_get_instrument_by_underlying_proto_rawDesc
)

func file_request_get_instrument_by_underlying_proto_rawDescGZIP() []byte {
	file_request_get_instrument_by_underlying_proto_rawDescOnce.Do(func() {
		file_request_get_instrument_by_underlying_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_get_instrument_by_underlying_proto_rawDescData)
	})
	return file_request_get_instrument_by_underlying_proto_rawDescData
}

var file_request_get_instrument_by_underlying_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_get_instrument_by_underlying_proto_goTypes = []any{
	(*RequestGetInstrumentByUnderlying)(nil), // 0: rti.RequestGetInstrumentByUnderlying
}
var file_request_get_instrument_by_underlying_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_get_instrument_by_underlying_proto_init() }
func file_request_get_instrument_by_underlying_proto_init() {
	if File_request_get_instrument_by_underlying_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_get_instrument_by_underlying_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_get_instrument_by_underlying_proto_goTypes,
		DependencyIndexes: file_request_get_instrument_by_underlying_proto_depIdxs,
		MessageInfos:      file_request_get_instrument_by_underlying_proto_msgTypes,
	}.Build()
	File_request_get_instrument_by_underlying_proto = out.File
	file_request_get_instrument_by_underlying_proto_rawDesc = nil
	file_request_get_instrument_by_underlying_proto_goTypes = nil
	file_request_get_instrument_by_underlying_proto_depIdxs = nil
}
