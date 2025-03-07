package ship

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestShipSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()

	params := &xmhsdk.ShipParam{
		OrderId:    "1741232270",
		SubOrderId: "SubOrderId6",
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
				}},
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

func TestShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	item := xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}

	params := &xmhsdk.ShipParam{
		OrderId:    "1741234403",
		SubOrderId: "SubOrderId6",
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
				ItemList: []*xmhsdk.DItem{item.AddOneYearPp("")},
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
func TestShip(t *testing.T) {
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
	params := &xmhsdk.ShipParam{
		OrderId:    "1741315736",
		SubOrderId: "SubOrderId6",
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
					PpVariant: &xmhsdk.PPVariant{
						VariantID:   "",
						Name:        "3 Years",
						PriceString: "10.00",
						Currency:    "USD",
						Properties: map[string]string{
							"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
							"Plan ID":   "10802001",
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
				}, p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
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
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipWithShipTime(t *testing.T) {
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
	params := &xmhsdk.ShipParam{
		OrderId:    "1741315736",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:          strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode: "SF",
				ShipTrackNumber: strconv.Itoa(int(time.Now().Unix())),
				ShipStateString: "已发货",
				ShipPrice:       "6.00",
				//ActualShipSendTime: time.Now().Format(time.RFC3339),
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
					PpVariant: &xmhsdk.PPVariant{
						VariantID:   "",
						Name:        "3 Years",
						PriceString: "10.00",
						Currency:    "USD",
						Properties: map[string]string{
							"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
							"Plan ID":   "10802001",
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
				}, p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
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
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipMultiPpOrder(t *testing.T) {
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
		UnitNum:           "1",
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
		UnitNum:           "1",
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
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}

	params := &xmhsdk.ShipParam{
		OrderId:    "1741345418",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00")},
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
