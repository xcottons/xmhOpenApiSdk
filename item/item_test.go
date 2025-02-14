package item

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestSync(t *testing.T) {
	xmhsdk.AppId = "1000151"
	xmhsdk.AppSecret = "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
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
	xmhsdk.AppId = "1000151"
	xmhsdk.AppSecret = "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p"
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
	_, err := Sync(params)
	if err != nil {
		panic(err)
	}
}

func TestSyncWithFile(t *testing.T) {
	xmhsdk.AppId = "1000151"
	xmhsdk.AppSecret = "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ProductItemsParam{
		FileUrl: "https://sslstatic.xiaoyusan.com/img/hyju/item.b325b0e1753dc8ad.csv",
	}
	_, err := Sync(params)
	if err != nil {
		panic(err)
	}

}
