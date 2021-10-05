// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global.proto

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

// * 设置依据
type ESettingBasis int32

const (
	// * 未设置
	ESettingBasis_None ESettingBasis = 0
	// * 使用全局
	ESettingBasis_Global ESettingBasis = 1
	// * 自定义
	ESettingBasis_Custom ESettingBasis = 2
)

var ESettingBasis_name = map[int32]string{
	0: "None",
	1: "Global",
	2: "Custom",
}
var ESettingBasis_value = map[string]int32{
	"None":   0,
	"Global": 1,
	"Custom": 2,
}

func (x ESettingBasis) String() string {
	return proto.EnumName(ESettingBasis_name, int32(x))
}
func (ESettingBasis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{0}
}

// * 价格计算方式
type EPriceBasis int32

const (
	EPriceBasis__Unused EPriceBasis = 0
	// * 原价
	EPriceBasis_Original EPriceBasis = 1
	// * 会员折扣价
	EPriceBasis_Discount EPriceBasis = 2
	// * 自定义价格
	EPriceBasis_CustomBasis EPriceBasis = 3
)

var EPriceBasis_name = map[int32]string{
	0: "_Unused",
	1: "Original",
	2: "Discount",
	3: "CustomBasis",
}
var EPriceBasis_value = map[string]int32{
	"_Unused":     0,
	"Original":    1,
	"Discount":    2,
	"CustomBasis": 3,
}

func (x EPriceBasis) String() string {
	return proto.EnumName(EPriceBasis_name, int32(x))
}
func (EPriceBasis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{1}
}

// * 金额/提成依据
type EAmountBasis int32

const (
	EAmountBasis__12 EAmountBasis = 0
	// * 未设置
	EAmountBasis_NotSet EAmountBasis = 1
	// * 按金额
	EAmountBasis_amount EAmountBasis = 2
	// * 按百分比
	EAmountBasis_Percent EAmountBasis = 3
)

var EAmountBasis_name = map[int32]string{
	0: "_12",
	1: "NotSet",
	2: "amount",
	3: "Percent",
}
var EAmountBasis_value = map[string]int32{
	"_12":     0,
	"NotSet":  1,
	"amount":  2,
	"Percent": 3,
}

func (x EAmountBasis) String() string {
	return proto.EnumName(EAmountBasis_name, int32(x))
}
func (EAmountBasis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{2}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{1}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Int64 struct {
	Value                int64    `protobuf:"zigzag64,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{2}
}
func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (dst *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(dst, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32 struct {
	Value                int32    `protobuf:"zigzag32,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{3}
}
func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (dst *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(dst, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{4}
}
func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (dst *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(dst, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type StringMap struct {
	Value                map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StringMap) Reset()         { *m = StringMap{} }
func (m *StringMap) String() string { return proto.CompactTextString(m) }
func (*StringMap) ProtoMessage()    {}
func (*StringMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{5}
}
func (m *StringMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMap.Unmarshal(m, b)
}
func (m *StringMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMap.Marshal(b, m, deterministic)
}
func (dst *StringMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMap.Merge(dst, src)
}
func (m *StringMap) XXX_Size() int {
	return xxx_messageInfo_StringMap.Size(m)
}
func (m *StringMap) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMap.DiscardUnknown(m)
}

var xxx_messageInfo_StringMap proto.InternalMessageInfo

func (m *StringMap) GetValue() map[string]string {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringArray struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringArray) Reset()         { *m = StringArray{} }
func (m *StringArray) String() string { return proto.CompactTextString(m) }
func (*StringArray) ProtoMessage()    {}
func (*StringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{6}
}
func (m *StringArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringArray.Unmarshal(m, b)
}
func (m *StringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringArray.Marshal(b, m, deterministic)
}
func (dst *StringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringArray.Merge(dst, src)
}
func (m *StringArray) XXX_Size() int {
	return xxx_messageInfo_StringArray.Size(m)
}
func (m *StringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_StringArray.DiscardUnknown(m)
}

var xxx_messageInfo_StringArray proto.InternalMessageInfo

func (m *StringArray) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// 传输结果对象
type Result struct {
	// 状态码,如为0表示成功
	ErrCode int32 `protobuf:"zigzag32,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	// 消息
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	// * 数据字典
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{7}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Result) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *Result) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// 键值对
type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{8}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IdOrName struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdOrName) Reset()         { *m = IdOrName{} }
func (m *IdOrName) String() string { return proto.CompactTextString(m) }
func (*IdOrName) ProtoMessage()    {}
func (*IdOrName) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{9}
}
func (m *IdOrName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdOrName.Unmarshal(m, b)
}
func (m *IdOrName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdOrName.Marshal(b, m, deterministic)
}
func (dst *IdOrName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdOrName.Merge(dst, src)
}
func (m *IdOrName) XXX_Size() int {
	return xxx_messageInfo_IdOrName.Size(m)
}
func (m *IdOrName) XXX_DiscardUnknown() {
	xxx_messageInfo_IdOrName.DiscardUnknown(m)
}

var xxx_messageInfo_IdOrName proto.InternalMessageInfo

func (m *IdOrName) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IdOrName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IdAndRemark struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdAndRemark) Reset()         { *m = IdAndRemark{} }
func (m *IdAndRemark) String() string { return proto.CompactTextString(m) }
func (*IdAndRemark) ProtoMessage()    {}
func (*IdAndRemark) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{10}
}
func (m *IdAndRemark) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdAndRemark.Unmarshal(m, b)
}
func (m *IdAndRemark) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdAndRemark.Marshal(b, m, deterministic)
}
func (dst *IdAndRemark) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdAndRemark.Merge(dst, src)
}
func (m *IdAndRemark) XXX_Size() int {
	return xxx_messageInfo_IdAndRemark.Size(m)
}
func (m *IdAndRemark) XXX_DiscardUnknown() {
	xxx_messageInfo_IdAndRemark.DiscardUnknown(m)
}

var xxx_messageInfo_IdAndRemark proto.InternalMessageInfo

func (m *IdAndRemark) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IdAndRemark) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

// * 分页参数
type SPagingParams struct {
	// 开始记录数
	Begin int64 `protobuf:"zigzag64,1,opt,name=begin,proto3" json:"begin"`
	// 结束记录数
	End int64 `protobuf:"zigzag64,2,opt,name=end,proto3" json:"end"`
	// 条件
	Where string `protobuf:"bytes,3,opt,name=Where,proto3" json:"Where"`
	// 排序字段
	SortBy string `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by"`
	// 参数
	Parameters           map[string]string `protobuf:"bytes,5,rep,name=parameters,proto3" json:"parameters" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SPagingParams) Reset()         { *m = SPagingParams{} }
func (m *SPagingParams) String() string { return proto.CompactTextString(m) }
func (*SPagingParams) ProtoMessage()    {}
func (*SPagingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{11}
}
func (m *SPagingParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPagingParams.Unmarshal(m, b)
}
func (m *SPagingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPagingParams.Marshal(b, m, deterministic)
}
func (dst *SPagingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPagingParams.Merge(dst, src)
}
func (m *SPagingParams) XXX_Size() int {
	return xxx_messageInfo_SPagingParams.Size(m)
}
func (m *SPagingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SPagingParams.DiscardUnknown(m)
}

var xxx_messageInfo_SPagingParams proto.InternalMessageInfo

func (m *SPagingParams) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *SPagingParams) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *SPagingParams) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

func (m *SPagingParams) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *SPagingParams) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// * 分页结果
type SPagingResult struct {
	// * 代码
	ErrCode int32 `protobuf:"zigzag32,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	// * 消息
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	// * 总数
	Count int32 `protobuf:"zigzag32,3,opt,name=count,proto3" json:"count"`
	// * 数据
	Data string `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
	// * 额外的数据
	Extras               map[string]string `protobuf:"bytes,5,rep,name=extras,proto3" json:"extras" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SPagingResult) Reset()         { *m = SPagingResult{} }
func (m *SPagingResult) String() string { return proto.CompactTextString(m) }
func (*SPagingResult) ProtoMessage()    {}
func (*SPagingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{12}
}
func (m *SPagingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPagingResult.Unmarshal(m, b)
}
func (m *SPagingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPagingResult.Marshal(b, m, deterministic)
}
func (dst *SPagingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPagingResult.Merge(dst, src)
}
func (m *SPagingResult) XXX_Size() int {
	return xxx_messageInfo_SPagingResult.Size(m)
}
func (m *SPagingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SPagingResult.DiscardUnknown(m)
}

var xxx_messageInfo_SPagingResult proto.InternalMessageInfo

func (m *SPagingResult) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SPagingResult) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SPagingResult) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SPagingResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *SPagingResult) GetExtras() map[string]string {
	if m != nil {
		return m.Extras
	}
	return nil
}

// 树形节点
type STreeNode struct {
	// 文本
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	// 值
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	// 图标,icon与JS树形控件冲突
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon"`
	// 是否展开
	Expand bool `protobuf:"varint,4,opt,name=expand,proto3" json:"expand"`
	// 延迟加载
	Lazy bool `protobuf:"varint,5,opt,name=lazy,proto3" json:"lazy"`
	// 其他数据
	Data map[string]string `protobuf:"bytes,6,rep,name=data,proto3" json:"data" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 子节点
	Children []*STreeNode `protobuf:"bytes,7,rep,name=children,proto3" json:"children"`
	// 子节点编号
	Id int64 `protobuf:"varint,8,opt,name=Id,proto3" json:"Id"`
	// 是否目录，通常Children有元素,则为true
	Folder               bool     `protobuf:"varint,10,opt,name=folder,proto3" json:"folder"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *STreeNode) Reset()         { *m = STreeNode{} }
func (m *STreeNode) String() string { return proto.CompactTextString(m) }
func (*STreeNode) ProtoMessage()    {}
func (*STreeNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{13}
}
func (m *STreeNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_STreeNode.Unmarshal(m, b)
}
func (m *STreeNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_STreeNode.Marshal(b, m, deterministic)
}
func (dst *STreeNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_STreeNode.Merge(dst, src)
}
func (m *STreeNode) XXX_Size() int {
	return xxx_messageInfo_STreeNode.Size(m)
}
func (m *STreeNode) XXX_DiscardUnknown() {
	xxx_messageInfo_STreeNode.DiscardUnknown(m)
}

var xxx_messageInfo_STreeNode proto.InternalMessageInfo

func (m *STreeNode) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *STreeNode) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *STreeNode) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *STreeNode) GetExpand() bool {
	if m != nil {
		return m.Expand
	}
	return false
}

func (m *STreeNode) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *STreeNode) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *STreeNode) GetChildren() []*STreeNode {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *STreeNode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *STreeNode) GetFolder() bool {
	if m != nil {
		return m.Folder
	}
	return false
}

type Id struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{14}
}
func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (dst *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(dst, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MemberId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberId) Reset()         { *m = MemberId{} }
func (m *MemberId) String() string { return proto.CompactTextString(m) }
func (*MemberId) ProtoMessage()    {}
func (*MemberId) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{15}
}
func (m *MemberId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberId.Unmarshal(m, b)
}
func (m *MemberId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberId.Marshal(b, m, deterministic)
}
func (dst *MemberId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberId.Merge(dst, src)
}
func (m *MemberId) XXX_Size() int {
	return xxx_messageInfo_MemberId.Size(m)
}
func (m *MemberId) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberId.DiscardUnknown(m)
}

var xxx_messageInfo_MemberId proto.InternalMessageInfo

func (m *MemberId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MerchantId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantId) Reset()         { *m = MerchantId{} }
func (m *MerchantId) String() string { return proto.CompactTextString(m) }
func (*MerchantId) ProtoMessage()    {}
func (*MerchantId) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{16}
}
func (m *MerchantId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantId.Unmarshal(m, b)
}
func (m *MerchantId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantId.Marshal(b, m, deterministic)
}
func (dst *MerchantId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantId.Merge(dst, src)
}
func (m *MerchantId) XXX_Size() int {
	return xxx_messageInfo_MerchantId.Size(m)
}
func (m *MerchantId) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantId.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantId proto.InternalMessageInfo

func (m *MerchantId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SellerId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellerId) Reset()         { *m = SellerId{} }
func (m *SellerId) String() string { return proto.CompactTextString(m) }
func (*SellerId) ProtoMessage()    {}
func (*SellerId) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{17}
}
func (m *SellerId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellerId.Unmarshal(m, b)
}
func (m *SellerId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellerId.Marshal(b, m, deterministic)
}
func (dst *SellerId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellerId.Merge(dst, src)
}
func (m *SellerId) XXX_Size() int {
	return xxx_messageInfo_SellerId.Size(m)
}
func (m *SellerId) XXX_DiscardUnknown() {
	xxx_messageInfo_SellerId.DiscardUnknown(m)
}

var xxx_messageInfo_SellerId proto.InternalMessageInfo

func (m *SellerId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BuyerId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyerId) Reset()         { *m = BuyerId{} }
func (m *BuyerId) String() string { return proto.CompactTextString(m) }
func (*BuyerId) ProtoMessage()    {}
func (*BuyerId) Descriptor() ([]byte, []int) {
	return fileDescriptor_global_8674c57b3e883acf, []int{18}
}
func (m *BuyerId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyerId.Unmarshal(m, b)
}
func (m *BuyerId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyerId.Marshal(b, m, deterministic)
}
func (dst *BuyerId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyerId.Merge(dst, src)
}
func (m *BuyerId) XXX_Size() int {
	return xxx_messageInfo_BuyerId.Size(m)
}
func (m *BuyerId) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyerId.DiscardUnknown(m)
}

var xxx_messageInfo_BuyerId proto.InternalMessageInfo

func (m *BuyerId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*String)(nil), "String")
	proto.RegisterType((*Int64)(nil), "Int64")
	proto.RegisterType((*Int32)(nil), "Int32")
	proto.RegisterType((*Bool)(nil), "Bool")
	proto.RegisterType((*StringMap)(nil), "StringMap")
	proto.RegisterMapType((map[string]string)(nil), "StringMap.ValueEntry")
	proto.RegisterType((*StringArray)(nil), "StringArray")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterMapType((map[string]string)(nil), "Result.DataEntry")
	proto.RegisterType((*Pair)(nil), "Pair")
	proto.RegisterType((*IdOrName)(nil), "IdOrName")
	proto.RegisterType((*IdAndRemark)(nil), "IdAndRemark")
	proto.RegisterType((*SPagingParams)(nil), "SPagingParams")
	proto.RegisterMapType((map[string]string)(nil), "SPagingParams.ParametersEntry")
	proto.RegisterType((*SPagingResult)(nil), "SPagingResult")
	proto.RegisterMapType((map[string]string)(nil), "SPagingResult.ExtrasEntry")
	proto.RegisterType((*STreeNode)(nil), "STreeNode")
	proto.RegisterMapType((map[string]string)(nil), "STreeNode.DataEntry")
	proto.RegisterType((*Id)(nil), "Id")
	proto.RegisterType((*MemberId)(nil), "MemberId")
	proto.RegisterType((*MerchantId)(nil), "MerchantId")
	proto.RegisterType((*SellerId)(nil), "SellerId")
	proto.RegisterType((*BuyerId)(nil), "BuyerId")
	proto.RegisterEnum("ESettingBasis", ESettingBasis_name, ESettingBasis_value)
	proto.RegisterEnum("EPriceBasis", EPriceBasis_name, EPriceBasis_value)
	proto.RegisterEnum("EAmountBasis", EAmountBasis_name, EAmountBasis_value)
}

func init() { proto.RegisterFile("global.proto", fileDescriptor_global_8674c57b3e883acf) }

var fileDescriptor_global_8674c57b3e883acf = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x6d, 0x8f, 0xdb, 0x44,
	0x10, 0x6e, 0xe2, 0xc4, 0xf1, 0x8d, 0xaf, 0xd4, 0xb7, 0x1c, 0x60, 0x22, 0x68, 0x2b, 0x23, 0xd0,
	0xe9, 0x90, 0x16, 0x35, 0xe5, 0xa5, 0xbc, 0x4a, 0x97, 0x36, 0xaa, 0xf2, 0x21, 0x69, 0xe4, 0x40,
	0x91, 0xf8, 0x12, 0x6d, 0xec, 0xc1, 0xb1, 0x6a, 0xef, 0x46, 0xeb, 0x0d, 0xaa, 0xf9, 0x2d, 0xfc,
	0x02, 0x7e, 0x13, 0x5f, 0xf9, 0x1f, 0x68, 0x77, 0x9d, 0xb7, 0x42, 0x4f, 0x1c, 0x7c, 0xf2, 0xcc,
	0x3c, 0xf3, 0xcc, 0xcc, 0xb3, 0x1e, 0x7b, 0xe1, 0x34, 0x2b, 0xc4, 0x92, 0x15, 0x74, 0x2d, 0x85,
	0x12, 0x51, 0x0f, 0xba, 0xa3, 0x72, 0xad, 0xea, 0xe8, 0x2e, 0xb8, 0x73, 0x25, 0x73, 0x9e, 0x91,
	0x73, 0xe8, 0xfe, 0xc2, 0x8a, 0x0d, 0x86, 0xad, 0xfb, 0xad, 0x8b, 0x93, 0xd8, 0x3a, 0xd1, 0xfb,
	0xd0, 0x1d, 0x73, 0xf5, 0xf9, 0xa7, 0xc7, 0x30, 0x39, 0x86, 0x1f, 0x0e, 0x8e, 0xe1, 0xb3, 0x2d,
	0xfc, 0x1e, 0x74, 0x86, 0x42, 0x14, 0xc7, 0xa8, 0xb7, 0x45, 0x25, 0x9c, 0xd8, 0xde, 0x13, 0xb6,
	0x26, 0x1f, 0xef, 0x53, 0x9c, 0x0b, 0x7f, 0xf0, 0x16, 0xdd, 0x41, 0xf4, 0xb9, 0x8e, 0x8f, 0xb8,
	0x92, 0x75, 0xc3, 0xec, 0x3f, 0x02, 0xd8, 0x07, 0x49, 0x00, 0xce, 0x0b, 0xac, 0x9b, 0xb9, 0xb5,
	0xb9, 0xef, 0xd7, 0x3e, 0xd0, 0xf2, 0x55, 0xfb, 0x51, 0x2b, 0xfa, 0x00, 0x7c, 0x5b, 0xf8, 0x4a,
	0x4a, 0x56, 0x1f, 0x0e, 0xe6, 0xec, 0x45, 0xff, 0xd6, 0x02, 0x37, 0xc6, 0x6a, 0x53, 0x28, 0xf2,
	0x2e, 0x78, 0x28, 0xe5, 0x22, 0x11, 0xe9, 0x56, 0x5a, 0x0f, 0xa5, 0x7c, 0x2c, 0x52, 0x24, 0xef,
	0x80, 0x36, 0x17, 0x65, 0x95, 0x35, 0x6d, 0x5c, 0x94, 0x72, 0x52, 0x65, 0xe4, 0x43, 0xe8, 0xa4,
	0x4c, 0xb1, 0xd0, 0x31, 0x4a, 0xce, 0xa8, 0x2d, 0x45, 0x9f, 0x30, 0xc5, 0xac, 0x0a, 0x03, 0xf7,
	0xbf, 0x80, 0x93, 0x5d, 0xe8, 0x46, 0x1a, 0x28, 0x74, 0x66, 0x2c, 0x97, 0xff, 0x96, 0x13, 0x51,
	0xf0, 0xc6, 0xe9, 0x33, 0x39, 0x65, 0x25, 0x92, 0x37, 0xa0, 0x3d, 0x4e, 0x0d, 0xc5, 0x89, 0xdb,
	0xe3, 0x94, 0x10, 0xe8, 0xe8, 0x78, 0x43, 0x30, 0x76, 0xf4, 0x19, 0xf8, 0xe3, 0xf4, 0x8a, 0xa7,
	0x31, 0x96, 0x4c, 0xbe, 0xf8, 0x1b, 0xe5, 0x6d, 0x70, 0xa5, 0x41, 0xb6, 0xb2, 0xad, 0x17, 0xfd,
	0xd9, 0x82, 0xdb, 0xf3, 0x19, 0xcb, 0x72, 0x9e, 0xcd, 0x98, 0x64, 0x65, 0xa5, 0xc7, 0x59, 0x62,
	0x96, 0xf3, 0xed, 0xce, 0x18, 0x47, 0x8f, 0x8d, 0x3c, 0x35, 0x64, 0x12, 0x6b, 0x53, 0xe7, 0xfd,
	0xb8, 0x42, 0x89, 0xa1, 0x63, 0xc7, 0x36, 0x8e, 0x3e, 0xdf, 0x4a, 0x48, 0xb5, 0x58, 0xd6, 0x61,
	0xc7, 0x36, 0xd2, 0xee, 0xb0, 0x26, 0xdf, 0x01, 0xac, 0x75, 0x03, 0x54, 0x28, 0xab, 0xb0, 0x6b,
	0x4e, 0xf9, 0x2e, 0x3d, 0x6a, 0x4d, 0x67, 0xbb, 0x04, 0x7b, 0xe4, 0x07, 0x8c, 0xfe, 0xb7, 0x70,
	0xe7, 0x15, 0xf8, 0x46, 0xc7, 0xff, 0xc7, 0x5e, 0xe7, 0xff, 0x58, 0x92, 0x73, 0xe8, 0x26, 0x62,
	0xc3, 0x95, 0xd1, 0x7c, 0x16, 0x5b, 0x47, 0xbf, 0x0e, 0xb3, 0x3a, 0x56, 0xb0, 0xb1, 0xc9, 0x00,
	0x5c, 0x7c, 0xa9, 0x24, 0xdb, 0x4a, 0xed, 0xd3, 0xa3, 0xee, 0x74, 0x64, 0x40, 0x2b, 0xb3, 0xc9,
	0xec, 0x7f, 0x09, 0xfe, 0x41, 0xf8, 0x46, 0xf2, 0x7e, 0x6f, 0xc3, 0xc9, 0xfc, 0x7b, 0x89, 0x38,
	0xd5, 0xf3, 0x9f, 0x43, 0x57, 0xe5, 0xaa, 0xd8, 0xfd, 0x15, 0x8c, 0xf3, 0xcf, 0x6c, 0x3d, 0x7c,
	0x9e, 0x08, 0xde, 0xbc, 0x45, 0x63, 0xeb, 0x65, 0xc1, 0x97, 0x6b, 0xc6, 0x53, 0x23, 0xc9, 0x8b,
	0x1b, 0x4f, 0xe7, 0x16, 0xec, 0xd7, 0x3a, 0xec, 0x9a, 0xa8, 0xb1, 0xc9, 0x45, 0x23, 0xde, 0x35,
	0x32, 0xcf, 0xe9, 0x6e, 0x8a, 0x57, 0x3f, 0x1d, 0xf2, 0x11, 0x78, 0xc9, 0x2a, 0x2f, 0x52, 0x89,
	0x3c, 0xec, 0x99, 0x6c, 0xd8, 0x67, 0xc7, 0x3b, 0xac, 0x59, 0x5d, 0xef, 0x70, 0x75, 0x7f, 0x16,
	0x45, 0x8a, 0x32, 0x04, 0x3b, 0x8d, 0xf5, 0xfe, 0xfb, 0xa7, 0xd8, 0xd7, 0x0d, 0x34, 0xfe, 0x7c,
	0xf7, 0x7b, 0x73, 0x62, 0xeb, 0x44, 0xf7, 0xc1, 0x9b, 0x60, 0xb9, 0x44, 0xf9, 0xda, 0x8c, 0x08,
	0x60, 0x82, 0x32, 0x59, 0x31, 0xae, 0xae, 0xab, 0x32, 0xc7, 0xa2, 0xb8, 0xa6, 0xca, 0x3d, 0xe8,
	0x0d, 0x37, 0xf5, 0xeb, 0x13, 0x2e, 0x1f, 0xc0, 0xed, 0xd1, 0x1c, 0x95, 0xca, 0x79, 0x36, 0x64,
	0x55, 0x5e, 0x11, 0x0f, 0x3a, 0x53, 0xc1, 0x31, 0xb8, 0x45, 0x00, 0xdc, 0xa7, 0xe6, 0x5e, 0x08,
	0x5a, 0xda, 0x7e, 0xbc, 0xa9, 0x94, 0x28, 0x83, 0xf6, 0xe5, 0x53, 0xf0, 0x47, 0x33, 0x99, 0x27,
	0x68, 0x09, 0x3e, 0xf4, 0x16, 0x3f, 0xf0, 0x4d, 0x85, 0x69, 0x70, 0x8b, 0x9c, 0x82, 0xf7, 0x4c,
	0xe6, 0x59, 0xce, 0x0d, 0xeb, 0x14, 0xbc, 0x27, 0x79, 0x65, 0xb6, 0x37, 0x68, 0x93, 0x3b, 0xe0,
	0xdb, 0x1a, 0x86, 0x17, 0x38, 0x97, 0xdf, 0xc0, 0xe9, 0xe8, 0xaa, 0xd4, 0xa8, 0xad, 0xd4, 0x03,
	0x67, 0xf1, 0x60, 0x60, 0x3b, 0x4f, 0x85, 0x9a, 0xa3, 0xb2, 0x9d, 0x59, 0xd9, 0x54, 0xf0, 0xa1,
	0x37, 0x43, 0x99, 0x20, 0x57, 0x81, 0x33, 0xbc, 0x07, 0x6f, 0x26, 0xa2, 0xa4, 0x59, 0xae, 0x56,
	0x9b, 0x25, 0xcd, 0xc4, 0x40, 0x50, 0xb9, 0x4e, 0x7e, 0xf2, 0xe8, 0x27, 0x5f, 0x9b, 0x7b, 0x6c,
	0xe9, 0x9a, 0xc7, 0xc3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x60, 0xd6, 0x50, 0x84, 0xde, 0x06,
	0x00, 0x00,
}
