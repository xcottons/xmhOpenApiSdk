package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestTestCancelSpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000160"
	xmhsdk.AppSecret = "xzs6We8YcKdpHrGQn8XHpyGCNg0a7sd7"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{ServiceOrderId: "20250314XSP83011364A832"},
		},
	}
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
			{ServiceOrderId: "20250610XPP11676136F0F65"},
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
