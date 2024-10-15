// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: response_easy_to_borrow_list.proto

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

type ResponseEasyToBorrowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                   // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg         []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                             // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode []string `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"` // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode          []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                // PB_OFFSET + MNM_RESPONSE_CODE
	BrokerDealer    *string  `protobuf:"bytes,154612,opt,name=broker_dealer,json=brokerDealer" json:"broker_dealer,omitempty"`              // PB_OFFSET + MNM_BROKER_DEALER
	Symbol          *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                              // PB_OFFSET + MNM_SYMBOL
	SymbolName      *string  `protobuf:"bytes,100003,opt,name=symbol_name,json=symbolName" json:"symbol_name,omitempty"`                    // PB_OFFSET + MNM_SYMBOL_NAME
	QtyAvailable    *int32   `protobuf:"varint,154613,opt,name=qty_available,json=qtyAvailable" json:"qty_available,omitempty"`             // PB_OFFSET + MNM_TOTAL_AVAILABLE_QTY
	QtyNeeded       *int32   `protobuf:"varint,154614,opt,name=qty_needed,json=qtyNeeded" json:"qty_needed,omitempty"`                      // PB_OFFSET + MNM_TOTAL_USED_QTY
	Borrowable      *bool    `protobuf:"varint,110353,opt,name=borrowable" json:"borrowable,omitempty"`                                     // PB_OFFSET + MNM_SHORT_LIST_INDICATOR
}

func (x *ResponseEasyToBorrowList) Reset() {
	*x = ResponseEasyToBorrowList{}
	mi := &file_response_easy_to_borrow_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseEasyToBorrowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseEasyToBorrowList) ProtoMessage() {}

func (x *ResponseEasyToBorrowList) ProtoReflect() protoreflect.Message {
	mi := &file_response_easy_to_borrow_list_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseEasyToBorrowList.ProtoReflect.Descriptor instead.
func (*ResponseEasyToBorrowList) Descriptor() ([]byte, []int) {
	return file_response_easy_to_borrow_list_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseEasyToBorrowList) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseEasyToBorrowList) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseEasyToBorrowList) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseEasyToBorrowList) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseEasyToBorrowList) GetBrokerDealer() string {
	if x != nil && x.BrokerDealer != nil {
		return *x.BrokerDealer
	}
	return ""
}

func (x *ResponseEasyToBorrowList) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *ResponseEasyToBorrowList) GetSymbolName() string {
	if x != nil && x.SymbolName != nil {
		return *x.SymbolName
	}
	return ""
}

func (x *ResponseEasyToBorrowList) GetQtyAvailable() int32 {
	if x != nil && x.QtyAvailable != nil {
		return *x.QtyAvailable
	}
	return 0
}

func (x *ResponseEasyToBorrowList) GetQtyNeeded() int32 {
	if x != nil && x.QtyNeeded != nil {
		return *x.QtyNeeded
	}
	return 0
}

func (x *ResponseEasyToBorrowList) GetBorrowable() bool {
	if x != nil && x.Borrowable != nil {
		return *x.Borrowable
	}
	return false
}

var File_response_easy_to_borrow_list_proto protoreflect.FileDescriptor

var file_response_easy_to_borrow_list_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x65, 0x61, 0x73, 0x79, 0x5f,
	0x74, 0x6f, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xf2, 0x02, 0x0a, 0x18, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x61, 0x73, 0x79, 0x54, 0x6f, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x9e, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x25, 0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x18, 0xf4, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0xa3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x71, 0x74, 0x79, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0xf5, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x71,
	0x74, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x71,
	0x74, 0x79, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0xf6, 0xb7, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x71, 0x74, 0x79, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0a,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x91, 0xde, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_response_easy_to_borrow_list_proto_rawDescOnce sync.Once
	file_response_easy_to_borrow_list_proto_rawDescData = file_response_easy_to_borrow_list_proto_rawDesc
)

func file_response_easy_to_borrow_list_proto_rawDescGZIP() []byte {
	file_response_easy_to_borrow_list_proto_rawDescOnce.Do(func() {
		file_response_easy_to_borrow_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_easy_to_borrow_list_proto_rawDescData)
	})
	return file_response_easy_to_borrow_list_proto_rawDescData
}

var file_response_easy_to_borrow_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_easy_to_borrow_list_proto_goTypes = []any{
	(*ResponseEasyToBorrowList)(nil), // 0: rti.ResponseEasyToBorrowList
}
var file_response_easy_to_borrow_list_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_easy_to_borrow_list_proto_init() }
func file_response_easy_to_borrow_list_proto_init() {
	if File_response_easy_to_borrow_list_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_easy_to_borrow_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_easy_to_borrow_list_proto_goTypes,
		DependencyIndexes: file_response_easy_to_borrow_list_proto_depIdxs,
		MessageInfos:      file_response_easy_to_borrow_list_proto_msgTypes,
	}.Build()
	File_response_easy_to_borrow_list_proto = out.File
	file_response_easy_to_borrow_list_proto_rawDesc = nil
	file_response_easy_to_borrow_list_proto_goTypes = nil
	file_response_easy_to_borrow_list_proto_depIdxs = nil
}
