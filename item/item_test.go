package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestSpOrder(t *testing.T) {
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New(&xmhsdk.AuthParam{
		AppId:     "1000151",
		AppSecret: "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p",
	})
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "VIPER3@qq.com",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:         "OrderId3",
			SubOrderId:      "SubOrderId6",
			TotalPayPrice:   "150.00",
			Currency:        "USD",
			OrderState:      xmhsdk.OERDER_STATE_UNPAID,
			InsuredPayPrice: "3.00",
			ItemList: []*xmhsdk.DItem{{
				ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
				SkuId:             "SKU001",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "USD",
				UnitPrice:         "100.00",
				UnitNum:           "2",
				TotalPrice:        "200.00",
				PreferentialPrice: "50.00",
				TotalPayPrice:     "150.00",
			}, {
				ItemId:            "666666",
				SkuId:             "SKU002",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "USD",
				UnitPrice:         "100.00",
				UnitNum:           "2",
				TotalPrice:        "200.00",
				PreferentialPrice: "50.00",
				TotalPayPrice:     "150.00",
			}},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}
