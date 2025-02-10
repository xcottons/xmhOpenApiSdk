package item

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestSync(t *testing.T) {
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New(&xmhsdk.AuthParam{
		AppId:     "1000151",
		AppSecret: "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p",
	})
	params := &xmhsdk.ProductItemsParam{
		Items: []*xmhsdk.ProductItem{
			{
				PlatformItemId: "123",
				ItemState:      1,
				ItemName:       "TestProductItem",
				Currency:       "CNY",
				ItemPriceExt:   "123",
				ItemDetailExt:  "123",
				PicLink:        []string{"https://www.google.com"},
				ItemLink:       "https://www.google.com",
				Variants:       nil,
			},
		},
	}
	_, err := Sync(params)
	if err != nil {
		panic(err)
	}
}

func TestSyncWithVariants(t *testing.T) {
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New(&xmhsdk.AuthParam{
		AppId:     "1000151",
		AppSecret: "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p",
	})
	params := &xmhsdk.ProductItemsParam{
		Items: []*xmhsdk.ProductItem{
			{
				PlatformItemId: "123",
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
						Price:     "12.00",
						ImageUrl:  "https://www.google.com",
					},
				},
			},
			{
				PlatformItemId: "123",
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
					Price:     "13.00",
					ImageUrl:  "https://www.google.com",
				}},
			},
		},
	}
	_, err := Sync(params)
	if err != nil {
		panic(err)
	}
}
