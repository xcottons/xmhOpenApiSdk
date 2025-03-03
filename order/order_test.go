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
		UserEmail:    "VIPER3@qq.com",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
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
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
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
		},
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
		UserEmail: "VIPER3@qq.com",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:         "order666666",
			SubOrderId:      "SubOrderId6",
			TotalPayPrice:   "603.00",
			Currency:        "USD",
			OrderState:      xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice: "0.00",
			TaxPrice:        "0.00",
			ShipPrice:       "0.00",
			PayTime:         "2025-01-02T15:04:05Z07:00",
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
		UserEmail: "VIPER3@qq.com",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:         "order6666668",
			SubOrderId:      "SubOrderId6",
			TotalPayPrice:   "603.00",
			Currency:        "USD",
			OrderState:      xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice: "3.00",
			TaxPrice:        "0.00",
			ShipPrice:       "0.00",
			PayTime:         "2025-01-02T15:04:05Z07:00",
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
		UserEmail: "VIPER3@qq.com",
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
		UserEmail:    "VIPER3@qq.com",
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
			ItemList: []*xmhsdk.DItem{
				//	{
				//	ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
				//	SkuId:             "SKU001",
				//	ItemName:          "Durable Concrete Shirt",
				//	Currency:          "USD",
				//	UnitPrice:         "100.00",
				//	UnitNum:           "2",
				//	TotalPrice:        "200.00",
				//	PreferentialPrice: "50.00",
				//	TotalPayPrice:     "150.00",
				//	Properties: map[string]string{
				//		"Plan ID": "10802003",
				//	},
				//	PpVariant: &xmhsdk.PPVariant{
				//		VariantID:   "",
				//		Name:        "1 Years",
				//		PriceString: "10.00",
				//		Currency:    "USD",
				//		Properties: map[string]string{
				//			"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
				//			"Plan ID":   "10802003",
				//			"Product":   "Durable Concrete Shirt",
				//		},
				//	},
				//}, {
				//	ItemId:            "666666",
				//	SkuId:             "SKU002",
				//	ItemName:          "Durable Concrete Shirt",
				//	Currency:          "USD",
				//	UnitPrice:         "100.00",
				//	UnitNum:           "2",
				//	TotalPrice:        "200.00",
				//	PreferentialPrice: "50.00",
				//	TotalPayPrice:     "150.00",
				//},
				p1.AddThreeYearPp("12.00"),
				p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
			},
		},
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:             "SKU001",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "100.00",
					UnitNum:           "1",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
					PpVariant: &xmhsdk.PPVariant{
						VariantID:   "",
						Name:        "3 Years",
						PriceString: "10.00",
						Currency:    "USD",
						Properties: map[string]string{
							"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
							"Plan ID":   xmhsdk.PPPlanCodeFor2Years,
							"Product":   "Durable Concrete Shirt",
						},
					},
				},
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
