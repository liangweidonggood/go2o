// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_service.proto

package proto // import "."

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 获取产品模型
	GetModel(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*SProductModel, error)
	// 获取产品模型
	GetModels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductModelListResponse, error)
	// 获取属性
	GetAttr(ctx context.Context, in *ProductAttrId, opts ...grpc.CallOption) (*SProductAttr, error)
	// 获取属性项
	GetAttrItem_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductAttrItem, error)
	// 获取模型属性Html
	GetModelAttrsHtml(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*String, error)
	// 保存产品模型
	SaveModel(ctx context.Context, in *SProductModel, opts ...grpc.CallOption) (*Result, error)
	// 删除产品模型
	DeleteProModel_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// Get 产品品牌
	GetProBrand_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductBrand, error)
	// Save 产品品牌
	SaveProBrand_(ctx context.Context, in *SProductBrand, opts ...grpc.CallOption) (*Result, error)
	// Delete 产品品牌
	DeleteProBrand_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取所有产品品牌
	GetBrands(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductBrandListResponse, error)
	// 获取全部分类
	GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductCategoriesResponse, error)
	// 获取商品分类
	GetCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductCategory, error)
	// 获取商品分类和选项
	//    rpc GetCategoryAndOptions(mchId int64, id int32) (*product.Category,
	//    domain.IOptionStore)
	DeleteCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 保存产品分类
	SaveCategory(ctx context.Context, in *SProductCategory, opts ...grpc.CallOption) (*Result, error)
	// 根据上级编号获取分类列表
	GetChildren(ctx context.Context, in *CategoryParentId, opts ...grpc.CallOption) (*ProductCategoriesResponse, error)
	// 获取产品值
	GetProductValue(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*SProduct, error)
	// 保存产品
	SaveProduct(ctx context.Context, in *SProduct, opts ...grpc.CallOption) (*SaveProductResponse, error)
	// 删除产品
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存货品描述
	SaveProductInfo(ctx context.Context, in *ProductInfoRequest, opts ...grpc.CallOption) (*Result, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetModel(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*SProductModel, error) {
	out := new(SProductModel)
	err := c.cc.Invoke(ctx, "/ProductService/GetModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetModels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductModelListResponse, error) {
	out := new(ProductModelListResponse)
	err := c.cc.Invoke(ctx, "/ProductService/GetModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAttr(ctx context.Context, in *ProductAttrId, opts ...grpc.CallOption) (*SProductAttr, error) {
	out := new(SProductAttr)
	err := c.cc.Invoke(ctx, "/ProductService/GetAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAttrItem_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductAttrItem, error) {
	out := new(SProductAttrItem)
	err := c.cc.Invoke(ctx, "/ProductService/GetAttrItem_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetModelAttrsHtml(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/ProductService/GetModelAttrsHtml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveModel(ctx context.Context, in *SProductModel, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/SaveModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProModel_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProModel_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProBrand_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductBrand, error) {
	out := new(SProductBrand)
	err := c.cc.Invoke(ctx, "/ProductService/GetProBrand_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProBrand_(ctx context.Context, in *SProductBrand, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/SaveProBrand_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProBrand_(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProBrand_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetBrands(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductBrandListResponse, error) {
	out := new(ProductBrandListResponse)
	err := c.cc.Invoke(ctx, "/ProductService/GetBrands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductCategoriesResponse, error) {
	out := new(ProductCategoriesResponse)
	err := c.cc.Invoke(ctx, "/ProductService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductCategory, error) {
	out := new(SProductCategory)
	err := c.cc.Invoke(ctx, "/ProductService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveCategory(ctx context.Context, in *SProductCategory, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/SaveCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetChildren(ctx context.Context, in *CategoryParentId, opts ...grpc.CallOption) (*ProductCategoriesResponse, error) {
	out := new(ProductCategoriesResponse)
	err := c.cc.Invoke(ctx, "/ProductService/GetChildren", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductValue(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*SProduct, error) {
	out := new(SProduct)
	err := c.cc.Invoke(ctx, "/ProductService/GetProductValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProduct(ctx context.Context, in *SProduct, opts ...grpc.CallOption) (*SaveProductResponse, error) {
	out := new(SaveProductResponse)
	err := c.cc.Invoke(ctx, "/ProductService/SaveProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProductInfo(ctx context.Context, in *ProductInfoRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ProductService/SaveProductInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// 获取产品模型
	GetModel(context.Context, *ProductModelId) (*SProductModel, error)
	// 获取产品模型
	GetModels(context.Context, *Empty) (*ProductModelListResponse, error)
	// 获取属性
	GetAttr(context.Context, *ProductAttrId) (*SProductAttr, error)
	// 获取属性项
	GetAttrItem_(context.Context, *Int64) (*SProductAttrItem, error)
	// 获取模型属性Html
	GetModelAttrsHtml(context.Context, *Int64) (*String, error)
	// 保存产品模型
	SaveModel(context.Context, *SProductModel) (*Result, error)
	// 删除产品模型
	DeleteProModel_(context.Context, *Int64) (*Result, error)
	// Get 产品品牌
	GetProBrand_(context.Context, *Int64) (*SProductBrand, error)
	// Save 产品品牌
	SaveProBrand_(context.Context, *SProductBrand) (*Result, error)
	// Delete 产品品牌
	DeleteProBrand_(context.Context, *Int64) (*Result, error)
	// 获取所有产品品牌
	GetBrands(context.Context, *Empty) (*ProductBrandListResponse, error)
	// 获取全部分类
	GetCategories(context.Context, *Empty) (*ProductCategoriesResponse, error)
	// 获取商品分类
	GetCategory(context.Context, *Int64) (*SProductCategory, error)
	// 获取商品分类和选项
	//    rpc GetCategoryAndOptions(mchId int64, id int32) (*product.Category,
	//    domain.IOptionStore)
	DeleteCategory(context.Context, *Int64) (*Result, error)
	// 保存产品分类
	SaveCategory(context.Context, *SProductCategory) (*Result, error)
	// 根据上级编号获取分类列表
	GetChildren(context.Context, *CategoryParentId) (*ProductCategoriesResponse, error)
	// 获取产品值
	GetProductValue(context.Context, *ProductId) (*SProduct, error)
	// 保存产品
	SaveProduct(context.Context, *SProduct) (*SaveProductResponse, error)
	// 删除产品
	DeleteProduct(context.Context, *DeleteProductRequest) (*Result, error)
	// 保存货品描述
	SaveProductInfo(context.Context, *ProductInfoRequest) (*Result, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductModelId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetModel(ctx, req.(*ProductModelId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetModels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAttrId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAttr(ctx, req.(*ProductAttrId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAttrItem__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAttrItem_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetAttrItem_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAttrItem_(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetModelAttrsHtml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetModelAttrsHtml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetModelAttrsHtml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetModelAttrsHtml(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SProductModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SaveModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveModel(ctx, req.(*SProductModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProModel__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProModel_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProModel_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProModel_(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProBrand__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProBrand_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetProBrand_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProBrand_(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProBrand__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SProductBrand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProBrand_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SaveProBrand_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProBrand_(ctx, req.(*SProductBrand))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProBrand__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProBrand_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProBrand_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProBrand_(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetBrands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetBrands(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategory(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteCategory(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SProductCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SaveCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveCategory(ctx, req.(*SProductCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryParentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetChildren(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetChildren",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetChildren(ctx, req.(*CategoryParentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetProductValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductValue(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SaveProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProduct(ctx, req.(*SProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/SaveProductInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProductInfo(ctx, req.(*ProductInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModel",
			Handler:    _ProductService_GetModel_Handler,
		},
		{
			MethodName: "GetModels",
			Handler:    _ProductService_GetModels_Handler,
		},
		{
			MethodName: "GetAttr",
			Handler:    _ProductService_GetAttr_Handler,
		},
		{
			MethodName: "GetAttrItem_",
			Handler:    _ProductService_GetAttrItem__Handler,
		},
		{
			MethodName: "GetModelAttrsHtml",
			Handler:    _ProductService_GetModelAttrsHtml_Handler,
		},
		{
			MethodName: "SaveModel",
			Handler:    _ProductService_SaveModel_Handler,
		},
		{
			MethodName: "DeleteProModel_",
			Handler:    _ProductService_DeleteProModel__Handler,
		},
		{
			MethodName: "GetProBrand_",
			Handler:    _ProductService_GetProBrand__Handler,
		},
		{
			MethodName: "SaveProBrand_",
			Handler:    _ProductService_SaveProBrand__Handler,
		},
		{
			MethodName: "DeleteProBrand_",
			Handler:    _ProductService_DeleteProBrand__Handler,
		},
		{
			MethodName: "GetBrands",
			Handler:    _ProductService_GetBrands_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _ProductService_GetCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ProductService_GetCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _ProductService_DeleteCategory_Handler,
		},
		{
			MethodName: "SaveCategory",
			Handler:    _ProductService_SaveCategory_Handler,
		},
		{
			MethodName: "GetChildren",
			Handler:    _ProductService_GetChildren_Handler,
		},
		{
			MethodName: "GetProductValue",
			Handler:    _ProductService_GetProductValue_Handler,
		},
		{
			MethodName: "SaveProduct",
			Handler:    _ProductService_SaveProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "SaveProductInfo",
			Handler:    _ProductService_SaveProductInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service.proto",
}

func init() {
	proto.RegisterFile("product_service.proto", fileDescriptor_product_service_2703bb90cc56b302)
}

var fileDescriptor_product_service_2703bb90cc56b302 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4b, 0x8f, 0xd3, 0x30,
	0x10, 0xc7, 0x73, 0x6a, 0xa8, 0xb7, 0x0f, 0xd5, 0xd0, 0x03, 0x3e, 0x06, 0x84, 0x10, 0x0f, 0xc3,
	0xf2, 0xba, 0x20, 0x21, 0xb1, 0x80, 0x4a, 0x24, 0x90, 0xaa, 0x56, 0xe2, 0xc0, 0xa5, 0xca, 0xae,
	0x87, 0x12, 0xc9, 0x8d, 0x83, 0x3d, 0x59, 0x69, 0xbf, 0x11, 0x1f, 0x13, 0xd9, 0xce, 0xc3, 0x49,
	0x8b, 0xf6, 0x14, 0xcd, 0xfc, 0x7f, 0xff, 0xf1, 0x3c, 0x14, 0xb2, 0x2c, 0xb5, 0x12, 0xd5, 0x15,
	0xee, 0x0c, 0xe8, 0xeb, 0xfc, 0x0a, 0x78, 0xa9, 0x15, 0x2a, 0x36, 0xd9, 0x4b, 0x75, 0x99, 0xc9,
	0x3a, 0x5a, 0x1e, 0xc0, 0x98, 0x6c, 0x0f, 0x2f, 0x6a, 0xd8, 0xa7, 0x5f, 0xfd, 0x8d, 0xc9, 0x6c,
	0xed, 0x33, 0x5b, 0xef, 0xa6, 0xcf, 0xc9, 0x9d, 0x15, 0xe0, 0x77, 0x25, 0x40, 0xd2, 0x39, 0xaf,
	0x45, 0x17, 0xa6, 0x82, 0xcd, 0xf8, 0x36, 0xcc, 0x24, 0x11, 0x7d, 0x49, 0xc6, 0x0d, 0x6e, 0xe8,
	0x88, 0x7f, 0x39, 0x94, 0x78, 0xc3, 0xee, 0xf7, 0x7c, 0xdf, 0x72, 0x83, 0x1b, 0x30, 0xa5, 0x2a,
	0x0c, 0x24, 0x11, 0x7d, 0x42, 0xe2, 0x15, 0xe0, 0x47, 0x44, 0x4d, 0x67, 0x0d, 0x67, 0xa3, 0x54,
	0xb0, 0x69, 0x5b, 0xde, 0x26, 0x92, 0x88, 0x3e, 0x25, 0x93, 0x9a, 0x4d, 0x11, 0x0e, 0x3b, 0x3a,
	0xe2, 0x69, 0x81, 0xef, 0xde, 0xb0, 0x45, 0x0f, 0xb4, 0x5a, 0x12, 0xd1, 0x47, 0x64, 0xd1, 0xb4,
	0x62, 0xb3, 0xe6, 0x2b, 0x1e, 0x64, 0xeb, 0x88, 0xf9, 0x16, 0x75, 0x5e, 0xec, 0x1d, 0x37, 0xde,
	0x66, 0xd7, 0xe0, 0x47, 0x1c, 0x4c, 0xc4, 0x62, 0xbe, 0x01, 0x53, 0x49, 0x4c, 0x22, 0xfa, 0x90,
	0xcc, 0x3f, 0x83, 0x04, 0x84, 0xb5, 0x56, 0x4e, 0xdc, 0x05, 0xd5, 0x5a, 0xea, 0xb1, 0x6b, 0x71,
	0xad, 0xd5, 0x85, 0xce, 0x0a, 0xd1, 0x21, 0x5d, 0x61, 0x27, 0xb8, 0xc1, 0xa7, 0xf6, 0xdd, 0x0e,
	0x1d, 0x20, 0xff, 0x7b, 0x7b, 0x50, 0x38, 0xa0, 0xfc, 0xf2, 0x9d, 0x7e, 0x62, 0xf9, 0x2e, 0x3f,
	0x58, 0xfe, 0x5b, 0x32, 0x5d, 0x01, 0x7e, 0xca, 0x10, 0xf6, 0x4a, 0xe7, 0xd0, 0xb9, 0x58, 0xe3,
	0xea, 0xb4, 0xde, 0xcd, 0xce, 0x3a, 0xdb, 0xcd, 0x89, 0x33, 0x34, 0x52, 0x12, 0xd1, 0x07, 0x64,
	0xe6, 0x5b, 0x3f, 0xc2, 0x83, 0xce, 0x9f, 0x91, 0x89, 0xdd, 0x45, 0x8b, 0x1c, 0x57, 0x0a, 0xe9,
	0x0f, 0xfe, 0xf9, 0xdf, 0xb9, 0x14, 0x1a, 0x0a, 0xba, 0xe0, 0x0d, 0xb4, 0xce, 0x34, 0x14, 0x98,
	0x8a, 0x5b, 0xdb, 0x9f, 0xfb, 0x1b, 0x59, 0xe2, 0x47, 0x26, 0x2b, 0xa0, 0xa4, 0x31, 0xa4, 0x82,
	0x8d, 0xdb, 0xc7, 0xdd, 0x4e, 0xcf, 0xea, 0x2b, 0xd9, 0x04, 0xed, 0x34, 0x76, 0x8f, 0x07, 0x42,
	0x50, 0xfd, 0x9c, 0x4c, 0xdb, 0x5b, 0x39, 0xcf, 0x92, 0xf7, 0xe2, 0x0d, 0xfc, 0xa9, 0xc0, 0x60,
	0x38, 0xd0, 0x39, 0x99, 0x07, 0xb5, 0xd2, 0xe2, 0x97, 0xa2, 0x77, 0x79, 0x10, 0x1d, 0x5b, 0x2e,
	0xc6, 0x3f, 0x63, 0xfe, 0xde, 0xfd, 0xb5, 0x97, 0x23, 0xf7, 0x79, 0xfd, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x62, 0x7f, 0x4d, 0xe3, 0xfa, 0x03, 0x00, 0x00,
}
