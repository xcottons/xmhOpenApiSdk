package calc

import (
	"testing"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
)

func TestSpCalc(t *testing.T) {
	xmhsdk.AppId = "1601704"
	xmhsdk.AppSecret = "NAXnV842UJSjAWqhq53da9sw78JiranX"
	xmhsdk.SignSecret = "ueeshop-idc"
	xmhsdk.SetEnv(xmhsdk.EnvIdc)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		CartToken:                "11111111122222222222222223333333333jjjjjjjjjaaaaaaaa",
		IsShippingProtectionOpen: true,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice:     "13",
			TotalPrice:        "13",
			PreferentialPrice: "0",
			Currency:          "AUD",
		},
	}
	calc, err := Calc(params)
	if err != nil {
		t.Errorf("calc price error: %s", err)
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}

func TestPpCalc(t *testing.T) {
	xmhsdk.AppId = "1601056"
	xmhsdk.AppSecret = "tt8ZfdeUi5UrMqHVT17b9NUj9O4qAvcP"
	xmhsdk.SignSecret = "crealityusa-idc"
	xmhsdk.SetEnv(xmhsdk.EnvIdc)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		IsShippingProtectionOpen: false,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice:     "150.00",
			Currency:          "USD",
			PreferentialPrice: "0.0",
			TotalPrice:        "0.0",
			ItemList: []*xmhsdk.DItem{
				{
					ItemId:            "2026010901-US-1",
					SkuId:             "2026010901-US",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "JPY",
					UnitPrice:         "100.00",
					UnitNum:           "2",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
					//VariantId:         "123",
				},
				//{
				//	ItemId:            "2026010901-US",
				//	SkuId:             "2026010901-US",
				//	ItemName:          "Durable Concrete Shirt",
				//	Currency:          "JPY",
				//	UnitPrice:         "100.00",
				//	UnitNum:           "2",
				//	TotalPrice:        "200.00",
				//	PreferentialPrice: "50.00",
				//	TotalPayPrice:     "150.00",
				//	//VariantId:         "123",
				//}
				{
					ItemId: "1002110184",
					SkuId:  "1002110184",
					//VariantId:         "53226298835309",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "JPY",
					UnitPrice:         "100.00",
					UnitNum:           "2",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
				},
			},
		},
	}
	calc, err := Calc(params)
	if err != nil {
		t.Errorf("calc price error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}
func TestShopifyPpCalc(t *testing.T) {
	xmhsdk.AppId = "1000146"
	xmhsdk.AppSecret = "oOBzwiEPLIx33dvQfbIS5iR0GN94mFv2"
	xmhsdk.SignSecret = ""
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		IsShippingProtectionOpen: false,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice:     "150.00",
			Currency:          "USD",
			PreferentialPrice: "0.0",
			TotalPrice:        "0.0",
			ItemList: []*xmhsdk.DItem{{
				ItemId:            "7598670512227",
				SkuId:             "SKU001",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "JPY",
				UnitPrice:         "100.00",
				UnitNum:           "2",
				TotalPrice:        "200.00",
				PreferentialPrice: "50.00",
				TotalPayPrice:     "150.00",
				//VariantId:         "123",
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
	xmhsdk.AppId = "1601056"
	xmhsdk.AppSecret = "tt8ZfdeUi5UrMqHVT17b9NUj9O4qAvcP"
	xmhsdk.SignSecret = "crealityusa-idc"
	xmhsdk.SetEnv(xmhsdk.EnvIdc)
	auth.New()
	params := &xmhsdk.PpCalcParams{
		/*
		  "currency" : "CAD",
		    "itemId" : "3301120004",
		    "itemName" : "CR-Silk PLA",
		    "skuId" : "3301120004",
		    "unitPrice" : "1.00",
		    "variantId" : "50350464827692"
		*/
		ItemList: []*xmhsdk.DItem{
			//	{
			//	ItemId:            "2026010901-US",
			//	SkuId:             "2026010901-US",
			//	ItemName:          "Durable Concrete Shirt",
			//	Currency:          "JPY",
			//	UnitPrice:         "100.00",
			//	UnitNum:           "2",
			//	TotalPrice:        "200.00",
			//	PreferentialPrice: "50.00",
			//	TotalPayPrice:     "150.00",
			//	//VariantId:         "123",
			//}, {
			//	ItemId:            "3301120004",
			//	SkuId:             "3301120004",
			//	VariantId:         "50350464827692",
			//	ItemName:          "Durable Concrete Shirt",
			//	Currency:          "JPY",
			//	UnitPrice:         "100.00",
			//	UnitNum:           "2",
			//	TotalPrice:        "200.00",
			//	PreferentialPrice: "50.00",
			//	TotalPayPrice:     "150.00",
			//},
			{
				ItemId:            "1002110184",
				SkuId:             "1002110184",
				VariantId:         "50350464827692",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "JPY",
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

func TestEuSpCalc(t *testing.T) {
	xmhsdk.AppId = "10001"
	xmhsdk.AppSecret = "sX6QltIfMJEobQqxsdRFt3w2Jr8jlZaM"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	params := &xmhsdk.CalcParams{
		UserID:                   "VIPER3",
		UserEmail:                "VIPER3@qq.com",
		BuyerIP:                  "113.89.35.162",
		CartToken:                "11111111122222222222222223333333333jjjjjjjjjaaaaaaaa",
		IsShippingProtectionOpen: true,
		OrderInfo: &xmhsdk.DOrder{
			TotalPayPrice:     "303.00",
			TotalPrice:        "303.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
		},
	}
	calc, err := Calc(params)
	if err != nil {
		t.Errorf("calc price error: %s", err)
	}
	t.Logf("calc price result: %s", xmhsdk.ToStr(calc))
}

func TestEuPpCalc(t *testing.T) {
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

func TestEuOnlyPpCalc(t *testing.T) {
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
