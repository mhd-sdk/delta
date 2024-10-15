// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: indicator_prices.proto

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

// below enum is just for reference only, not used in this message
type IndicatorPrices_PresenceBits int32

const (
	IndicatorPrices_OPENING_INDICATOR IndicatorPrices_PresenceBits = 1
	IndicatorPrices_CLOSING_INDICATOR IndicatorPrices_PresenceBits = 2
)

// Enum value maps for IndicatorPrices_PresenceBits.
var (
	IndicatorPrices_PresenceBits_name = map[int32]string{
		1: "OPENING_INDICATOR",
		2: "CLOSING_INDICATOR",
	}
	IndicatorPrices_PresenceBits_value = map[string]int32{
		"OPENING_INDICATOR": 1,
		"CLOSING_INDICATOR": 2,
	}
)

func (x IndicatorPrices_PresenceBits) Enum() *IndicatorPrices_PresenceBits {
	p := new(IndicatorPrices_PresenceBits)
	*p = x
	return p
}

func (x IndicatorPrices_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndicatorPrices_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_indicator_prices_proto_enumTypes[0].Descriptor()
}

func (IndicatorPrices_PresenceBits) Type() protoreflect.EnumType {
	return &file_indicator_prices_proto_enumTypes[0]
}

func (x IndicatorPrices_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *IndicatorPrices_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = IndicatorPrices_PresenceBits(num)
	return nil
}

// Deprecated: Use IndicatorPrices_PresenceBits.Descriptor instead.
func (IndicatorPrices_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_indicator_prices_proto_rawDescGZIP(), []int{0, 0}
}

type IndicatorPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId       *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                    // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol           *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                               // PB_OFFSET + MNM_SYMBOL
	Exchange         *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                           // PB_OFFSET + MNM_EXCHANGE
	PresenceBits     *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`              // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits        *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`                       // PB_OFFSET + MNM_DISPLAY_INDICATOR
	IsSnapshot       *bool    `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`                    // PB_OFFSET + MNM_UPDATE_TYPE
	OpeningIndicator *float64 `protobuf:"fixed64,154522,opt,name=opening_indicator,json=openingIndicator" json:"opening_indicator,omitempty"` // PB_OFFSET + MNM_OPENING_INDICATOR
	ClosingIndicator *float64 `protobuf:"fixed64,154064,opt,name=closing_indicator,json=closingIndicator" json:"closing_indicator,omitempty"` // PB_OFFSET + MNM_CLOSING_INDICATOR
	Ssboe            *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                                // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs            *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                                // PB_OFFSET + MNM_USECS
}

func (x *IndicatorPrices) Reset() {
	*x = IndicatorPrices{}
	mi := &file_indicator_prices_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndicatorPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndicatorPrices) ProtoMessage() {}

func (x *IndicatorPrices) ProtoReflect() protoreflect.Message {
	mi := &file_indicator_prices_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndicatorPrices.ProtoReflect.Descriptor instead.
func (*IndicatorPrices) Descriptor() ([]byte, []int) {
	return file_indicator_prices_proto_rawDescGZIP(), []int{0}
}

func (x *IndicatorPrices) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *IndicatorPrices) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *IndicatorPrices) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *IndicatorPrices) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *IndicatorPrices) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *IndicatorPrices) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *IndicatorPrices) GetOpeningIndicator() float64 {
	if x != nil && x.OpeningIndicator != nil {
		return *x.OpeningIndicator
	}
	return 0
}

func (x *IndicatorPrices) GetClosingIndicator() float64 {
	if x != nil && x.ClosingIndicator != nil {
		return *x.ClosingIndicator
	}
	return 0
}

func (x *IndicatorPrices) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *IndicatorPrices) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_indicator_prices_proto protoreflect.FileDescriptor

var file_indicator_prices_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xa3, 0x03,
	0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94,
	0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x92, 0x8d,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x69, 0x74,
	0x73, 0x18, 0xcb, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72,
	0x42, 0x69, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0xa9, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2d, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x9a, 0xb7, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x11, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0xd0, 0xb3, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd4,
	0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x16, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x63, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54, 0x4f,
	0x52, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_indicator_prices_proto_rawDescOnce sync.Once
	file_indicator_prices_proto_rawDescData = file_indicator_prices_proto_rawDesc
)

func file_indicator_prices_proto_rawDescGZIP() []byte {
	file_indicator_prices_proto_rawDescOnce.Do(func() {
		file_indicator_prices_proto_rawDescData = protoimpl.X.CompressGZIP(file_indicator_prices_proto_rawDescData)
	})
	return file_indicator_prices_proto_rawDescData
}

var file_indicator_prices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_indicator_prices_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_indicator_prices_proto_goTypes = []any{
	(IndicatorPrices_PresenceBits)(0), // 0: rti.IndicatorPrices.PresenceBits
	(*IndicatorPrices)(nil),           // 1: rti.IndicatorPrices
}
var file_indicator_prices_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_indicator_prices_proto_init() }
func file_indicator_prices_proto_init() {
	if File_indicator_prices_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_indicator_prices_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indicator_prices_proto_goTypes,
		DependencyIndexes: file_indicator_prices_proto_depIdxs,
		EnumInfos:         file_indicator_prices_proto_enumTypes,
		MessageInfos:      file_indicator_prices_proto_msgTypes,
	}.Build()
	File_indicator_prices_proto = out.File
	file_indicator_prices_proto_rawDesc = nil
	file_indicator_prices_proto_goTypes = nil
	file_indicator_prices_proto_depIdxs = nil
}
