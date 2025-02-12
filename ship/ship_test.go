package ship

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestShip(t *testing.T) {
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New(&xmhsdk.AuthParam{
		AppId:     "1000151",
		AppSecret: "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p",
	})

	params := &xmhsdk.ShipParam{
		OrderId:    "order6666668",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipCompanyCode: "SF",
				ShipTrackNumber: "123456789",
				ShipStateString: "已发货",
				ShipPrice:       "6.00",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{
					{ItemId: "97760109-16ad-40c9-9385-caba381a26aa",
						SkuId:             "SKU001",
						ItemName:          "Durable Concrete Shirt",
						Currency:          "USD",
						UnitPrice:         "100.00",
						UnitNum:           "1",
						TotalPrice:        "200.00",
						PreferentialPrice: "50.00",
						TotalPayPrice:     "150.00"},
				},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}
