package calc

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestSpCalc(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		CartToken:                "11111111122222222222222223333333333jjjjjjjjjaaaaaaaa",
		IsShippingProtectionOpen: true,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice: "100.00",
			Currency:      "USD",
		},
	}
	calc, err := Calc(params)
	if err != nil {
		t.Errorf("calc price error: %s", err)
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}

func TestPpCalc(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		IsShippingProtectionOpen: true,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice:     "150.00",
			Currency:          "USD",
			PreferentialPrice: "0.0",
			TotalPrice:        "0.0",
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
	calc, err := Calc(params)
	if err != nil {
		t.Errorf("calc price error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}

func TestOnlyPpCalc(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PpCalcParams{
		ItemList: []*xmhsdk.DItem{{
			ItemId: "16GA5QRB5TO01_110",
			//SkuId:             "SKU001",
			ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
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
	}
	calc, err := PpCalc(params)
	if err != nil {
		t.Errorf("calc price error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}
