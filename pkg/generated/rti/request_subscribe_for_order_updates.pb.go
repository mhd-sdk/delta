// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: request_subscribe_for_order_updates.proto

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

type RequestSubscribeForOrderUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"` // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`           // PB_OFFSET + MNM_USER_MSG
	FcmId      *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                 // PB_OFFSET + MNM_FCM_ID
	IbId       *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                    // PB_OFFSET + MNM_IB_ID
	AccountId  *string  `protobuf:"bytes,154008,opt,name=account_id,json=accountId" json:"account_id,omitempty"`     // PB_OFFSET + MNM_ACCOUNT_ID
}

func (x *RequestSubscribeForOrderUpdates) Reset() {
	*x = RequestSubscribeForOrderUpdates{}
	mi := &file_request_subscribe_for_order_updates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestSubscribeForOrderUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSubscribeForOrderUpdates) ProtoMessage() {}

func (x *RequestSubscribeForOrderUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_request_subscribe_for_order_updates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSubscribeForOrderUpdates.ProtoReflect.Descriptor instead.
func (*RequestSubscribeForOrderUpdates) Descriptor() ([]byte, []int) {
	return file_request_subscribe_for_order_updates_proto_rawDescGZIP(), []int{0}
}

func (x *RequestSubscribeForOrderUpdates) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestSubscribeForOrderUpdates) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestSubscribeForOrderUpdates) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *RequestSubscribeForOrderUpdates) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *RequestSubscribeForOrderUpdates) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

var File_request_subscribe_for_order_updates_proto protoreflect.FileDescriptor

var file_request_subscribe_for_order_updates_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69,
	0x22, 0xb2, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d,
	0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x05, 0x69, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x62, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x98, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_request_subscribe_for_order_updates_proto_rawDescOnce sync.Once
	file_request_subscribe_for_order_updates_proto_rawDescData = file_request_subscribe_for_order_updates_proto_rawDesc
)

func file_request_subscribe_for_order_updates_proto_rawDescGZIP() []byte {
	file_request_subscribe_for_order_updates_proto_rawDescOnce.Do(func() {
		file_request_subscribe_for_order_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_subscribe_for_order_updates_proto_rawDescData)
	})
	return file_request_subscribe_for_order_updates_proto_rawDescData
}

var file_request_subscribe_for_order_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_subscribe_for_order_updates_proto_goTypes = []any{
	(*RequestSubscribeForOrderUpdates)(nil), // 0: rti.RequestSubscribeForOrderUpdates
}
var file_request_subscribe_for_order_updates_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_subscribe_for_order_updates_proto_init() }
func file_request_subscribe_for_order_updates_proto_init() {
	if File_request_subscribe_for_order_updates_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_request_subscribe_for_order_updates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_subscribe_for_order_updates_proto_goTypes,
		DependencyIndexes: file_request_subscribe_for_order_updates_proto_depIdxs,
		MessageInfos:      file_request_subscribe_for_order_updates_proto_msgTypes,
	}.Build()
	File_request_subscribe_for_order_updates_proto = out.File
	file_request_subscribe_for_order_updates_proto_rawDesc = nil
	file_request_subscribe_for_order_updates_proto_goTypes = nil
	file_request_subscribe_for_order_updates_proto_depIdxs = nil
}
