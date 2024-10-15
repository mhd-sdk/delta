// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: request_volume_profile_minute_bars.proto

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

type RequestVolumeProfileMinuteBars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId    *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg       []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	Symbol        *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`
	Exchange      *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`
	BarTypePeriod *int32   `protobuf:"varint,119215,opt,name=bar_type_period,json=barTypePeriod" json:"bar_type_period,omitempty"`
	StartIndex    *int32   `protobuf:"varint,153002,opt,name=start_index,json=startIndex" json:"start_index,omitempty"`
	FinishIndex   *int32   `protobuf:"varint,153003,opt,name=finish_index,json=finishIndex" json:"finish_index,omitempty"`
	UserMaxCount  *int32   `protobuf:"varint,154020,opt,name=user_max_count,json=userMaxCount" json:"user_max_count,omitempty"`
	ResumeBars    *bool    `protobuf:"varint,153642,opt,name=resume_bars,json=resumeBars" json:"resume_bars,omitempty"`
}

func (x *RequestVolumeProfileMinuteBars) Reset() {
	*x = RequestVolumeProfileMinuteBars{}
	mi := &file_request_volume_profile_minute_bars_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVolumeProfileMinuteBars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVolumeProfileMinuteBars) ProtoMessage() {}

func (x *RequestVolumeProfileMinuteBars) ProtoReflect() protoreflect.Message {
	mi := &file_request_volume_profile_minute_bars_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVolumeProfileMinuteBars.ProtoReflect.Descriptor instead.
func (*RequestVolumeProfileMinuteBars) Descriptor() ([]byte, []int) {
	return file_request_volume_profile_minute_bars_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVolumeProfileMinuteBars) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestVolumeProfileMinuteBars) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestVolumeProfileMinuteBars) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *RequestVolumeProfileMinuteBars) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestVolumeProfileMinuteBars) GetBarTypePeriod() int32 {
	if x != nil && x.BarTypePeriod != nil {
		return *x.BarTypePeriod
	}
	return 0
}

func (x *RequestVolumeProfileMinuteBars) GetStartIndex() int32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

func (x *RequestVolumeProfileMinuteBars) GetFinishIndex() int32 {
	if x != nil && x.FinishIndex != nil {
		return *x.FinishIndex
	}
	return 0
}

func (x *RequestVolumeProfileMinuteBars) GetUserMaxCount() int32 {
	if x != nil && x.UserMaxCount != nil {
		return *x.UserMaxCount
	}
	return 0
}

func (x *RequestVolumeProfileMinuteBars) GetResumeBars() bool {
	if x != nil && x.ResumeBars != nil {
		return *x.ResumeBars
	}
	return false
}

var File_request_volume_profile_minute_bars_proto protoreflect.FileDescriptor

var file_request_volume_profile_minute_bars_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x5f,
	0x62, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22,
	0xd5, 0x02, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x42, 0x61,
	0x72, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0xaf, 0xa3,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0xaa, 0xab, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0xab, 0xab, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xa4,
	0xb3, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x78, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x62,
	0x61, 0x72, 0x73, 0x18, 0xaa, 0xb0, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x42, 0x61, 0x72, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_request_volume_profile_minute_bars_proto_rawDescOnce sync.Once
	file_request_volume_profile_minute_bars_proto_rawDescData = file_request_volume_profile_minute_bars_proto_rawDesc
)

func file_request_volume_profile_minute_bars_proto_rawDescGZIP() []byte {
	file_request_volume_profile_minute_bars_proto_rawDescOnce.Do(func() {
		file_request_volume_profile_minute_bars_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_volume_profile_minute_bars_proto_rawDescData)
	})
	return file_request_volume_profile_minute_bars_proto_rawDescData
}

var file_request_volume_profile_minute_bars_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_volume_profile_minute_bars_proto_goTypes = []any{
	(*RequestVolumeProfileMinuteBars)(nil), // 0: rti.RequestVolumeProfileMinuteBars
}
var file_request_volume_profile_minute_bars_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_volume_profile_minute_bars_proto_init() }
func file_request_volume_profile_minute_bars_proto_init() {
	if File_request_volume_profile_minute_bars_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_volume_profile_minute_bars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_volume_profile_minute_bars_proto_goTypes,
		DependencyIndexes: file_request_volume_profile_minute_bars_proto_depIdxs,
		MessageInfos:      file_request_volume_profile_minute_bars_proto_msgTypes,
	}.Build()
	File_request_volume_profile_minute_bars_proto = out.File
	file_request_volume_profile_minute_bars_proto_rawDesc = nil
	file_request_volume_profile_minute_bars_proto_goTypes = nil
	file_request_volume_profile_minute_bars_proto_depIdxs = nil
}
