package item

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ProductItemsParam{
		RequestId: strconv.Itoa(int(time.Now().Unix())),
		Items: []*xmhsdk.ProductItem{
			{
				PlatformItemId: "16GA5QRB5TO01_110",
				ItemState:      1,
				ItemName:       "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency:       "CAD",
				ItemPriceExt:   "{}",
				ItemDetailExt:  "{}",
				PicLink:        []string{"https://www.google.com"},
				ItemLink:       "https://www.google.com",
				Variants:       nil,
				SkuId:          "",
				Price:          "1199.99",
			},
		},
	}
	err := Sync(params)
	if err != nil {
		panic(err)
	}
}

func TestSyncWithVariants(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ProductItemsParam{
		Items: []*xmhsdk.ProductItem{
			{
				PlatformItemId: "PlatformItemId1112",
				ItemState:      1,
				ItemName:       "TestProductItem1",
				Currency:       "CNY",
				ItemPriceExt:   "123",
				ItemDetailExt:  "123",
				PicLink:        []string{"https://www.google.com"},
				ItemLink:       "https://www.google.com",
				Variants: []xmhsdk.OpenApiPlatformProductItemVariant{
					{
						VariantId: "TestProductItemV1",
						Name:      "TestProductItemVariant1",
						SkuId:     "TestProductItemS1",
						Price:     "1.00",
						ImageUrl:  "https://www.google.com",
					},
					{
						VariantId: "TestProductItemV2",
						Name:      "TestProductItemVariant1",
						SkuId:     "TestProductItemS2",
						Price:     "1.00",
						ImageUrl:  "https://www.google.com",
					},
				},
			},
			{
				PlatformItemId: "PlatformItemId123",
				ItemState:      1,
				ItemName:       "TestProductItem2",
				Currency:       "CNY",
				ItemPriceExt:   "123",
				ItemDetailExt:  "123",
				PicLink:        []string{"https://www.google.com"},
				ItemLink:       "https://www.google.com",
				Variants: []xmhsdk.OpenApiPlatformProductItemVariant{{
					VariantId: "TestProductItemV2",
					Name:      "TestProductItemVariant2",
					SkuId:     "TestProductItemS2",
					Price:     "123.00",
					ImageUrl:  "https://www.google.com",
				}, {
					VariantId: "TestProductItemV1",
					Name:      "TestProductItemVariant1",
					SkuId:     "TestProductItemS1",
					Price:     "223.00",
					ImageUrl:  "https://www.google.com",
				}},
			},
		},
	}
	err := Sync(params)
	if err != nil {
		panic(err)
	}
}

func TestSyncWithFile(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ProductItemsParam{
		FileUrl: "https://sslstatic.xiaoyusan.com/img/hyju/item.b325b0e1753dc8ad.csv",
	}
	err := Sync(params)
	if err != nil {
		panic(err)
	}

}

func TestEuSync(t *testing.T) {
	xmhsdk.AppId = "10001"
	xmhsdk.AppSecret = "sX6QltIfMJEobQqxsdRFt3w2Jr8jlZaM"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvEuBeta)
	auth.New()
	params := &xmhsdk.ProductItemsParam{
		RequestId: strconv.Itoa(int(time.Now().Unix())),
		Items: []*xmhsdk.ProductItem{
			{
				PlatformItemId: "16GA5QRB5TO01_111",
				ItemState:      1,
				ItemName:       "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency:       "CAD",
				ItemPriceExt:   "{}",
				ItemDetailExt:  "{}",
				PicLink:        []string{"https://www.google.com"},
				ItemLink:       "https://www.google.com",
				Variants:       nil,
				SkuId:          "",
				Price:          "1199.99",
			},
		},
	}
	err := Sync(params)
	if err != nil {
		panic(err)
	}
}
