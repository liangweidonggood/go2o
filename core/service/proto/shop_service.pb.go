// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TurnShopRequest struct {
	ShopId               int64    `protobuf:"zigzag64,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`
	On                   bool     `protobuf:"varint,2,opt,name=on,proto3" json:"on"`
	Reason               string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TurnShopRequest) Reset()         { *m = TurnShopRequest{} }
func (m *TurnShopRequest) String() string { return proto.CompactTextString(m) }
func (*TurnShopRequest) ProtoMessage()    {}
func (*TurnShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{0}
}
func (m *TurnShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnShopRequest.Unmarshal(m, b)
}
func (m *TurnShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnShopRequest.Marshal(b, m, deterministic)
}
func (dst *TurnShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnShopRequest.Merge(dst, src)
}
func (m *TurnShopRequest) XXX_Size() int {
	return xxx_messageInfo_TurnShopRequest.Size(m)
}
func (m *TurnShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnShopRequest proto.InternalMessageInfo

func (m *TurnShopRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *TurnShopRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *TurnShopRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// 店铺
type SShop struct {
	// * 店铺编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// * 商户编号
	MerchantId int64 `protobuf:"varint,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id"`
	// * 店铺名称
	ShopName string `protobuf:"bytes,3,opt,name=ShopName,proto3" json:"ShopName"`
	// 店铺标题
	ShopTitle string `protobuf:"bytes,4,opt,name=ShopTitle,proto3" json:"ShopTitle"`
	// 店铺公告
	ShopNotice string `protobuf:"bytes,5,opt,name=ShopNotice,proto3" json:"ShopNotice"`
	// 标志
	Flag int32 `protobuf:"varint,6,opt,name=Flag,proto3" json:"Flag"`
	// * 店铺标志
	Logo string `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo"`
	// * 个性化域名
	Alias string `protobuf:"bytes,8,opt,name=Alias,proto3" json:"Alias"`
	// * 自定义 域名
	Host string `protobuf:"bytes,9,opt,name=Host,proto3" json:"Host"`
	// * 电话
	Telephone string `protobuf:"bytes,10,opt,name=Telephone,proto3" json:"Telephone"`
	// * 状态
	State                int32    `protobuf:"varint,11,opt,name=State,proto3" json:"State"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShop) Reset()         { *m = SShop{} }
func (m *SShop) String() string { return proto.CompactTextString(m) }
func (*SShop) ProtoMessage()    {}
func (*SShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{1}
}
func (m *SShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShop.Unmarshal(m, b)
}
func (m *SShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShop.Marshal(b, m, deterministic)
}
func (dst *SShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShop.Merge(dst, src)
}
func (m *SShop) XXX_Size() int {
	return xxx_messageInfo_SShop.Size(m)
}
func (m *SShop) XXX_DiscardUnknown() {
	xxx_messageInfo_SShop.DiscardUnknown(m)
}

var xxx_messageInfo_SShop proto.InternalMessageInfo

func (m *SShop) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShop) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SShop) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SShop) GetShopTitle() string {
	if m != nil {
		return m.ShopTitle
	}
	return ""
}

func (m *SShop) GetShopNotice() string {
	if m != nil {
		return m.ShopNotice
	}
	return ""
}

func (m *SShop) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SShop) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShop) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShop) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShop) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *SShop) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

// 店铺设置
type SShopConfig struct {
	// * 店铺标志
	Logo string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo"`
	// * 自定义 域名
	Host string `protobuf:"bytes,5,opt,name=Host,proto3" json:"Host"`
	// * 个性化域名
	Alias string `protobuf:"bytes,6,opt,name=Alias,proto3" json:"Alias"`
	// * 电话
	Tel                  string   `protobuf:"bytes,7,opt,name=Tel,proto3" json:"Tel"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShopConfig) Reset()         { *m = SShopConfig{} }
func (m *SShopConfig) String() string { return proto.CompactTextString(m) }
func (*SShopConfig) ProtoMessage()    {}
func (*SShopConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{2}
}
func (m *SShopConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShopConfig.Unmarshal(m, b)
}
func (m *SShopConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShopConfig.Marshal(b, m, deterministic)
}
func (dst *SShopConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShopConfig.Merge(dst, src)
}
func (m *SShopConfig) XXX_Size() int {
	return xxx_messageInfo_SShopConfig.Size(m)
}
func (m *SShopConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SShopConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SShopConfig proto.InternalMessageInfo

func (m *SShopConfig) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShopConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShopConfig) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShopConfig) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

// 店铺编号
type ShopId struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopId) Reset()         { *m = ShopId{} }
func (m *ShopId) String() string { return proto.CompactTextString(m) }
func (*ShopId) ProtoMessage()    {}
func (*ShopId) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{3}
}
func (m *ShopId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopId.Unmarshal(m, b)
}
func (m *ShopId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopId.Marshal(b, m, deterministic)
}
func (dst *ShopId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopId.Merge(dst, src)
}
func (m *ShopId) XXX_Size() int {
	return xxx_messageInfo_ShopId.Size(m)
}
func (m *ShopId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopId.DiscardUnknown(m)
}

var xxx_messageInfo_ShopId proto.InternalMessageInfo

func (m *ShopId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 门店编号
type StoreId struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreId) Reset()         { *m = StoreId{} }
func (m *StoreId) String() string { return proto.CompactTextString(m) }
func (*StoreId) ProtoMessage()    {}
func (*StoreId) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{4}
}
func (m *StoreId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreId.Unmarshal(m, b)
}
func (m *StoreId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreId.Marshal(b, m, deterministic)
}
func (dst *StoreId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreId.Merge(dst, src)
}
func (m *StoreId) XXX_Size() int {
	return xxx_messageInfo_StoreId.Size(m)
}
func (m *StoreId) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreId.DiscardUnknown(m)
}

var xxx_messageInfo_StoreId proto.InternalMessageInfo

func (m *StoreId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 店铺
type SStore struct {
	Id           int64  `protobuf:"zigzag64,1,opt,name=id,proto3" json:"id"`
	MerchantId   int64  `protobuf:"zigzag64,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Alias        string `protobuf:"bytes,4,opt,name=Alias,proto3" json:"Alias"`
	State        int32  `protobuf:"zigzag32,5,opt,name=State,proto3" json:"State"`
	OpeningState int32  `protobuf:"zigzag32,8,opt,name=OpeningState,proto3" json:"OpeningState"`
	StorePhone   string `protobuf:"bytes,9,opt,name=StorePhone,proto3" json:"StorePhone"`
	StoreNotice  string `protobuf:"bytes,11,opt,name=StoreNotice,proto3" json:"StoreNotice"`
	Province     int32  `protobuf:"varint,12,opt,name=province,proto3" json:"province"`
	City         int32  `protobuf:"varint,13,opt,name=City,proto3" json:"City"`
	District     int32  `protobuf:"varint,14,opt,name=District,proto3" json:"District"`
	// 地区名称
	Address string `protobuf:"bytes,15,opt,name=Address,proto3" json:"Address"`
	// 详细地址
	DetailAddress string `protobuf:"bytes,16,opt,name=DetailAddress,proto3" json:"DetailAddress"`
	// 纬度
	Lat float64 `protobuf:"fixed64,17,opt,name=Lat,proto3" json:"Lat"`
	// 经度
	Lng float64 `protobuf:"fixed64,18,opt,name=Lng,proto3" json:"Lng"`
	// 覆盖范围(公里)
	CoverRadius int32 `protobuf:"varint,19,opt,name=CoverRadius,proto3" json:"CoverRadius"`
	// 序号
	SortNum              int32    `protobuf:"varint,20,opt,name=SortNum,proto3" json:"SortNum"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SStore) Reset()         { *m = SStore{} }
func (m *SStore) String() string { return proto.CompactTextString(m) }
func (*SStore) ProtoMessage()    {}
func (*SStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{5}
}
func (m *SStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SStore.Unmarshal(m, b)
}
func (m *SStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SStore.Marshal(b, m, deterministic)
}
func (dst *SStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SStore.Merge(dst, src)
}
func (m *SStore) XXX_Size() int {
	return xxx_messageInfo_SStore.Size(m)
}
func (m *SStore) XXX_DiscardUnknown() {
	xxx_messageInfo_SStore.DiscardUnknown(m)
}

var xxx_messageInfo_SStore proto.InternalMessageInfo

func (m *SStore) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SStore) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SStore) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SStore) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SStore) GetOpeningState() int32 {
	if m != nil {
		return m.OpeningState
	}
	return 0
}

func (m *SStore) GetStorePhone() string {
	if m != nil {
		return m.StorePhone
	}
	return ""
}

func (m *SStore) GetStoreNotice() string {
	if m != nil {
		return m.StoreNotice
	}
	return ""
}

func (m *SStore) GetProvince() int32 {
	if m != nil {
		return m.Province
	}
	return 0
}

func (m *SStore) GetCity() int32 {
	if m != nil {
		return m.City
	}
	return 0
}

func (m *SStore) GetDistrict() int32 {
	if m != nil {
		return m.District
	}
	return 0
}

func (m *SStore) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SStore) GetDetailAddress() string {
	if m != nil {
		return m.DetailAddress
	}
	return ""
}

func (m *SStore) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *SStore) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *SStore) GetCoverRadius() int32 {
	if m != nil {
		return m.CoverRadius
	}
	return 0
}

func (m *SStore) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

// 检查店铺结果
type CheckShopResponse struct {
	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`
	// 店铺开通状态,0:未开通 1:已开通 2:待审核 3:审核未通过
	Status               int32    `protobuf:"varint,2,opt,name=Status,proto3" json:"Status"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckShopResponse) Reset()         { *m = CheckShopResponse{} }
func (m *CheckShopResponse) String() string { return proto.CompactTextString(m) }
func (*CheckShopResponse) ProtoMessage()    {}
func (*CheckShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_303a641d623ae92b, []int{6}
}
func (m *CheckShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckShopResponse.Unmarshal(m, b)
}
func (m *CheckShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckShopResponse.Marshal(b, m, deterministic)
}
func (dst *CheckShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckShopResponse.Merge(dst, src)
}
func (m *CheckShopResponse) XXX_Size() int {
	return xxx_messageInfo_CheckShopResponse.Size(m)
}
func (m *CheckShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckShopResponse proto.InternalMessageInfo

func (m *CheckShopResponse) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *CheckShopResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CheckShopResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*TurnShopRequest)(nil), "TurnShopRequest")
	proto.RegisterType((*SShop)(nil), "SShop")
	proto.RegisterType((*SShopConfig)(nil), "SShopConfig")
	proto.RegisterType((*ShopId)(nil), "ShopId")
	proto.RegisterType((*StoreId)(nil), "StoreId")
	proto.RegisterType((*SStore)(nil), "SStore")
	proto.RegisterType((*CheckShopResponse)(nil), "CheckShopResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	// * 获取店铺,shopId
	GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error)
	// 删除商店
	DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *ShopId, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error) {
	out := new(CheckShopResponse)
	err := c.cc.Invoke(ctx, "/ShopService/CheckMerchantShopState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error) {
	out := new(SStore)
	err := c.cc.Invoke(ctx, "/ShopService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/TurnShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveOfflineShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	// * 获取店铺,shopId
	GetShop(context.Context, *ShopId) (*SShop, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(context.Context, *MerchantId) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(context.Context, *StoreId) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(context.Context, *String) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(context.Context, *TurnShopRequest) (*Result, error)
	// 保存门店
	SaveShop(context.Context, *SShop) (*Result, error)
	// 保存门店
	SaveOfflineShop(context.Context, *SStore) (*Result, error)
	// 删除商店
	DeleteStore(context.Context, *StoreId) (*Result, error)
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*ShopId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CheckMerchantShopState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/CheckMerchantShopState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_TurnShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).TurnShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/TurnShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).TurnShop(ctx, req.(*TurnShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveShop(ctx, req.(*SShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveOfflineShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveOfflineShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, req.(*SStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "CheckMerchantShopState",
			Handler:    _ShopService_CheckMerchantShopState_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _ShopService_GetStore_Handler,
		},
		{
			MethodName: "QueryShopByHost",
			Handler:    _ShopService_QueryShopByHost_Handler,
		},
		{
			MethodName: "TurnShop",
			Handler:    _ShopService_TurnShop_Handler,
		},
		{
			MethodName: "SaveShop",
			Handler:    _ShopService_SaveShop_Handler,
		},
		{
			MethodName: "SaveOfflineShop",
			Handler:    _ShopService_SaveOfflineShop_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _ShopService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop_service.proto",
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_shop_service_303a641d623ae92b) }

var fileDescriptor_shop_service_303a641d623ae92b = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xce, 0xaf, 0xed, 0x8c, 0x81, 0xc0, 0x82, 0x38, 0xab, 0x88, 0x43, 0x22, 0x8b, 0xa3, 0x93,
	0x2b, 0x1f, 0x89, 0x53, 0xf5, 0x86, 0x2b, 0x08, 0x6a, 0x8b, 0x54, 0xa0, 0xb5, 0x73, 0x55, 0xb5,
	0xaa, 0x8c, 0xb3, 0x38, 0x2b, 0x9c, 0x5d, 0xd7, 0x5e, 0x47, 0xe2, 0x4d, 0xfa, 0x3c, 0xbd, 0xed,
	0x4b, 0x55, 0x3b, 0xfe, 0x89, 0x01, 0x71, 0xe5, 0xf9, 0xbe, 0x99, 0x1d, 0x7f, 0x3b, 0x3f, 0x0b,
	0x24, 0x5b, 0xca, 0xe4, 0x7b, 0xc6, 0xd2, 0x35, 0x0f, 0x99, 0x9b, 0xa4, 0x52, 0xc9, 0xd1, 0x56,
	0x14, 0xcb, 0xbb, 0x20, 0x2e, 0x90, 0xe3, 0xc1, 0x70, 0x9e, 0xa7, 0xc2, 0x5f, 0xca, 0xc4, 0x63,
	0x3f, 0x72, 0x96, 0x29, 0xf2, 0x17, 0x98, 0x78, 0x8c, 0x2f, 0x68, 0x7b, 0xd2, 0x9e, 0x12, 0xcf,
	0xd0, 0xf0, 0x6a, 0x41, 0x76, 0xa0, 0x23, 0x05, 0xed, 0x4c, 0xda, 0x53, 0xcb, 0xeb, 0x48, 0x41,
	0x0e, 0xc1, 0x48, 0x59, 0x90, 0x49, 0x41, 0xbb, 0x93, 0xf6, 0x74, 0xe0, 0x95, 0xc8, 0xf9, 0xd9,
	0x81, 0xbe, 0xaf, 0x33, 0xea, 0x13, 0x65, 0x96, 0xae, 0xd7, 0xe1, 0x0b, 0x32, 0x06, 0x7b, 0xc5,
	0xd2, 0x70, 0x19, 0x08, 0xa5, 0xd3, 0x77, 0xd0, 0x01, 0x15, 0x75, 0xb5, 0x20, 0x23, 0xb0, 0xf4,
	0xc1, 0x9b, 0x60, 0xc5, 0xca, 0xa4, 0x35, 0x26, 0x47, 0x30, 0xd0, 0xf6, 0x9c, 0xab, 0x98, 0xd1,
	0x1e, 0x3a, 0x37, 0x04, 0x39, 0x06, 0xc0, 0x48, 0xa9, 0x78, 0xc8, 0x68, 0x1f, 0xdd, 0x0d, 0x86,
	0x10, 0xe8, 0xbd, 0x8b, 0x83, 0x88, 0x1a, 0x93, 0xf6, 0xb4, 0xef, 0xa1, 0xad, 0xb9, 0x58, 0x46,
	0x92, 0x9a, 0x18, 0x8d, 0x36, 0x39, 0x80, 0xfe, 0x79, 0xcc, 0x83, 0x8c, 0x5a, 0x48, 0x16, 0x40,
	0x47, 0x7e, 0x90, 0x99, 0xa2, 0x83, 0x22, 0x52, 0xdb, 0x5a, 0xcf, 0x9c, 0xc5, 0x2c, 0x59, 0x4a,
	0xc1, 0x28, 0x14, 0x7a, 0x6a, 0x42, 0xe7, 0xf1, 0x55, 0xa0, 0x18, 0xb5, 0xf1, 0x87, 0x05, 0x70,
	0xbe, 0x81, 0x8d, 0x95, 0x99, 0x49, 0x71, 0xcf, 0x37, 0x02, 0x7a, 0x0d, 0x01, 0xd5, 0xaf, 0xfa,
	0x8d, 0x5f, 0xd5, 0xa2, 0x8c, 0xa6, 0xa8, 0x5d, 0xe8, 0xce, 0x59, 0x5c, 0xaa, 0xd7, 0xa6, 0x73,
	0x0c, 0x86, 0x5f, 0xf4, 0xea, 0x00, 0xfa, 0xeb, 0x20, 0xce, 0x59, 0x59, 0xfc, 0x02, 0x38, 0x63,
	0x30, 0x7d, 0x25, 0x53, 0xf6, 0x6a, 0xc0, 0xef, 0x2e, 0x18, 0x3e, 0x86, 0x34, 0x7a, 0x47, 0x5e,
	0xeb, 0x1d, 0x79, 0xd2, 0x3b, 0x02, 0x3d, 0xb1, 0xe9, 0x1b, 0xda, 0x1b, 0xe1, 0xbd, 0xa6, 0xf0,
	0xba, 0x36, 0xfa, 0x8e, 0x7b, 0x65, 0x6d, 0x88, 0x03, 0x5b, 0xb7, 0x09, 0x13, 0x5c, 0x44, 0x85,
	0xd3, 0x42, 0xe7, 0x13, 0x0e, 0xbb, 0xac, 0xd5, 0x7d, 0xc2, 0xa2, 0x0f, 0xca, 0x2e, 0xd7, 0x0c,
	0x99, 0x80, 0x8d, 0xa8, 0x1c, 0x03, 0x1b, 0x03, 0x9a, 0x94, 0x9e, 0xb0, 0x24, 0x95, 0x6b, 0x2e,
	0x42, 0x46, 0xb7, 0xb0, 0x35, 0x35, 0xd6, 0x37, 0x98, 0x71, 0xf5, 0x48, 0xb7, 0x8b, 0x19, 0xd1,
	0xb6, 0x8e, 0xbf, 0xe4, 0x99, 0x4a, 0x79, 0xa8, 0xe8, 0x4e, 0x11, 0x5f, 0x61, 0x42, 0xc1, 0x3c,
	0x5f, 0x2c, 0x52, 0x96, 0x65, 0x74, 0x88, 0x7f, 0xaa, 0x20, 0x39, 0x81, 0xed, 0x4b, 0xa6, 0x02,
	0x1e, 0x57, 0xfe, 0x5d, 0xf4, 0x3f, 0x25, 0x75, 0x03, 0x3f, 0x06, 0x8a, 0xee, 0x4d, 0xda, 0xd3,
	0xb6, 0xa7, 0x4d, 0x64, 0x44, 0x44, 0x49, 0xc9, 0x88, 0x48, 0xdf, 0x68, 0x26, 0xd7, 0x2c, 0xf5,
	0x82, 0x05, 0xcf, 0x33, 0xba, 0x8f, 0x12, 0x9a, 0x94, 0x56, 0xe1, 0xcb, 0x54, 0xdd, 0xe4, 0x2b,
	0x7a, 0x80, 0xde, 0x0a, 0x3a, 0x5f, 0x61, 0x6f, 0xb6, 0x64, 0xe1, 0x43, 0xb1, 0xdd, 0x59, 0x22,
	0x45, 0xc6, 0x9e, 0xaf, 0x77, 0xb7, 0x5e, 0xef, 0x43, 0x30, 0x74, 0x91, 0xf3, 0x0c, 0x7b, 0xdb,
	0xf7, 0x4a, 0x54, 0xac, 0xf9, 0x2a, 0x48, 0x1f, 0x36, 0x6b, 0xae, 0xd1, 0xe9, 0xaf, 0x0e, 0xd8,
	0x3a, 0xb3, 0x5f, 0x3c, 0x2f, 0xe4, 0x08, 0xcc, 0xf7, 0x4c, 0xe1, 0xde, 0x9b, 0x6e, 0x31, 0x86,
	0x23, 0xc3, 0xc5, 0x71, 0x77, 0x5a, 0xe4, 0x0c, 0x0e, 0x51, 0xcb, 0x75, 0x39, 0x30, 0x78, 0x12,
	0x7b, 0x6a, 0xbb, 0xd7, 0xf5, 0x10, 0x8d, 0x88, 0xfb, 0x42, 0xb1, 0xd3, 0x22, 0x63, 0xb0, 0x74,
	0x6a, 0x9c, 0x4b, 0xcb, 0x2d, 0x47, 0x78, 0x64, 0xba, 0xc5, 0xa8, 0x3a, 0x2d, 0x72, 0x02, 0xc3,
	0xcf, 0x39, 0x4b, 0x1f, 0xf5, 0xb9, 0x8b, 0x47, 0xdc, 0x19, 0xd3, 0xf5, 0x55, 0xca, 0x45, 0x34,
	0x32, 0xdc, 0x2b, 0xa1, 0xde, 0xbe, 0x71, 0x5a, 0xe4, 0x5f, 0xb0, 0xaa, 0xc7, 0x8e, 0xec, 0xba,
	0xcf, 0xde, 0xbd, 0x91, 0xe9, 0x7a, 0x2c, 0xcb, 0x63, 0xe5, 0xb4, 0xc8, 0xdf, 0x60, 0xf9, 0xc1,
	0x9a, 0x61, 0x60, 0x79, 0x85, 0xa6, 0xfb, 0x1f, 0x18, 0x6a, 0xf7, 0xed, 0xfd, 0x7d, 0xcc, 0x05,
	0xab, 0x6e, 0x8c, 0x5a, 0x9a, 0x61, 0x0e, 0xd8, 0x97, 0x2c, 0x66, 0x8a, 0xbd, 0x14, 0x5e, 0xc5,
	0x5c, 0x8c, 0x61, 0x3f, 0x94, 0x2b, 0x37, 0xe2, 0x6a, 0x99, 0xdf, 0xb9, 0x91, 0x3c, 0x95, 0x6e,
	0x9a, 0x84, 0x5f, 0x2c, 0xf7, 0xbf, 0x33, 0x7c, 0xa0, 0xef, 0x0c, 0xfc, 0xfc, 0xff, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x63, 0x60, 0x57, 0xcb, 0x05, 0x00, 0x00,
}
