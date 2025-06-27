package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestEuIDcSpOrder(t *testing.T) {
	xmhsdk.AppId = "1001"
	xmhsdk.AppSecret = "WkC2LaMgs7ckzxsBtDNSJR9YZb9B0wU3"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvEuIdc)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "wejnoospkonr@hldrive.com",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
		OrderInfo: &xmhsdk.DOrder{
			OrderId: strconv.Itoa(int(time.Now().Unix())),
			//SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "303.00",
			TotalPrice:        "303.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Add(-8 * time.Hour).Format(time.RFC3339),
			//OrderCreateTime:   time.Now().In(time.UTC).Add(-8 * time.Hour).Format(time.RFC3339),
			OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					//City:     "亚特兰提斯",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId:    "97760109-16ad-40c9-9385-caba381a26aa",
				SkuId:     "SKU001",
				ItemName:  "Durable Concrete Shirt",
				Currency:  "USD",
				UnitPrice: "0",
				UnitNum:   "2",
				//TotalPrice:        "0",
				//PreferentialPrice: "0",
				TotalPayPrice: "150.00",
			},
			//{
			//	ItemId:            "666666",
			//	SkuId:             "SKU002",
			//	ItemName:          "Durable Concrete Shirt",
			//	Currency:          "USD",
			//	UnitPrice:         "100.00",
			//	UnitNum:           "2",
			//	TotalPrice:        "200.00",
			//	PreferentialPrice: "50.00",
			//	TotalPayPrice:     "150.00",
			//}
			},
		},
		//ShipInfoList: []*xmhsdk.ShipInfo{
		//	{
		//		ShipId:             strconv.Itoa(int(time.Now().Unix())),
		//		ShipCompanyCode:    "SF",
		//		ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
		//		ShipStateString:    "已发货",
		//		ShipPrice:          "6.00",
		//		ActualShipSendTime: time.Now().Format(time.RFC3339),
		//		ShipOtherInfo: &xmhsdk.ShipAddress{
		//			Country:  "中国",
		//			Province: "浙江省",
		//			City:     "杭州市",
		//		},
		//		ItemList: []*xmhsdk.DItem{{
		//			ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		//			SkuId:             "SKU001",
		//			ItemName:          "Durable Concrete Shirt",
		//			Currency:          "USD",
		//			UnitPrice:         "100.00",
		//			UnitNum:           "2",
		//			TotalPrice:        "200.00",
		//			PreferentialPrice: "50.00",
		//			TotalPayPrice:     "150.00",
		//		}},
		//	},
		//},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}

func TestEuIDcPpOrder(t *testing.T) {
	xmhsdk.AppId = "1001"
	xmhsdk.AppSecret = "WkC2LaMgs7ckzxsBtDNSJR9YZb9B0wU3"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvEuIdc)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "wejnoospkonr@hldrive.com",
		UserId:    "VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:           strconv.Itoa(int(time.Now().Unix())),
			SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "303.00",
			TotalPrice:        "303.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "0",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "16GA5QRB5TO01_110",
				//SkuId:             "SKU001",
				//OrderGoodsId: strconv.Itoa(int(time.Now().UnixNano()) + 3),
				ItemName: "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency: "USD",
				//UnitPrice:         "100.00",
				UnitNum: "3",
				//TotalPrice:        "200.00",
				//PreferentialPrice: "50.00",
				TotalPayPrice:   "150.00",
				InsuredPayPrice: "36.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "16GA5QRB5TO01_110",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
					},
				},
			},
			//{
			//	ItemId: "16GA5QRB5TO01_110",
			//	//SkuId:             "SKU001",
			//	//OrderGoodsId: strconv.Itoa(int(time.Now().UnixNano()) + 5),
			//	ItemName: "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
			//	Currency: "USD",
			//	//UnitPrice:         "100.00",
			//	UnitNum: "2",
			//	//TotalPrice:        "200.00",
			//	//PreferentialPrice: "50.00",
			//	TotalPayPrice:   "150.00",
			//	InsuredPayPrice: "24.0",
			//	//Properties: map[string]string{
			//	//	"Plan ID": "10802003",
			//	//},
			//	PpVariant: &xmhsdk.PPVariant{
			//		VariantID:   "",
			//		Name:        xmhsdk.PPPlanCodeFor1Years,
			//		PriceString: "12.00",
			//		Currency:    "USD",
			//		Properties: map[string]string{
			//			"Reference": "16GA5QRB5TO01_110",
			//			"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
			//			"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
			//		},
			//	},
			//}
			},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}
