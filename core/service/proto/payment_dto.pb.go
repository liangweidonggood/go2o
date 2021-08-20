// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/payment_dto.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PaymentAccountType int32

const (
	// 余额支付
	PaymentAccountType_PA_Balance PaymentAccountType = 0
	// 钱包
	PaymentAccountType_PA_Wallet PaymentAccountType = 1
	// 快捷支付
	PaymentAccountType_PA_QuickPay PaymentAccountType = 2
)

var PaymentAccountType_name = map[int32]string{
	0: "PA_Balance",
	1: "PA_Wallet",
	2: "PA_QuickPay",
}
var PaymentAccountType_value = map[string]int32{
	"PA_Balance":  0,
	"PA_Wallet":   1,
	"PA_QuickPay": 2,
}

func (x PaymentAccountType) String() string {
	return proto.EnumName(PaymentAccountType_name, int32(x))
}
func (PaymentAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payment_dto_db0f6f3aadf5c05e, []int{0}
}

type PayPriorityRequest struct {
	OwnerId              int64              `protobuf:"zigzag64,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	Account              PaymentAccountType `protobuf:"varint,2,opt,name=account,proto3,enum=PaymentAccountType" json:"account"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PayPriorityRequest) Reset()         { *m = PayPriorityRequest{} }
func (m *PayPriorityRequest) String() string { return proto.CompactTextString(m) }
func (*PayPriorityRequest) ProtoMessage()    {}
func (*PayPriorityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_dto_db0f6f3aadf5c05e, []int{0}
}
func (m *PayPriorityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayPriorityRequest.Unmarshal(m, b)
}
func (m *PayPriorityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayPriorityRequest.Marshal(b, m, deterministic)
}
func (dst *PayPriorityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayPriorityRequest.Merge(dst, src)
}
func (m *PayPriorityRequest) XXX_Size() int {
	return xxx_messageInfo_PayPriorityRequest.Size(m)
}
func (m *PayPriorityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayPriorityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayPriorityRequest proto.InternalMessageInfo

func (m *PayPriorityRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *PayPriorityRequest) GetAccount() PaymentAccountType {
	if m != nil {
		return m.Account
	}
	return PaymentAccountType_PA_Balance
}

func init() {
	proto.RegisterType((*PayPriorityRequest)(nil), "PayPriorityRequest")
	proto.RegisterEnum("PaymentAccountType", PaymentAccountType_name, PaymentAccountType_value)
}

func init() {
	proto.RegisterFile("message/payment_dto.proto", fileDescriptor_payment_dto_db0f6f3aadf5c05e)
}

var fileDescriptor_payment_dto_db0f6f3aadf5c05e = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x0f, 0xb6, 0x8e, 0x58, 0xc3, 0xf6, 0xd2, 0x9e, 0x2c, 0x9e, 0x8a, 0xe0, 0x16,
	0xea, 0xd1, 0x53, 0x8a, 0x17, 0x6f, 0x6b, 0x10, 0x04, 0x0f, 0x2e, 0xd3, 0xcd, 0x10, 0x17, 0x93,
	0x4c, 0xdc, 0x4c, 0x90, 0xfd, 0xf6, 0xc2, 0x06, 0x2f, 0x9e, 0x1e, 0xef, 0x0f, 0x3f, 0x78, 0xb0,
	0x69, 0x69, 0x18, 0xb0, 0xa6, 0x7d, 0x8f, 0xb1, 0xa5, 0x4e, 0x6c, 0x25, 0xac, 0xfb, 0xc0, 0xc2,
	0xb7, 0x1f, 0xa0, 0x0c, 0x46, 0x13, 0x3c, 0x07, 0x2f, 0xb1, 0xa4, 0xef, 0x91, 0x06, 0x51, 0x1b,
	0x58, 0xf0, 0x4f, 0x47, 0xc1, 0xfa, 0x6a, 0x9d, 0x6d, 0xb3, 0x9d, 0x2a, 0xe7, 0xc9, 0x3f, 0x57,
	0xea, 0x1e, 0xe6, 0xe8, 0x1c, 0x8f, 0x9d, 0xac, 0x67, 0xdb, 0x6c, 0xb7, 0x3c, 0xac, 0xb4, 0x99,
	0xa8, 0xc5, 0x14, 0xbf, 0xc6, 0x9e, 0xca, 0xbf, 0xcd, 0xdd, 0x53, 0xe2, 0xff, 0xab, 0xd5, 0x12,
	0xc0, 0x14, 0xf6, 0x88, 0x0d, 0x76, 0x8e, 0xf2, 0x33, 0x75, 0x05, 0x17, 0xa6, 0xb0, 0x6f, 0xd8,
	0x34, 0x24, 0x79, 0xa6, 0xae, 0xe1, 0xd2, 0x14, 0xf6, 0x65, 0xf4, 0xee, 0xcb, 0x60, 0xcc, 0x67,
	0xc7, 0x1b, 0x58, 0x39, 0x6e, 0x75, 0xed, 0xe5, 0x73, 0x3c, 0xe9, 0x9a, 0x0f, 0xac, 0x43, 0xef,
	0xde, 0x17, 0x7a, 0xff, 0x98, 0x6e, 0x9c, 0xce, 0x93, 0x3c, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0xa0, 0x29, 0xaf, 0xea, 0x00, 0x00, 0x00,
}
