// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/bank_dto.proto

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

type BankCardAddRequest struct {
	OwnerId              int64          `protobuf:"zigzag64,1,opt,name=OwnerId,proto3" json:"OwnerId"`
	Value                *SBankCardInfo `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BankCardAddRequest) Reset()         { *m = BankCardAddRequest{} }
func (m *BankCardAddRequest) String() string { return proto.CompactTextString(m) }
func (*BankCardAddRequest) ProtoMessage()    {}
func (*BankCardAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bank_dto_8fb36c412cb2487e, []int{0}
}
func (m *BankCardAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardAddRequest.Unmarshal(m, b)
}
func (m *BankCardAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardAddRequest.Marshal(b, m, deterministic)
}
func (dst *BankCardAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardAddRequest.Merge(dst, src)
}
func (m *BankCardAddRequest) XXX_Size() int {
	return xxx_messageInfo_BankCardAddRequest.Size(m)
}
func (m *BankCardAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardAddRequest proto.InternalMessageInfo

func (m *BankCardAddRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *BankCardAddRequest) GetValue() *SBankCardInfo {
	if m != nil {
		return m.Value
	}
	return nil
}

type BankCardRequest struct {
	OwnerId              int64    `protobuf:"varint,1,opt,name=OwnerId,proto3" json:"OwnerId"`
	BankCardNo           string   `protobuf:"bytes,2,opt,name=BankCardNo,proto3" json:"BankCardNo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankCardRequest) Reset()         { *m = BankCardRequest{} }
func (m *BankCardRequest) String() string { return proto.CompactTextString(m) }
func (*BankCardRequest) ProtoMessage()    {}
func (*BankCardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bank_dto_8fb36c412cb2487e, []int{1}
}
func (m *BankCardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardRequest.Unmarshal(m, b)
}
func (m *BankCardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardRequest.Marshal(b, m, deterministic)
}
func (dst *BankCardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardRequest.Merge(dst, src)
}
func (m *BankCardRequest) XXX_Size() int {
	return xxx_messageInfo_BankCardRequest.Size(m)
}
func (m *BankCardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardRequest proto.InternalMessageInfo

func (m *BankCardRequest) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *BankCardRequest) GetBankCardNo() string {
	if m != nil {
		return m.BankCardNo
	}
	return ""
}

type BankCardListResponse struct {
	OwnerId              int64            `protobuf:"zigzag64,1,opt,name=OwnerId,proto3" json:"OwnerId"`
	Value                []*SBankCardInfo `protobuf:"bytes,2,rep,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BankCardListResponse) Reset()         { *m = BankCardListResponse{} }
func (m *BankCardListResponse) String() string { return proto.CompactTextString(m) }
func (*BankCardListResponse) ProtoMessage()    {}
func (*BankCardListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bank_dto_8fb36c412cb2487e, []int{2}
}
func (m *BankCardListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardListResponse.Unmarshal(m, b)
}
func (m *BankCardListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardListResponse.Marshal(b, m, deterministic)
}
func (dst *BankCardListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardListResponse.Merge(dst, src)
}
func (m *BankCardListResponse) XXX_Size() int {
	return xxx_messageInfo_BankCardListResponse.Size(m)
}
func (m *BankCardListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardListResponse proto.InternalMessageInfo

func (m *BankCardListResponse) GetOwnerId() int64 {
	if m != nil {
		return m.OwnerId
	}
	return 0
}

func (m *BankCardListResponse) GetValue() []*SBankCardInfo {
	if m != nil {
		return m.Value
	}
	return nil
}

// * 会员银行卡
type SBankCardInfo struct {
	// * 银行名称
	BankName string `protobuf:"bytes,2,opt,name=BankName,proto3" json:"BankName"`
	// * 账户名
	AccountName string `protobuf:"bytes,3,opt,name=AccountName,proto3" json:"AccountName"`
	// * 账号
	AccountNo string `protobuf:"bytes,4,opt,name=AccountNo,proto3" json:"AccountNo"`
	// 银行编号
	BankId int32 `protobuf:"varint,5,opt,name=BankId,proto3" json:"BankId"`
	// 银行卡代码
	BankCode string `protobuf:"bytes,6,opt,name=BankCode,proto3" json:"BankCode"`
	// 快捷支付授权码
	AuthCode string `protobuf:"bytes,7,opt,name=AuthCode,proto3" json:"AuthCode"`
	// * 网点名称
	Network string `protobuf:"bytes,8,opt,name=Network,proto3" json:"Network"`
	// * 状态
	State int32 `protobuf:"zigzag32,9,opt,name=State,proto3" json:"State"`
	// 更新时间
	UpdateTime           int64    `protobuf:"zigzag64,10,opt,name=UpdateTime,proto3" json:"UpdateTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SBankCardInfo) Reset()         { *m = SBankCardInfo{} }
func (m *SBankCardInfo) String() string { return proto.CompactTextString(m) }
func (*SBankCardInfo) ProtoMessage()    {}
func (*SBankCardInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bank_dto_8fb36c412cb2487e, []int{3}
}
func (m *SBankCardInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SBankCardInfo.Unmarshal(m, b)
}
func (m *SBankCardInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SBankCardInfo.Marshal(b, m, deterministic)
}
func (dst *SBankCardInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SBankCardInfo.Merge(dst, src)
}
func (m *SBankCardInfo) XXX_Size() int {
	return xxx_messageInfo_SBankCardInfo.Size(m)
}
func (m *SBankCardInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SBankCardInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SBankCardInfo proto.InternalMessageInfo

func (m *SBankCardInfo) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *SBankCardInfo) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *SBankCardInfo) GetAccountNo() string {
	if m != nil {
		return m.AccountNo
	}
	return ""
}

func (m *SBankCardInfo) GetBankId() int32 {
	if m != nil {
		return m.BankId
	}
	return 0
}

func (m *SBankCardInfo) GetBankCode() string {
	if m != nil {
		return m.BankCode
	}
	return ""
}

func (m *SBankCardInfo) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func (m *SBankCardInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *SBankCardInfo) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SBankCardInfo) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*BankCardAddRequest)(nil), "BankCardAddRequest")
	proto.RegisterType((*BankCardRequest)(nil), "BankCardRequest")
	proto.RegisterType((*BankCardListResponse)(nil), "BankCardListResponse")
	proto.RegisterType((*SBankCardInfo)(nil), "SBankCardInfo")
}

func init() { proto.RegisterFile("message/bank_dto.proto", fileDescriptor_bank_dto_8fb36c412cb2487e) }

var fileDescriptor_bank_dto_8fb36c412cb2487e = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x53, 0x78, 0xcb, 0x9f, 0x21, 0xaf, 0xc6, 0x95, 0x90, 0x8d, 0x31, 0xda, 0x10, 0x0f,
	0x3d, 0x95, 0x04, 0x8f, 0x9e, 0x80, 0x13, 0xd1, 0x60, 0xb2, 0x20, 0x07, 0x2f, 0x66, 0xe9, 0x8e,
	0x40, 0xb0, 0x9d, 0xda, 0x6e, 0xc3, 0x87, 0xf0, 0x4b, 0x9b, 0x5d, 0x58, 0xc1, 0x03, 0x89, 0xa7,
	0xe6, 0xf7, 0xfc, 0xda, 0x67, 0xa7, 0x93, 0x85, 0x4e, 0x82, 0x45, 0x21, 0x97, 0xd8, 0x5b, 0xc8,
	0x74, 0xf3, 0xa6, 0x34, 0x45, 0x59, 0x4e, 0x9a, 0xba, 0x33, 0x60, 0x43, 0x99, 0x6e, 0x46, 0x32,
	0x57, 0x03, 0xa5, 0x04, 0x7e, 0x96, 0x58, 0x68, 0xc6, 0xa1, 0xfe, 0xbc, 0x4d, 0x31, 0x1f, 0x2b,
	0xee, 0x05, 0x5e, 0xc8, 0x84, 0x43, 0x76, 0x07, 0xfe, 0x5c, 0x7e, 0x94, 0xc8, 0x2b, 0x81, 0x17,
	0xb6, 0xfa, 0x67, 0xd1, 0xd4, 0x7d, 0x3e, 0x4e, 0xdf, 0x49, 0xec, 0x64, 0xf7, 0x11, 0xce, 0x5d,
	0x7c, 0xa2, 0xb2, 0x7a, 0xa8, 0xbc, 0x01, 0x70, 0x2f, 0x4f, 0xc8, 0xf6, 0x36, 0xc5, 0x51, 0xd2,
	0x9d, 0x43, 0xdb, 0xd1, 0xd3, 0xba, 0xd0, 0x02, 0x8b, 0x8c, 0xd2, 0x02, 0xff, 0x36, 0x64, 0xf5,
	0xf4, 0x90, 0x5f, 0x15, 0xf8, 0xff, 0x4b, 0xb0, 0x2b, 0x68, 0x18, 0x9e, 0xc8, 0x04, 0xf7, 0x73,
	0xfc, 0x30, 0x0b, 0xa0, 0x35, 0x88, 0x63, 0x2a, 0x53, 0x6d, 0x75, 0xd5, 0xea, 0xe3, 0x88, 0x5d,
	0x43, 0xd3, 0x21, 0xf1, 0x7f, 0xd6, 0x1f, 0x02, 0xd6, 0x81, 0x9a, 0xe9, 0x1a, 0x2b, 0xee, 0x07,
	0x5e, 0xe8, 0x8b, 0x3d, 0xb9, 0x33, 0x47, 0xa4, 0x90, 0xd7, 0x0e, 0x67, 0x1a, 0x36, 0x6e, 0x50,
	0xea, 0x95, 0x75, 0xf5, 0x9d, 0x73, 0x6c, 0xfe, 0x7e, 0x82, 0x7a, 0x4b, 0xf9, 0x86, 0x37, 0xac,
	0x72, 0xc8, 0xda, 0xe0, 0x4f, 0xb5, 0xd4, 0xc8, 0x9b, 0x81, 0x17, 0x5e, 0x88, 0x1d, 0x98, 0x2d,
	0xbf, 0x64, 0x4a, 0x6a, 0x9c, 0xad, 0x13, 0xe4, 0x60, 0x17, 0x76, 0x94, 0x0c, 0x6f, 0xe1, 0x32,
	0xa6, 0x24, 0x5a, 0xae, 0xf5, 0xaa, 0x5c, 0x44, 0x4b, 0xea, 0x53, 0x94, 0x67, 0xf1, 0x6b, 0x23,
	0xea, 0x3d, 0xd8, 0x9b, 0xb2, 0xa8, 0xd9, 0xc7, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0x46, 0x76, 0xb4, 0x4a, 0x02, 0x00, 0x00,
}
