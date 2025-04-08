package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
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

func TestPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
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
			//OrderCreateTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
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
				InsuredPayPrice:   "20.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "10.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "Durable Concrete Shirt",
					},
				},
			}},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}

func TestPpOrders(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "b565f726-ff0f-434d-aa31-e3e3cee15c0c",
		SkuId:             "sku1",
		VariantId:         "var1",
		ItemName:          "Gorgeous Bronze Knife",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
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
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
			ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"), {
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

func TestMutiOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "wejnoospkonr@hldrive.com",
		UserId:    "VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:           strconv.Itoa(int(time.Now().Unix())),
			SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "303.00",
			TotalPrice:        "303.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_SHIPPED,
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             strconv.Itoa(int(time.Now().Unix())),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
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
				Properties: map[string]string{
					"Plan ID": "10802003",
				},
				InsuredPayPrice: "3.00",
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        "1 Years",
					PriceString: "10.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
						"Plan ID":   "10802003",
						"Product":   "Durable Concrete Shirt",
					},
				},
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
			}, p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00")},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}

func TestOrderWithShip(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "wejnoospkonr@hldrive.com",
		DisComputeId: "xcp-00222580000000000508829558",
		UserId:       "xcp-VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:           strconv.Itoa(int(time.Now().Unix())),
			SubOrderId:        strconv.Itoa(int(time.Now().Unix())),
			TotalPayPrice:     "603.00",
			TotalPrice:        "600.00",
			PreferentialPrice: "0.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_SHIPPED,
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             strconv.Itoa(int(time.Now().Unix())),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
			ItemList: []*xmhsdk.DItem{
				p1.AddOneYearPp("12.00"),
				p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
			},
		},
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipCompany:        "SF",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("12.00"),
				},
			},
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipCompany:        "SF",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
				},
			},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}

func TestZeroAmountSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "wejnoospkonr@hldrive.com",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:           strconv.Itoa(int(time.Now().Unix())),
			SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "303.00",
			TotalPrice:        "303.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "1.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
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

func TestFullOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	x := &xmhsdk.CancelOrderParam{
		OrderId:           "1741247025",
		SubOrderId:        "SubOrderId6",
		DisXmhShopOrderId: "9372759558",
		CancelId:          strconv.FormatInt(time.Now().Unix(), 10),
		CancelReason:      "全部退款",
		CancelReasonType:  "1",
		Currency:          "USD",
		TotalRefundPrice:  "0",
		CancelItems:       nil,
	}
	cancel, err := Cancel(x)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(cancel))

}
