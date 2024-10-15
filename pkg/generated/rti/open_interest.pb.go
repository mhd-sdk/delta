// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: open_interest.proto

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

type OpenInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId   *int32  `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`       // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol       *string `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                  // PB_OFFSET + MNM_SYMBOL
	Exchange     *string `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                              // PB_OFFSET + MNM_EXCHANGE
	IsSnapshot   *bool   `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`       // PB_OFFSET + MNM_UPDATE_TYPE
	ShouldClear  *bool   `protobuf:"varint,154571,opt,name=should_clear,json=shouldClear" json:"should_clear,omitempty"`    // PB_OFFSET + MNM_DISPLAY_INDICATOR
	OpenInterest *uint64 `protobuf:"varint,100064,opt,name=open_interest,json=openInterest" json:"open_interest,omitempty"` // PB_OFFSET + MNM_OPEN_INTEREST
	Ssboe        *int32  `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                   // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs        *int32  `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                   // PB_OFFSET + MNM_USECS
}

func (x *OpenInterest) Reset() {
	*x = OpenInterest{}
	mi := &file_open_interest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenInterest) ProtoMessage() {}

func (x *OpenInterest) ProtoReflect() protoreflect.Message {
	mi := &file_open_interest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenInterest.ProtoReflect.Descriptor instead.
func (*OpenInterest) Descriptor() ([]byte, []int) {
	return file_open_interest_proto_rawDescGZIP(), []int{0}
}

func (x *OpenInterest) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *OpenInterest) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *OpenInterest) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *OpenInterest) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *OpenInterest) GetShouldClear() bool {
	if x != nil && x.ShouldClear != nil {
		return *x.ShouldClear
	}
	return false
}

func (x *OpenInterest) GetOpenInterest() uint64 {
	if x != nil && x.OpenInterest != nil {
		return *x.OpenInterest
	}
	return 0
}

func (x *OpenInterest) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *OpenInterest) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_open_interest_proto protoreflect.FileDescriptor

var file_open_interest_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x88, 0x02, 0x0a, 0x0c, 0x4f,
	0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0xa9, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x18, 0xcb, 0xb7, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x25,
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18,
	0xe0, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd4,
	0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x16, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x63, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_open_interest_proto_rawDescOnce sync.Once
	file_open_interest_proto_rawDescData = file_open_interest_proto_rawDesc
)

func file_open_interest_proto_rawDescGZIP() []byte {
	file_open_interest_proto_rawDescOnce.Do(func() {
		file_open_interest_proto_rawDescData = protoimpl.X.CompressGZIP(file_open_interest_proto_rawDescData)
	})
	return file_open_interest_proto_rawDescData
}

var file_open_interest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_open_interest_proto_goTypes = []any{
	(*OpenInterest)(nil), // 0: rti.OpenInterest
}
var file_open_interest_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_open_interest_proto_init() }
func file_open_interest_proto_init() {
	if File_open_interest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_open_interest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_open_interest_proto_goTypes,
		DependencyIndexes: file_open_interest_proto_depIdxs,
		MessageInfos:      file_open_interest_proto_msgTypes,
	}.Build()
	File_open_interest_proto = out.File
	file_open_interest_proto_rawDesc = nil
	file_open_interest_proto_goTypes = nil
	file_open_interest_proto_depIdxs = nil
}
