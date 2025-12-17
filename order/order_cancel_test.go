package order

import (
	"strconv"
	"testing"
	"time"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
)

func TestTestCancelSpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{

			{
				CancelId:       strconv.FormatInt(time.Now().Unix(), 10),
				ServiceOrderId: "20251203XSP1167613823E65",
				CancelTime:     time.Now().Format(time.RFC3339),
			},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))
}
func TestPartialCancelSpServiceOrder(t *testing.T) {
	//xmhsdk.AppId = "1000168"
	//xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	//xmhsdk.SignSecret = "vevor-alpha"
	//xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	xmhsdk.AppId = "1000170"
	xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	xmhsdk.SignSecret = "vevor_beta"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{
				CancelId:       strconv.FormatInt(time.Now().Unix(), 10),
				ServiceOrderId: "20251209XSP1167613829A65",
				CancelTime:     time.Now().Format(time.RFC3339),
				CancelFee:      "5",
				Currency:       "USD",
				CancelItems: []*xmhsdk.DItem{
					{
						ItemId: "16GA5QRB5TO01_110",
						//SkuId:             "SKU001",
						//OrderGoodsId: "16GA5QRB5TO01_110",
						ItemName:  "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
						Currency:  "USD",
						UnitPrice: "48.00",
						UnitNum:   "2",
						//TotalPrice:        "200.00",
						//PreferentialPrice: "50.00",
						TotalPayPrice:   "4500.00",
						InsuredPayPrice: "36.0",
					},
				},
			},
		}}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))
}

func TestCancelPpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{
				CancelId:       strconv.FormatInt(time.Now().Unix(), 10),
				ServiceOrderId: "20251210XSP116761382BE65",
				//CancelId:       "1765371131",
				CancelTime: time.Now().Format(time.RFC3339),
				CancelFee:  "5",
				Currency:   "USD",
				CancelItems: []*xmhsdk.DItem{
					{
						ItemId: "16GA5QRB5TO01_110",
						//SkuId:             "SKU001",
						//OrderGoodsId: "16GA5QRB5TO01_110",
						ItemName:  "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
						Currency:  "USD",
						UnitPrice: "48.00",
						UnitNum:   "1",
						//TotalPrice:        "200.00",
						//PreferentialPrice: "50.00",
						TotalPayPrice:   "2000.00",
						InsuredPayPrice: "36.0",
					},
				},
			},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))
}

func TestPartialCancelPpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{
				CancelId:       strconv.Itoa(int(time.Now().Unix())),
				ServiceOrderId: "20250314XPP11675436B4D58",
				CancelFee:      "10",
				Currency:       "USD",
				CancelItems: []*xmhsdk.DItem{{
					ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:             "SKU001",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "100.00",
					UnitNum:           "1",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
					InsuredPayPrice:   "20.0",
				}},
			},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))

}

func TestPartialCancelMutilPpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
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
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{
				CancelId:       strconv.Itoa(int(time.Now().Unix())),
				ServiceOrderId: "20250314XPP11675436B3358",
				CancelFee:      "5",
				Currency:       "USD",
				CancelItems:    []*xmhsdk.DItem{p1},
			},
			{
				CancelId:       strconv.Itoa(int(time.Now().Unix())),
				ServiceOrderId: "20250314XPP11675436B3458",
				CancelFee:      "5",
				Currency:       "USD",
				CancelItems:    []*xmhsdk.DItem{p2},
			},
			{
				CancelId:       strconv.Itoa(int(time.Now().Unix())),
				ServiceOrderId: "20250314XPP11675436B3558",
				CancelFee:      "5",
				Currency:       "USD",
				CancelItems:    []*xmhsdk.DItem{p3},
			},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))

}

func TestPpPartialCancelMutilPpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000170"
	xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	xmhsdk.SignSecret = "vevor_beta"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId: "VEVOR_US_1.5KWTSSGJ0000001V1",
		//SkuId:             "SKU001",
		OrderGoodsId: "7316476058016931840",
		ItemName:     "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
		Currency:     "USD",
		//UnitPrice:         "100.00",
		UnitNum: "1",
		//TotalPrice:        "200.00",
		//PreferentialPrice: "50.00",
		TotalPayPrice:   "150.00",
		InsuredPayPrice: "24.0",
	}
	//p2 := &xmhsdk.DItem{
	//	ItemId:            "b565f726-ff0f-434d-aa31-e3e3cee15c0c",
	//	SkuId:             "sku1",
	//	VariantId:         "var1",
	//	ItemName:          "Gorgeous Bronze Knife",
	//	Currency:          "USD",
	//	UnitPrice:         "100.00",
	//	UnitNum:           "1",
	//	TotalPrice:        "200.00",
	//	PreferentialPrice: "50.00",
	//	TotalPayPrice:     "150.00",
	//}
	//p3 := &xmhsdk.DItem{
	//	ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
	//	SkuId:             "SKU001",
	//	ItemName:          "Durable Concrete Shirt",
	//	Currency:          "USD",
	//	UnitPrice:         "100.00",
	//	UnitNum:           "1",
	//	TotalPrice:        "200.00",
	//	PreferentialPrice: "50.00",
	//	TotalPayPrice:     "150.00",
	//}
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{
				CancelId:       strconv.Itoa(int(time.Now().Unix())),
				ServiceOrderId: "20251210XPP83015378B936",
				CancelFee:      "5",
				//Currency:       "USD",
				CancelTime: time.Now().Format(time.RFC3339),

				CancelItems: []*xmhsdk.DItem{p1},
			},
			//{
			//	CancelId:       strconv.Itoa(int(time.Now().Unix())),
			//	ServiceOrderId: "20250314XPP11675436B3458",
			//	CancelFee:      "5",
			//	Currency:       "USD",
			//	CancelItems:    []*xmhsdk.DItem{p2},
			//},
			//{
			//	CancelId:       strconv.Itoa(int(time.Now().Unix())),
			//	ServiceOrderId: "20250314XPP11675436B3558",
			//	CancelFee:      "5",
			//	Currency:       "USD",
			//	CancelItems:    []*xmhsdk.DItem{p3},
			//},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))

}
