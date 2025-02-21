/**
 * Copyright 2015 @ 56x.net.
 * name : base_c
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package restapi

import (
	"context"
	"github.com/ixre/go2o/core/service"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// 获取传入的商户接口编号和密钥
func getUserInfo(c echo.Context) (string, string) {
	r := c.Request()
	apiId := r.FormValue("merchant_id")
	apiSecret := r.FormValue("secret")
	if len(apiId) == 0 {
		apiId = r.URL.Query().Get("merchant_id")
	}

	//todo: 兼容partner_id  ,将删除
	if len(apiId) == 0 {
		apiId = r.FormValue("partner_id")
		if len(apiId) == 0 {
			apiId = r.URL.Query().Get("partner_id")
		}
	}

	if len(apiSecret) == 0 {
		apiSecret = r.URL.Query().Get("secret")
	}
	return apiId, apiSecret
}

// 检查是否有权限调用接口(商户)
func chkMerchantApiSecret(c echo.Context) bool {
	i, s := getUserInfo(c)
	ok, mchId := CheckApiPermission(i, s)
	if ok {
		c.Set("merchant_id", mchId)
	}
	return ok
}

// 检查会员令牌信息
func checkMemberToken(c echo.Context) bool {
	r := c.Request()
	memberId, _ := util.I64Err(strconv.Atoi(r.FormValue("member_id")))
	token := r.FormValue("member_token")
	trans, cli, err := service.MemberServiceClient()
	if err == nil {
		defer trans.Close()
		if b, _ := cli.CheckToken(context.TODO(),
			&proto.CheckTokenRequest{
				MemberId: memberId,
				Token:    token,
			}); b.Value {
			c.Set("member_id", memberId)
			return true
		}
	}
	return false
}

// 获取商户编号
func getMerchantId(c echo.Context) int64 {
	return c.Get("merchant_id").(int64)
}

// 获取会员编号
func GetMemberId(c echo.Context) int64 {
	return c.Get("member_id").(int64)
}

func ApiTest(c echo.Context) error {
	return c.String(http.StatusOK, "It's working!")
}

// 检查是否有权限
func CheckApiPermission(apiId string, secret string) (bool, int64) {
	if len(apiId) != 0 && len(secret) != 0 {
		mchId := GetMerchantIdByApiId(apiId)
		var apiInfo = GetMerchantApiInfo(mchId)
		if apiInfo != nil {
			return apiInfo.ApiSecret == secret, mchId
		}
	}
	return false, 0
}
