// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: trade_statistics.proto

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
type TradeStatistics_PresenceBits int32

const (
	TradeStatistics_OPEN TradeStatistics_PresenceBits = 1
	TradeStatistics_HIGH TradeStatistics_PresenceBits = 2
	TradeStatistics_LOW  TradeStatistics_PresenceBits = 4
)

// Enum value maps for TradeStatistics_PresenceBits.
var (
	TradeStatistics_PresenceBits_name = map[int32]string{
		1: "OPEN",
		2: "HIGH",
		4: "LOW",
	}
	TradeStatistics_PresenceBits_value = map[string]int32{
		"OPEN": 1,
		"HIGH": 2,
		"LOW":  4,
	}
)

func (x TradeStatistics_PresenceBits) Enum() *TradeStatistics_PresenceBits {
	p := new(TradeStatistics_PresenceBits)
	*p = x
	return p
}

func (x TradeStatistics_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeStatistics_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_trade_statistics_proto_enumTypes[0].Descriptor()
}

func (TradeStatistics_PresenceBits) Type() protoreflect.EnumType {
	return &file_trade_statistics_proto_enumTypes[0]
}

func (x TradeStatistics_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TradeStatistics_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TradeStatistics_PresenceBits(num)
	return nil
}

// Deprecated: Use TradeStatistics_PresenceBits.Descriptor instead.
func (TradeStatistics_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_trade_statistics_proto_rawDescGZIP(), []int{0, 0}
}

type TradeStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId   *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`       // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol       *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                  // PB_OFFSET + MNM_SYMBOL
	Exchange     *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                              // PB_OFFSET + MNM_EXCHANGE
	PresenceBits *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"` // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits    *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`          // PB_OFFSET + MNM_DISPLAY_INDICATOR
	IsSnapshot   *bool    `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`       // PB_OFFSET + MNM_UPDATE_TYPE
	OpenPrice    *float64 `protobuf:"fixed64,100019,opt,name=open_price,json=openPrice" json:"open_price,omitempty"`         // PB_OFFSET + MNM_OPEN_PRICE
	HighPrice    *float64 `protobuf:"fixed64,100012,opt,name=high_price,json=highPrice" json:"high_price,omitempty"`         // PB_OFFSET + MNM_HIGH_PRICE
	LowPrice     *float64 `protobuf:"fixed64,100013,opt,name=low_price,json=lowPrice" json:"low_price,omitempty"`            // PB_OFFSET + MNM_LOW_PRICE
	Ssboe        *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                   // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs        *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                   // PB_OFFSET + MNM_USECS
	SourceSsboe  *int32   `protobuf:"varint,150400,opt,name=source_ssboe,json=sourceSsboe" json:"source_ssboe,omitempty"`    // PB_OFFSET + MNM_SOURCE_SSBOE
	SourceUsecs  *int32   `protobuf:"varint,150401,opt,name=source_usecs,json=sourceUsecs" json:"source_usecs,omitempty"`    // PB_OFFSET + MNM_SOURCE_USECS
	SourceNsecs  *int32   `protobuf:"varint,150404,opt,name=source_nsecs,json=sourceNsecs" json:"source_nsecs,omitempty"`    // PB_OFFSET + MNM_SOURCE_NSECS
	JopSsboe     *int32   `protobuf:"varint,150600,opt,name=jop_ssboe,json=jopSsboe" json:"jop_ssboe,omitempty"`             // PB_OFFSET + MNM_JOP_SSBOE
	JopNsecs     *int32   `protobuf:"varint,150604,opt,name=jop_nsecs,json=jopNsecs" json:"jop_nsecs,omitempty"`             // PB_OFFSET + MNM_JOP_NSECS
}

func (x *TradeStatistics) Reset() {
	*x = TradeStatistics{}
	mi := &file_trade_statistics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeStatistics) ProtoMessage() {}

func (x *TradeStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_trade_statistics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeStatistics.ProtoReflect.Descriptor instead.
func (*TradeStatistics) Descriptor() ([]byte, []int) {
	return file_trade_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *TradeStatistics) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *TradeStatistics) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *TradeStatistics) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *TradeStatistics) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *TradeStatistics) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *TradeStatistics) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *TradeStatistics) GetOpenPrice() float64 {
	if x != nil && x.OpenPrice != nil {
		return *x.OpenPrice
	}
	return 0
}

func (x *TradeStatistics) GetHighPrice() float64 {
	if x != nil && x.HighPrice != nil {
		return *x.HighPrice
	}
	return 0
}

func (x *TradeStatistics) GetLowPrice() float64 {
	if x != nil && x.LowPrice != nil {
		return *x.LowPrice
	}
	return 0
}

func (x *TradeStatistics) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *TradeStatistics) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

func (x *TradeStatistics) GetSourceSsboe() int32 {
	if x != nil && x.SourceSsboe != nil {
		return *x.SourceSsboe
	}
	return 0
}

func (x *TradeStatistics) GetSourceUsecs() int32 {
	if x != nil && x.SourceUsecs != nil {
		return *x.SourceUsecs
	}
	return 0
}

func (x *TradeStatistics) GetSourceNsecs() int32 {
	if x != nil && x.SourceNsecs != nil {
		return *x.SourceNsecs
	}
	return 0
}

func (x *TradeStatistics) GetJopSsboe() int32 {
	if x != nil && x.JopSsboe != nil {
		return *x.JopSsboe
	}
	return 0
}

func (x *TradeStatistics) GetJopNsecs() int32 {
	if x != nil && x.JopNsecs != nil {
		return *x.JopNsecs
	}
	return 0
}

var File_trade_statistics_proto protoreflect.FileDescriptor

var file_trade_statistics_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xc2, 0x04,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
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
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f,
	0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xac, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6c, 0x6f, 0x77,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xad, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f,
	0x65, 0x18, 0xd4, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65,
	0x12, 0x16, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0x80, 0x97, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x23, 0x0a,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0x81, 0x97,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x65,
	0x63, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x65,
	0x63, 0x73, 0x18, 0x84, 0x97, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x73, 0x65, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x6f, 0x70, 0x5f, 0x73,
	0x73, 0x62, 0x6f, 0x65, 0x18, 0xc8, 0x98, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f,
	0x70, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x6f, 0x70, 0x5f, 0x6e, 0x73,
	0x65, 0x63, 0x73, 0x18, 0xcc, 0x98, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x70,
	0x4e, 0x73, 0x65, 0x63, 0x73, 0x22, 0x2b, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57,
	0x10, 0x04, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_trade_statistics_proto_rawDescOnce sync.Once
	file_trade_statistics_proto_rawDescData = file_trade_statistics_proto_rawDesc
)

func file_trade_statistics_proto_rawDescGZIP() []byte {
	file_trade_statistics_proto_rawDescOnce.Do(func() {
		file_trade_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_statistics_proto_rawDescData)
	})
	return file_trade_statistics_proto_rawDescData
}

var file_trade_statistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trade_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trade_statistics_proto_goTypes = []any{
	(TradeStatistics_PresenceBits)(0), // 0: rti.TradeStatistics.PresenceBits
	(*TradeStatistics)(nil),           // 1: rti.TradeStatistics
}
var file_trade_statistics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trade_statistics_proto_init() }
func file_trade_statistics_proto_init() {
	if File_trade_statistics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trade_statistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_statistics_proto_goTypes,
		DependencyIndexes: file_trade_statistics_proto_depIdxs,
		EnumInfos:         file_trade_statistics_proto_enumTypes,
		MessageInfos:      file_trade_statistics_proto_msgTypes,
	}.Build()
	File_trade_statistics_proto = out.File
	file_trade_statistics_proto_rawDesc = nil
	file_trade_statistics_proto_goTypes = nil
	file_trade_statistics_proto_depIdxs = nil
}
