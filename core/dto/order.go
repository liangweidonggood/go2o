/**
 * Copyright 2015 @ 56x.net.
 * name : order
 * author : jarryliu
 * date : 2016-07-08 16:46
 * description :
 * history :
 */
package dto

type (
	/*
	   o.order_no,po.order_no as parent_no,
	       o.vendor_id,o.shop_id,s.shop_name,
	       o.item_amount,o.discount_amount,o.express_fee,
	       o.package_fee,o.final_fee,o.status
	*/

	// MemberPagingOrderDto 会员订单分页对象
	MemberPagingOrderDto struct {
		// 订单编号
		OrderId int64
		// 买家
		BuyerId int64
		// 买家用户名
		BuyerUser string
		// 店铺编号
		ShopId int64
		// 店铺名称
		ShopName string
		// 订单号
		OrderNo string
		// 商品数量
		ItemCount int
		// 商品总金额
		ItemAmount int64
		// 抵扣金额
		DiscountAmount int64
		// 优惠金额
		DeductAmount int64
		// 快递费
		ExpressFee int64
		// 包装费
		PackageFee int64
		// 最终金额
		FinalAmount int64
		// 是否支付
		IsPaid int32
		// 状态
		Status int32
		// 状态文本
		StatusText string
		// 下单时间
		CreateTime int64
		// 订单商品
		Items []*OrderItem
	}

	PagedVendorOrder struct {
		Id        int64
		OrderNo   string
		ParentNo  string
		BuyerId   int
		BuyerName string
		// 订单详情,主要描述订单的内容
		Details string
		//VendorId    int
		//ShopId      int
		//ShopName    string
		ItemAmount     int64
		DiscountAmount int64
		ExpressFee     int64
		PackageFee     int64
		IsPaid         bool
		FinalAmount    int64
		Status         int
		StatusText     string
		CreateTime     int64
		Items          []*OrderItem
		Data           map[string]string
	}

	/*
	   SELECT si.id,si.order_id,si.snap_id,sn.sku_id,sn.goods_title,sn.img,
	           si.quantity,si.fee,si.final_fee
	*/

	// OrderItem 订单商品项
	OrderItem struct {
		// 编号
		Id int
		// 订单编号
		OrderId int64
		// 商品快照编号
		SnapshotId int
		// Sku编号
		SkuId int
		// 商品编号
		ItemId int32
		// 商品标题
		ItemTitle string
		// 商品图片
		Image string
		// 商品单价
		Price int64
		// 商品实际单价
		FinalPrice int64
		// 商品数量
		Quantity int
		// 退货数量
		ReturnQuantity int
		// 商品总金额
		Amount int64
		// 商品实际总金额
		FinalAmount int64
		// 是否已发货
		IsShipped int
	}
)
