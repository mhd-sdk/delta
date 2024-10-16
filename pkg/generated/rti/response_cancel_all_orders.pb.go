// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: response_cancel_all_orders.proto

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

type ResponseCancelAllOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"` // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`           // PB_OFFSET + MNM_USER_MSG
	RpCode     []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`              // PB_OFFSET + MNM_RESPONSE_CODE
}

func (x *ResponseCancelAllOrders) Reset() {
	*x = ResponseCancelAllOrders{}
	mi := &file_response_cancel_all_orders_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseCancelAllOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCancelAllOrders) ProtoMessage() {}

func (x *ResponseCancelAllOrders) ProtoReflect() protoreflect.Message {
	mi := &file_response_cancel_all_orders_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCancelAllOrders.ProtoReflect.Descriptor instead.
func (*ResponseCancelAllOrders) Descriptor() ([]byte, []int) {
	return file_response_cancel_all_orders_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseCancelAllOrders) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseCancelAllOrders) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseCancelAllOrders) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

var File_response_cancel_all_orders_proto protoreflect.FileDescriptor

var file_response_cancel_all_orders_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x74, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_response_cancel_all_orders_proto_rawDescOnce sync.Once
	file_response_cancel_all_orders_proto_rawDescData = file_response_cancel_all_orders_proto_rawDesc
)

func file_response_cancel_all_orders_proto_rawDescGZIP() []byte {
	file_response_cancel_all_orders_proto_rawDescOnce.Do(func() {
		file_response_cancel_all_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_cancel_all_orders_proto_rawDescData)
	})
	return file_response_cancel_all_orders_proto_rawDescData
}

var file_response_cancel_all_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_cancel_all_orders_proto_goTypes = []any{
	(*ResponseCancelAllOrders)(nil), // 0: rti.ResponseCancelAllOrders
}
var file_response_cancel_all_orders_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_cancel_all_orders_proto_init() }
func file_response_cancel_all_orders_proto_init() {
	if File_response_cancel_all_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_cancel_all_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_cancel_all_orders_proto_goTypes,
		DependencyIndexes: file_response_cancel_all_orders_proto_depIdxs,
		MessageInfos:      file_response_cancel_all_orders_proto_msgTypes,
	}.Build()
	File_response_cancel_all_orders_proto = out.File
	file_response_cancel_all_orders_proto_rawDesc = nil
	file_response_cancel_all_orders_proto_goTypes = nil
	file_response_cancel_all_orders_proto_depIdxs = nil
}
