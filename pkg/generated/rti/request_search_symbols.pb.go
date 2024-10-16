// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: request_search_symbols.proto

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

type RequestSearchSymbols_Pattern int32

const (
	RequestSearchSymbols_EQUALS   RequestSearchSymbols_Pattern = 1
	RequestSearchSymbols_CONTAINS RequestSearchSymbols_Pattern = 2
)

// Enum value maps for RequestSearchSymbols_Pattern.
var (
	RequestSearchSymbols_Pattern_name = map[int32]string{
		1: "EQUALS",
		2: "CONTAINS",
	}
	RequestSearchSymbols_Pattern_value = map[string]int32{
		"EQUALS":   1,
		"CONTAINS": 2,
	}
)

func (x RequestSearchSymbols_Pattern) Enum() *RequestSearchSymbols_Pattern {
	p := new(RequestSearchSymbols_Pattern)
	*p = x
	return p
}

func (x RequestSearchSymbols_Pattern) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestSearchSymbols_Pattern) Descriptor() protoreflect.EnumDescriptor {
	return file_request_search_symbols_proto_enumTypes[0].Descriptor()
}

func (RequestSearchSymbols_Pattern) Type() protoreflect.EnumType {
	return &file_request_search_symbols_proto_enumTypes[0]
}

func (x RequestSearchSymbols_Pattern) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestSearchSymbols_Pattern) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestSearchSymbols_Pattern(num)
	return nil
}

// Deprecated: Use RequestSearchSymbols_Pattern.Descriptor instead.
func (RequestSearchSymbols_Pattern) EnumDescriptor() ([]byte, []int) {
	return file_request_search_symbols_proto_rawDescGZIP(), []int{0, 0}
}

type RequestSearchSymbols_InstrumentType int32

const (
	RequestSearchSymbols_FUTURE          RequestSearchSymbols_InstrumentType = 1
	RequestSearchSymbols_FUTURE_OPTION   RequestSearchSymbols_InstrumentType = 2
	RequestSearchSymbols_FUTURE_STRATEGY RequestSearchSymbols_InstrumentType = 3
	RequestSearchSymbols_EQUITY          RequestSearchSymbols_InstrumentType = 4
	RequestSearchSymbols_EQUITY_OPTION   RequestSearchSymbols_InstrumentType = 5
	RequestSearchSymbols_EQUITY_STRATEGY RequestSearchSymbols_InstrumentType = 6
	RequestSearchSymbols_INDEX           RequestSearchSymbols_InstrumentType = 7
	RequestSearchSymbols_INDEX_OPTION    RequestSearchSymbols_InstrumentType = 8
	RequestSearchSymbols_SPREAD          RequestSearchSymbols_InstrumentType = 9
	RequestSearchSymbols_SYNTHETIC       RequestSearchSymbols_InstrumentType = 10
)

// Enum value maps for RequestSearchSymbols_InstrumentType.
var (
	RequestSearchSymbols_InstrumentType_name = map[int32]string{
		1:  "FUTURE",
		2:  "FUTURE_OPTION",
		3:  "FUTURE_STRATEGY",
		4:  "EQUITY",
		5:  "EQUITY_OPTION",
		6:  "EQUITY_STRATEGY",
		7:  "INDEX",
		8:  "INDEX_OPTION",
		9:  "SPREAD",
		10: "SYNTHETIC",
	}
	RequestSearchSymbols_InstrumentType_value = map[string]int32{
		"FUTURE":          1,
		"FUTURE_OPTION":   2,
		"FUTURE_STRATEGY": 3,
		"EQUITY":          4,
		"EQUITY_OPTION":   5,
		"EQUITY_STRATEGY": 6,
		"INDEX":           7,
		"INDEX_OPTION":    8,
		"SPREAD":          9,
		"SYNTHETIC":       10,
	}
)

func (x RequestSearchSymbols_InstrumentType) Enum() *RequestSearchSymbols_InstrumentType {
	p := new(RequestSearchSymbols_InstrumentType)
	*p = x
	return p
}

func (x RequestSearchSymbols_InstrumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestSearchSymbols_InstrumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_search_symbols_proto_enumTypes[1].Descriptor()
}

func (RequestSearchSymbols_InstrumentType) Type() protoreflect.EnumType {
	return &file_request_search_symbols_proto_enumTypes[1]
}

func (x RequestSearchSymbols_InstrumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestSearchSymbols_InstrumentType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestSearchSymbols_InstrumentType(num)
	return nil
}

// Deprecated: Use RequestSearchSymbols_InstrumentType.Descriptor instead.
func (RequestSearchSymbols_InstrumentType) EnumDescriptor() ([]byte, []int) {
	return file_request_search_symbols_proto_rawDescGZIP(), []int{0, 1}
}

type RequestSearchSymbols struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId     *int32                               `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                                          // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg        []string                             `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                                                    // PB_OFFSET + MNM_USER_MSG
	SearchText     *string                              `protobuf:"bytes,120008,opt,name=search_text,json=searchText" json:"search_text,omitempty"`                                                           // PB_OFFSET + MNM_TEXT
	Exchange       *string                              `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                                                 // PB_OFFSET + MNM_EXCHANGE
	ProductCode    *string                              `protobuf:"bytes,100749,opt,name=product_code,json=productCode" json:"product_code,omitempty"`                                                        // PB_OFFSET + MNM_PRODUCT_CODE
	InstrumentType *RequestSearchSymbols_InstrumentType `protobuf:"varint,110116,opt,name=instrument_type,json=instrumentType,enum=rti.RequestSearchSymbols_InstrumentType" json:"instrument_type,omitempty"` // PB_OFFSET + MNM_INSTRUMENT_TYPE
	Pattern        *RequestSearchSymbols_Pattern        `protobuf:"varint,155008,opt,name=pattern,enum=rti.RequestSearchSymbols_Pattern" json:"pattern,omitempty"`                                            // PB_OFFSET + MNM_SEARCH_PATTERN
}

func (x *RequestSearchSymbols) Reset() {
	*x = RequestSearchSymbols{}
	mi := &file_request_search_symbols_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestSearchSymbols) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSearchSymbols) ProtoMessage() {}

func (x *RequestSearchSymbols) ProtoReflect() protoreflect.Message {
	mi := &file_request_search_symbols_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSearchSymbols.ProtoReflect.Descriptor instead.
func (*RequestSearchSymbols) Descriptor() ([]byte, []int) {
	return file_request_search_symbols_proto_rawDescGZIP(), []int{0}
}

func (x *RequestSearchSymbols) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestSearchSymbols) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestSearchSymbols) GetSearchText() string {
	if x != nil && x.SearchText != nil {
		return *x.SearchText
	}
	return ""
}

func (x *RequestSearchSymbols) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestSearchSymbols) GetProductCode() string {
	if x != nil && x.ProductCode != nil {
		return *x.ProductCode
	}
	return ""
}

func (x *RequestSearchSymbols) GetInstrumentType() RequestSearchSymbols_InstrumentType {
	if x != nil && x.InstrumentType != nil {
		return *x.InstrumentType
	}
	return RequestSearchSymbols_FUTURE
}

func (x *RequestSearchSymbols) GetPattern() RequestSearchSymbols_Pattern {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return RequestSearchSymbols_EQUALS
}

var File_request_search_symbols_proto protoreflect.FileDescriptor

var file_request_search_symbols_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x74, 0x69, 0x22, 0xa8, 0x04, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x0b,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0xc8, 0xa9, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x8d, 0x93,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xa4, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x18, 0x80, 0xbb, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x74, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x02, 0x22, 0xb0, 0x01, 0x0a, 0x0e,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x55,
	0x54, 0x55, 0x52, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x45, 0x51, 0x55, 0x49, 0x54, 0x59, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x51, 0x55, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x52, 0x41,
	0x54, 0x45, 0x47, 0x59, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10,
	0x07, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x50, 0x52, 0x45, 0x41, 0x44, 0x10, 0x09, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x59, 0x4e, 0x54, 0x48, 0x45, 0x54, 0x49, 0x43, 0x10, 0x0a, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_request_search_symbols_proto_rawDescOnce sync.Once
	file_request_search_symbols_proto_rawDescData = file_request_search_symbols_proto_rawDesc
)

func file_request_search_symbols_proto_rawDescGZIP() []byte {
	file_request_search_symbols_proto_rawDescOnce.Do(func() {
		file_request_search_symbols_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_search_symbols_proto_rawDescData)
	})
	return file_request_search_symbols_proto_rawDescData
}

var file_request_search_symbols_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_request_search_symbols_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_search_symbols_proto_goTypes = []any{
	(RequestSearchSymbols_Pattern)(0),        // 0: rti.RequestSearchSymbols.Pattern
	(RequestSearchSymbols_InstrumentType)(0), // 1: rti.RequestSearchSymbols.InstrumentType
	(*RequestSearchSymbols)(nil),             // 2: rti.RequestSearchSymbols
}
var file_request_search_symbols_proto_depIdxs = []int32{
	1, // 0: rti.RequestSearchSymbols.instrument_type:type_name -> rti.RequestSearchSymbols.InstrumentType
	0, // 1: rti.RequestSearchSymbols.pattern:type_name -> rti.RequestSearchSymbols.Pattern
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_request_search_symbols_proto_init() }
func file_request_search_symbols_proto_init() {
	if File_request_search_symbols_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_search_symbols_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_search_symbols_proto_goTypes,
		DependencyIndexes: file_request_search_symbols_proto_depIdxs,
		EnumInfos:         file_request_search_symbols_proto_enumTypes,
		MessageInfos:      file_request_search_symbols_proto_msgTypes,
	}.Build()
	File_request_search_symbols_proto = out.File
	file_request_search_symbols_proto_rawDesc = nil
	file_request_search_symbols_proto_goTypes = nil
	file_request_search_symbols_proto_depIdxs = nil
}
