/**
 * Copyright 2015 @ 56x.net.
 * name : return
 * author : jarryliu
 * date : 2016-07-17 17:29
 * description :
 * history :
 */
package afterSales

import (
	"errors"
	"math"

	afterSales "github.com/ixre/go2o/core/domain/interface/aftersales"
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/domain/interface/payment"
	"github.com/ixre/go2o/core/domain/tmp"
	"github.com/ixre/gof/db/orm"
)

var _ afterSales.IAfterSalesOrder = new(returnOrderImpl)
var _ afterSales.IReturnOrder = new(returnOrderImpl)
var _ afterSales.IReturnAfterSalesOrder = new(returnOrderImpl)

type returnOrderImpl struct {
	*afterSalesOrderImpl
	refValue    *afterSales.ReturnOrder
	memberRepo  member.IMemberRepo
	paymentRepo payment.IPaymentRepo
}

func newReturnOrderImpl(v *afterSalesOrderImpl, memberRepo member.IMemberRepo,
	paymentRepo payment.IPaymentRepo) *returnOrderImpl {
	if v.value.Type != afterSales.TypeReturn {
		panic(errors.New("售后单类型不是退货单"))
	}
	return &returnOrderImpl{
		afterSalesOrderImpl: v,
		memberRepo:          memberRepo,
		paymentRepo:         paymentRepo,
	}
}

func (r *returnOrderImpl) getValue() *afterSales.ReturnOrder {
	if r.refValue == nil {
		if r.GetDomainId() <= 0 {
			panic(errors.New("退货单还未提交"))
		}
		v := &afterSales.ReturnOrder{}
		if tmp.Orm.Get(r.GetDomainId(), v) != nil {
			panic(errors.New("退货单不存在"))
		}
		r.refValue = v
	}
	return r.refValue
}

// 获取售后单数据
func (r *returnOrderImpl) Value() afterSales.AfterSalesOrder {
	v := r.afterSalesOrderImpl.Value()
	v2 := r.getValue()
	v.Data = *v2
	return v
}

// 保存
func (r *returnOrderImpl) saveReturnOrder() error {
	_, err := orm.Save(tmp.Orm, r.refValue, int(r.GetDomainId()))
	return err
}

// 设置要退回货物信息
func (r *returnOrderImpl) SetItem(snapshotId int64, quantity int32) error {
	o := r.GetOrder()
	for _, v := range o.Items() {
		if v.SnapshotId == snapshotId {
			// 判断是否超过数量,减去已退货数量
			if v.Quantity-v.ReturnQuantity < quantity {
				return afterSales.ErrOutOfQuantity
			}
			// 设置退回商品
			r.value.SnapshotId = snapshotId
			r.value.Quantity = quantity
			return nil
		}
	}
	return afterSales.ErrNoSuchOrderItem
}

// 提交售后申请
func (r *returnOrderImpl) Submit() (int32, error) {
	//o := r.GetOrder()
	//if o.Value().State == order.StatCompleted {
	//    return 0, afterSales.ErrReturnAfterReceived
	//}
	id, err := r.afterSalesOrderImpl.Submit()
	// 提交退货单
	if err == nil {
		// 锁定退货数量
		err = r.GetOrder().Return(r.value.SnapshotId, r.value.Quantity)
		if err == nil {
			// 生成退货单
			err = r.submitReturnOrder()
		}
	}
	return id, err
}

// 提交退货单
func (r *returnOrderImpl) submitReturnOrder() (err error) {
	r.refValue = &afterSales.ReturnOrder{
		Id:       r.afterSalesOrderImpl.GetDomainId(),
		IsRefund: 0,
	}
	o := r.GetOrder()
	for _, v := range o.Items() {
		if v.SnapshotId == r.value.SnapshotId {
			price := v.FinalAmount / int64(v.Quantity) // 计算单价
			r.refValue.Amount = price * int64(r.value.Quantity)
			break
		}
	}
	if r.refValue.Amount <= 0 || math.IsNaN(float64(r.refValue.Amount)) {
		return afterSales.ErrOrderAmount
	}
	_, err = orm.Save(tmp.Orm, r.refValue, 0)
	return err
}

// 取消申请
func (r *returnOrderImpl) Cancel() error {
	err := r.afterSalesOrderImpl.Cancel()
	if err == nil {
		// 撤销退货数量
		err = r.GetOrder().RevertReturn(r.value.SnapshotId, r.value.Quantity)
	}
	return err
}

// 退回申请
func (r *returnOrderImpl) Reject(remark string) error {
	err := r.afterSalesOrderImpl.Reject(remark)
	if err == nil {
		// 撤销退货数量
		err = r.GetOrder().RevertReturn(r.value.SnapshotId, r.value.Quantity)
	}
	return err
}

// 完成退货
func (r *returnOrderImpl) Process() error {
	err := r.afterSalesOrderImpl.Process()
	if err == nil {
		err = r.handleReturn()
	}
	return err
}

// 处理退货
func (r *returnOrderImpl) handleReturn() error {

	//todo:??添加库存,或计入残次品

	v := r.getValue()
	if v.IsRefund == 1 {
		return nil
	}
	v.IsRefund = 1
	err := r.saveReturnOrder()
	if err == nil {
		err = r.backAmount(int(v.Amount))
	}
	return err
}

// 退款
func (r *returnOrderImpl) backAmount(amount int) error {
	o := r.GetOrder().GetValue()
	mm := r.memberRepo.GetMember(r.value.BuyerId)
	if mm == nil {
		return member.ErrNoSuchMember
	}
	//支付单与父订单关联。多个子订单合并付款
	po := r.paymentRepo.GetPaymentBySalesOrderId(o.OrderId)
	return po.Cancel()

	/**
	 *  重构支付单后调整，下面代码可能无用

	acc := mm.GetAccount()
	//如果支付单已清理数据，则全部退回到余额
	if po == nil {
		return acc.Refund(member.AccountBalance,
			member.KindRefund, "订单退款",
			o.OrderNo, amount,
			member.DefaultRelateUser)
	}
	//原路退回
	pv := po.Value()
	if pv.BalanceDiscount > 0 {
		if err := acc.Refund(member.AccountBalance, member.ChargeByRefund, "订单退款",
			o.OrderNo, pv.BalanceDiscount, member.DefaultRelateUser); err == nil {
			amount -= pv.BalanceDiscount
		}
	}
	//退积分
	if pv.IntegralDiscount > 0 {
	}
	//多退少补
	if pv.FinalFee > amount {
		amount = pv.FinalFee
	}
	//退到钱包账户
	if pv.PaymentSign == payment.SignWalletAccount {
		return acc.Refund(member.AccountWallet,
			member.KindWalletPaymentRefund,
			"订单退款", o.OrderNo, amount,
			member.DefaultRelateUser)
	}
	//原路退回，暂时不实现。直接退到钱包账户
	return acc.Refund(member.AccountWallet,
		member.KindWalletPaymentRefund,
		"订单退款", o.OrderNo, amount,
		member.DefaultRelateUser)
	*/
}



// 将换货的商品重新发货
func (r *returnOrderImpl) ReturnShipment(expressName string, expressOrder string, image string) error {
	if r.afterSalesOrderImpl.GetDomainId() <= 0 {
		panic(errors.New("换货单尚未提交"))
	}
	return r.afterSalesOrderImpl.ReturnShipment(expressName, expressOrder,image)	
}


// 接收换货
func (r *returnOrderImpl) ReturnReceive() error {
	return r.afterSalesOrderImpl.ReturnReceive()
}