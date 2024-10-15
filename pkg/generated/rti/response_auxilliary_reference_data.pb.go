// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.5
// source: response_auxilliary_reference_data.proto

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

// bit constants are defined using enum
type ResponseAuxilliaryReferenceData_PresenceBits int32

const (
	ResponseAuxilliaryReferenceData_SETTLEMENT_METHOD   ResponseAuxilliaryReferenceData_PresenceBits = 1
	ResponseAuxilliaryReferenceData_FIRST_NOTICE_DATE   ResponseAuxilliaryReferenceData_PresenceBits = 2
	ResponseAuxilliaryReferenceData_LAST_NOTICE_DATE    ResponseAuxilliaryReferenceData_PresenceBits = 4
	ResponseAuxilliaryReferenceData_FIRST_TRADING_DATE  ResponseAuxilliaryReferenceData_PresenceBits = 8
	ResponseAuxilliaryReferenceData_LAST_TRADING_DATE   ResponseAuxilliaryReferenceData_PresenceBits = 16
	ResponseAuxilliaryReferenceData_FIRST_DELIVERY_DATE ResponseAuxilliaryReferenceData_PresenceBits = 32
	ResponseAuxilliaryReferenceData_LAST_DELIVERY_DATE  ResponseAuxilliaryReferenceData_PresenceBits = 64
	ResponseAuxilliaryReferenceData_FIRST_POSITION_DATE ResponseAuxilliaryReferenceData_PresenceBits = 128
	ResponseAuxilliaryReferenceData_LAST_POSITION_DATE  ResponseAuxilliaryReferenceData_PresenceBits = 256
	ResponseAuxilliaryReferenceData_UNIT_OF_MEASURE     ResponseAuxilliaryReferenceData_PresenceBits = 512
	ResponseAuxilliaryReferenceData_UNIT_OF_MEASURE_QTY ResponseAuxilliaryReferenceData_PresenceBits = 1024
)

// Enum value maps for ResponseAuxilliaryReferenceData_PresenceBits.
var (
	ResponseAuxilliaryReferenceData_PresenceBits_name = map[int32]string{
		1:    "SETTLEMENT_METHOD",
		2:    "FIRST_NOTICE_DATE",
		4:    "LAST_NOTICE_DATE",
		8:    "FIRST_TRADING_DATE",
		16:   "LAST_TRADING_DATE",
		32:   "FIRST_DELIVERY_DATE",
		64:   "LAST_DELIVERY_DATE",
		128:  "FIRST_POSITION_DATE",
		256:  "LAST_POSITION_DATE",
		512:  "UNIT_OF_MEASURE",
		1024: "UNIT_OF_MEASURE_QTY",
	}
	ResponseAuxilliaryReferenceData_PresenceBits_value = map[string]int32{
		"SETTLEMENT_METHOD":   1,
		"FIRST_NOTICE_DATE":   2,
		"LAST_NOTICE_DATE":    4,
		"FIRST_TRADING_DATE":  8,
		"LAST_TRADING_DATE":   16,
		"FIRST_DELIVERY_DATE": 32,
		"LAST_DELIVERY_DATE":  64,
		"FIRST_POSITION_DATE": 128,
		"LAST_POSITION_DATE":  256,
		"UNIT_OF_MEASURE":     512,
		"UNIT_OF_MEASURE_QTY": 1024,
	}
)

func (x ResponseAuxilliaryReferenceData_PresenceBits) Enum() *ResponseAuxilliaryReferenceData_PresenceBits {
	p := new(ResponseAuxilliaryReferenceData_PresenceBits)
	*p = x
	return p
}

func (x ResponseAuxilliaryReferenceData_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseAuxilliaryReferenceData_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_response_auxilliary_reference_data_proto_enumTypes[0].Descriptor()
}

func (ResponseAuxilliaryReferenceData_PresenceBits) Type() protoreflect.EnumType {
	return &file_response_auxilliary_reference_data_proto_enumTypes[0]
}

func (x ResponseAuxilliaryReferenceData_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseAuxilliaryReferenceData_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseAuxilliaryReferenceData_PresenceBits(num)
	return nil
}

// Deprecated: Use ResponseAuxilliaryReferenceData_PresenceBits.Descriptor instead.
func (ResponseAuxilliaryReferenceData_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_response_auxilliary_reference_data_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseAuxilliaryReferenceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId        *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                        // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg           []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                  // PB_OFFSET + MNM_USER_MSG
	RpCode            []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                     // PB_OFFSET + MNM_RESPONSE_CODE
	PresenceBits      *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`                  // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits         *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`                           // PB_OFFSET + MNM_DISPLAY_INDICATOR
	Symbol            *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                   // PB_OFFSET + MNM_SYMBOL
	Exchange          *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                               // PB_OFFSET + MNM_EXCHANGE
	SettlementMethod  *string  `protobuf:"bytes,153294,opt,name=settlement_method,json=settlementMethod" json:"settlement_method,omitempty"`       // PB_OFFSET + MNM_SETTLEMENT_METHOD
	FirstNoticeDate   *string  `protobuf:"bytes,154932,opt,name=first_notice_date,json=firstNoticeDate" json:"first_notice_date,omitempty"`        // PB_OFFSET + MNM_FIRST_NOTICE_DATE
	LastNoticeDate    *string  `protobuf:"bytes,154933,opt,name=last_notice_date,json=lastNoticeDate" json:"last_notice_date,omitempty"`           // PB_OFFSET + MNM_LAST_NOTICE_DATE
	FirstTradingDate  *string  `protobuf:"bytes,154996,opt,name=first_trading_date,json=firstTradingDate" json:"first_trading_date,omitempty"`     // PB_OFFSET + MNM_FIRST_TRADING_DATE
	LastTradingDate   *string  `protobuf:"bytes,154236,opt,name=last_trading_date,json=lastTradingDate" json:"last_trading_date,omitempty"`        // PB_OFFSET + MNM_LAST_TRADING_DATE
	FirstDeliveryDate *string  `protobuf:"bytes,154994,opt,name=first_delivery_date,json=firstDeliveryDate" json:"first_delivery_date,omitempty"`  // PB_OFFSET + MNM_FIRST_DELIVERY_DATE
	LastDeliveryDate  *string  `protobuf:"bytes,154995,opt,name=last_delivery_date,json=lastDeliveryDate" json:"last_delivery_date,omitempty"`     // PB_OFFSET + MNM_LAST_DELIVERY_DATE
	FirstPositionDate *string  `protobuf:"bytes,154997,opt,name=first_position_date,json=firstPositionDate" json:"first_position_date,omitempty"`  // PB_OFFSET + MNM_FIRST_POSITION_DATE
	LastPositionDate  *string  `protobuf:"bytes,154998,opt,name=last_position_date,json=lastPositionDate" json:"last_position_date,omitempty"`     // PB_OFFSET + MNM_LAST_POSITION_DATE
	UnitOfMeasure     *string  `protobuf:"bytes,157023,opt,name=unit_of_measure,json=unitOfMeasure" json:"unit_of_measure,omitempty"`              // PB_OFFSET + MNM_UNIT_OF_MEASURE
	UnitOfMeasureQty  *float64 `protobuf:"fixed64,157024,opt,name=unit_of_measure_qty,json=unitOfMeasureQty" json:"unit_of_measure_qty,omitempty"` // PB_OFFSET + MNM_UNIT_OF_MEASURE_QTY
}

func (x *ResponseAuxilliaryReferenceData) Reset() {
	*x = ResponseAuxilliaryReferenceData{}
	mi := &file_response_auxilliary_reference_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseAuxilliaryReferenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAuxilliaryReferenceData) ProtoMessage() {}

func (x *ResponseAuxilliaryReferenceData) ProtoReflect() protoreflect.Message {
	mi := &file_response_auxilliary_reference_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAuxilliaryReferenceData.ProtoReflect.Descriptor instead.
func (*ResponseAuxilliaryReferenceData) Descriptor() ([]byte, []int) {
	return file_response_auxilliary_reference_data_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseAuxilliaryReferenceData) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseAuxilliaryReferenceData) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseAuxilliaryReferenceData) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseAuxilliaryReferenceData) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *ResponseAuxilliaryReferenceData) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *ResponseAuxilliaryReferenceData) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetSettlementMethod() string {
	if x != nil && x.SettlementMethod != nil {
		return *x.SettlementMethod
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetFirstNoticeDate() string {
	if x != nil && x.FirstNoticeDate != nil {
		return *x.FirstNoticeDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetLastNoticeDate() string {
	if x != nil && x.LastNoticeDate != nil {
		return *x.LastNoticeDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetFirstTradingDate() string {
	if x != nil && x.FirstTradingDate != nil {
		return *x.FirstTradingDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetLastTradingDate() string {
	if x != nil && x.LastTradingDate != nil {
		return *x.LastTradingDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetFirstDeliveryDate() string {
	if x != nil && x.FirstDeliveryDate != nil {
		return *x.FirstDeliveryDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetLastDeliveryDate() string {
	if x != nil && x.LastDeliveryDate != nil {
		return *x.LastDeliveryDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetFirstPositionDate() string {
	if x != nil && x.FirstPositionDate != nil {
		return *x.FirstPositionDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetLastPositionDate() string {
	if x != nil && x.LastPositionDate != nil {
		return *x.LastPositionDate
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetUnitOfMeasure() string {
	if x != nil && x.UnitOfMeasure != nil {
		return *x.UnitOfMeasure
	}
	return ""
}

func (x *ResponseAuxilliaryReferenceData) GetUnitOfMeasureQty() float64 {
	if x != nil && x.UnitOfMeasureQty != nil {
		return *x.UnitOfMeasureQty
	}
	return 0
}

var File_response_auxilliary_reference_data_proto protoreflect.FileDescriptor

var file_response_auxilliary_reference_data_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x61, 0x75, 0x78, 0x69, 0x6c,
	0x6c, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22,
	0x9a, 0x08, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x78, 0x69,
	0x6c, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e,
	0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18,
	0x92, 0x8d, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x62,
	0x69, 0x74, 0x73, 0x18, 0xcb, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x42, 0x69, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x11, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0xce, 0xad, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x0a,
	0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0xb4, 0xba, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0xb5, 0xba, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xf4, 0xba,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xfc, 0xb4, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xf2, 0xba, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xf3, 0xba,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xf5,
	0xba, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0xf6, 0xba, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0xdf, 0xca, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x71, 0x74, 0x79, 0x18, 0xe0, 0xca, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x51, 0x74, 0x79, 0x22, 0x95, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49,
	0x43, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x52,
	0x53, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x08, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x52, 0x53,
	0x54, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x20, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x40, 0x12, 0x18, 0x0a, 0x13, 0x46, 0x49, 0x52,
	0x53, 0x54, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x10, 0x80, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x80, 0x02, 0x12, 0x14, 0x0a, 0x0f,
	0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x4d, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x10,
	0x80, 0x04, 0x12, 0x18, 0x0a, 0x13, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x4d, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x51, 0x54, 0x59, 0x10, 0x80, 0x08, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_response_auxilliary_reference_data_proto_rawDescOnce sync.Once
	file_response_auxilliary_reference_data_proto_rawDescData = file_response_auxilliary_reference_data_proto_rawDesc
)

func file_response_auxilliary_reference_data_proto_rawDescGZIP() []byte {
	file_response_auxilliary_reference_data_proto_rawDescOnce.Do(func() {
		file_response_auxilliary_reference_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_auxilliary_reference_data_proto_rawDescData)
	})
	return file_response_auxilliary_reference_data_proto_rawDescData
}

var file_response_auxilliary_reference_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_auxilliary_reference_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_auxilliary_reference_data_proto_goTypes = []any{
	(ResponseAuxilliaryReferenceData_PresenceBits)(0), // 0: rti.ResponseAuxilliaryReferenceData.PresenceBits
	(*ResponseAuxilliaryReferenceData)(nil),           // 1: rti.ResponseAuxilliaryReferenceData
}
var file_response_auxilliary_reference_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_auxilliary_reference_data_proto_init() }
func file_response_auxilliary_reference_data_proto_init() {
	if File_response_auxilliary_reference_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_auxilliary_reference_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_auxilliary_reference_data_proto_goTypes,
		DependencyIndexes: file_response_auxilliary_reference_data_proto_depIdxs,
		EnumInfos:         file_response_auxilliary_reference_data_proto_enumTypes,
		MessageInfos:      file_response_auxilliary_reference_data_proto_msgTypes,
	}.Build()
	File_response_auxilliary_reference_data_proto = out.File
	file_response_auxilliary_reference_data_proto_rawDesc = nil
	file_response_auxilliary_reference_data_proto_goTypes = nil
	file_response_auxilliary_reference_data_proto_depIdxs = nil
}
