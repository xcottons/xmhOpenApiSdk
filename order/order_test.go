package order

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"github.com/cjay-shouhui/xmhOpenApiSdk/claim"
	shipapi "github.com/cjay-shouhui/xmhOpenApiSdk/ship"
)

const claimPipelineDefaultBatchSize = 100
const claimQueryRetryCount = 6
const claimQueryRetryInterval = 2 * time.Second

func initAlphaEnv(t *testing.T) {
	t.Helper()
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	if _, err := auth.New(); err != nil {
		t.Fatalf("auth.New error: %v", err)
	}
}

func getBatchSizeFromEnv() int {
	raw := os.Getenv("XMH_CLAIM_BATCH_SIZE")
	if raw == "" {
		return claimPipelineDefaultBatchSize
	}
	size, err := strconv.Atoi(raw)
	if err != nil || size <= 0 {
		return claimPipelineDefaultBatchSize
	}
	return size
}

func buildClaimPipelineOrderParam(orderID string, itemID string) *xmhsdk.PlatformOrderParam {
	now := time.Now()
	return &xmhsdk.PlatformOrderParam{
		UserEmail:    "yujianfx@xcotton.cn",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:           orderID,
			SubOrderId:        orderID,
			TotalPayPrice:     "153.00",
			TotalPrice:        "150.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "3.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           now.Format(time.RFC3339),
			PaySn:             orderID,
			OrderModifyTime:   now.Format(time.RFC3339),
			OrderCreateTime:   now.Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "United States",
					CityCode: "US",
					ZipCode:  "10001",
				},
			},
			ItemList: []*xmhsdk.DItem{
				{
					ItemId:            itemID,
					SkuId:             "SKU-CLAIM-BATCH-1",
					ItemName:          "Claim Pipeline Item",
					Currency:          "USD",
					UnitPrice:         "150.00",
					UnitNum:           "1",
					TotalPrice:        "150.00",
					PreferentialPrice: "0.00",
					TotalPayPrice:     "150.00",
				},
			},
		},
	}
}

func buildClaimPipelineShipParam(orderRes *xmhsdk.PlatformOrderResult, orderParam *xmhsdk.PlatformOrderParam, idx int) *xmhsdk.ShipParam {
	now := time.Now()
	item := orderParam.OrderInfo.ItemList[0]
	return &xmhsdk.ShipParam{
		OrderId:    orderRes.OrderId,
		SubOrderId: orderRes.OrderId,
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             fmt.Sprintf("%s-%d", orderRes.OrderId, idx),
				ShipCompanyCode:    "UPS",
				ShipCompany:        "UPS",
				ShipTrackNumber:    fmt.Sprintf("TRACK-%s-%d", orderRes.OrderId, idx),
				ShipStateString:    "SHIPPED",
				ShipPrice:          "0.00",
				ActualShipSendTime: now.Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:    "United States",
					CityCode:   "US",
					Province:   "California",
					City:       "Los Angeles",
					PostalCode: "90001",
				},
				ItemList: []*xmhsdk.DItem{
					{
						ItemId:        item.ItemId,
						SkuId:         item.SkuId,
						ItemName:      item.ItemName,
						Currency:      item.Currency,
						UnitPrice:     item.UnitPrice,
						UnitNum:       item.UnitNum,
						TotalPrice:    item.TotalPrice,
						TotalPayPrice: item.TotalPayPrice,
					},
				},
			},
		},
	}
}

func queryClaimableItemByOrderID(orderID string) (*xmhsdk.OpenApiClaimItem, error) {
	var lastErr error
	for i := 0; i < claimQueryRetryCount; i++ {
		items, err := claim.ClaimItemsQuery(&xmhsdk.ClaimItemsQueryParam{
			OrderID:            orderID,
			ClaimInsuranceType: 4,
		})
		if err != nil {
			lastErr = err
			time.Sleep(claimQueryRetryInterval)
			continue
		}
		if len(*items) > 0 {
			return (*items)[0], nil
		}
		lastErr = fmt.Errorf("no claim items found")
		time.Sleep(claimQueryRetryInterval)
	}
	return nil, fmt.Errorf("query claim item failed for order %s: %v", orderID, lastErr)
}

func queryClaimableItemByServiceOrderID(serviceOrderID string) (*xmhsdk.OpenApiClaimItem, error) {
	var lastErr error
	for i := 0; i < claimQueryRetryCount; i++ {
		items, err := claim.ServiceClaimItemsQuery(&xmhsdk.ServiceClaimItemsQueryParam{
			ServiceOrderId: serviceOrderID,
		})
		if err != nil {
			lastErr = err
			time.Sleep(claimQueryRetryInterval)
			continue
		}
		if len(*items) > 0 {
			return (*items)[0], nil
		}
		lastErr = fmt.Errorf("no claim items found")
		time.Sleep(claimQueryRetryInterval)
	}
	return nil, fmt.Errorf("query claim item failed for serviceOrderId %s: %v", serviceOrderID, lastErr)
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func buildClaimReportParam(serviceOrderID string, claimItem *xmhsdk.OpenApiClaimItem, orderItem *xmhsdk.DItem) *xmhsdk.OpenApiClaimReport {
	now := time.Now().Format(time.RFC3339)
	claimType := int(claimItem.ClaimType)
	if claimType <= 0 {
		claimType = 1
	}

	priceCurrency := claimItem.PriceCurrency
	if priceCurrency == "" && orderItem != nil {
		priceCurrency = orderItem.Currency
	}
	if priceCurrency == "" {
		priceCurrency = "USD"
	}

	itemUnitPrice := claimItem.ItemUnitPrice
	itemSumPrice := claimItem.ItemSumPrice
	if orderItem != nil {
		if itemUnitPrice == "" {
			itemUnitPrice = firstNonEmpty(orderItem.UnitPrice, orderItem.TotalPayPrice)
		}
		if itemSumPrice == "" {
			itemSumPrice = firstNonEmpty(orderItem.TotalPayPrice, orderItem.TotalPrice, orderItem.UnitPrice)
		}
	}

	claimPayout := claimItem.ClaimApplyMoney
	if claimPayout == "" {
		claimPayout = claimItem.Claimpayout
	}
	if claimPayout == "" {
		claimPayout = itemSumPrice
	}
	if claimPayout == "" {
		claimPayout = itemUnitPrice
	}
	if claimPayout == "" {
		claimPayout = "50.00"
	}

	itemNum := int32(1)
	if claimItem.MaxClaimNum > 0 && itemNum > claimItem.MaxClaimNum {
		itemNum = claimItem.MaxClaimNum
	}
	if orderItem != nil {
		if n, err := strconv.Atoi(orderItem.UnitNum); err == nil && n > 0 && claimItem.MaxClaimNum > 0 && int32(n) < claimItem.MaxClaimNum {
			itemNum = int32(n)
		}
	}

	productId := claimItem.ProductId
	itemId := claimItem.ItemId
	skuId := claimItem.SkuId
	orderGoodsId := claimItem.OrderGoodsId
	variantId := claimItem.VariantId
	itemName := claimItem.ItemName
	if orderItem != nil {
		productId = firstNonEmpty(productId, orderItem.ProductId)
		itemId = firstNonEmpty(itemId, orderItem.ItemId)
		skuId = firstNonEmpty(skuId, orderItem.SkuId)
		orderGoodsId = firstNonEmpty(orderGoodsId, orderItem.OrderGoodsId)
		variantId = firstNonEmpty(variantId, orderItem.VariantId)
		itemName = firstNonEmpty(itemName, orderItem.ItemName)
	}

	reportItem := &xmhsdk.OpenApiClaimItem{
		ProductId:     productId,
		ItemId:        itemId,
		SkuId:         skuId,
		OrderGoodsId:  orderGoodsId,
		PlanId:        claimItem.PlanId,
		VariantId:     variantId,
		ItemNum:       itemNum,
		ItemName:      itemName,
		ItemUnitPrice: itemUnitPrice,
		ItemSumPrice:  itemSumPrice,
		PriceCurrency: priceCurrency,
		Claimpayout:   claimPayout,
		XmhServiceId:  serviceOrderID,
	}

	return &xmhsdk.OpenApiClaimReport{
		ServiceOrderId:  serviceOrderID,
		ClaimItems:      []*xmhsdk.OpenApiClaimItem{reportItem},
		ClaimType:       claimType,
		FileLinks:       []string{"https://example.com/claim-proof.jpg"},
		Comments:        "batch claim from sdk pipeline test",
		ClaimPaymentObj: 1,
		PaymentMethod:   1,
		AccountInfo: xmhsdk.AccountInfo{
			AccountName:   "TestUser",
			AccountNumber: "1234567890",
			RoutingNumber: "021000021",
			Address:       "100 Main St, Los Angeles, CA",
		},
		ClaimReportTime:    now,
		LossOccurrenceTime: now,
		Describe:           "Auto generated batch claim",
	}
}

func TestBatchSyncOrderForClaimPipeline(t *testing.T) {
	initAlphaEnv(t)

	batchSize := getBatchSizeFromEnv()
	success := 0
	for i := 0; i < batchSize; i++ {
		orderID := strconv.FormatInt(time.Now().UnixNano()+int64(i), 10)
		itemID := fmt.Sprintf("claim-batch-item-%d", i)
		orderParam := buildClaimPipelineOrderParam(orderID, itemID)

		// Step 1: sync order
		orderResult, orderErr := New(orderParam)
		if orderErr != nil {
			t.Errorf("pipeline order failed idx=%d orderId=%s err=%v", i, orderID, orderErr)
			continue
		}

		// Step 2: sync shipping
		shipResult, shipErr := shipapi.New(buildClaimPipelineShipParam(orderResult, orderParam, i))
		if shipErr != nil {
			t.Errorf("pipeline ship failed idx=%d orderId=%s err=%v", i, orderResult.OrderId, shipErr)
			continue
		}

		// Step 3: get serviceOrderId and query claimable item
		serviceOrderID := orderResult.Insurance.SPInsureDetail.ServiceOrderID
		if serviceOrderID == "" {
			claimableItem, claimableErr := queryClaimableItemByOrderID(shipResult.OrderId)
			if claimableErr != nil {
				t.Errorf("pipeline serviceOrderId empty and order claim query failed idx=%d orderId=%s err=%v", i, shipResult.OrderId, claimableErr)
				continue
			}
			serviceOrderID = claimableItem.XmhServiceId
		}
		if serviceOrderID == "" {
			t.Errorf("pipeline serviceOrderId still empty idx=%d orderId=%s", i, shipResult.OrderId)
			continue
		}

		claimableItem, claimableErr := queryClaimableItemByServiceOrderID(serviceOrderID)
		if claimableErr != nil {
			t.Errorf("pipeline service claim item query failed idx=%d orderId=%s serviceOrderId=%s err=%v", i, shipResult.OrderId, serviceOrderID, claimableErr)
			continue
		}
		t.Logf("pipeline claimable item idx=%d productId=%s itemId=%s itemName=%s itemUnitPrice=%s itemSumPrice=%s claimType=%d",
			i, claimableItem.ProductId, claimableItem.ItemId, claimableItem.ItemName, claimableItem.ItemUnitPrice, claimableItem.ItemSumPrice, claimableItem.ClaimType)

		// Step 4: submit claim report
		claimResult, claimErr := claim.ClaimReport(buildClaimReportParam(serviceOrderID, claimableItem, orderParam.OrderInfo.ItemList[0]))
		if claimErr != nil {
			t.Errorf("pipeline claim report failed idx=%d orderId=%s serviceOrderId=%s err=%v", i, shipResult.OrderId, serviceOrderID, claimErr)
			continue
		}

		// Step 5: query the submitted claim to verify items
		time.Sleep(2 * time.Second)
		claimItems, queryErr := claim.ClaimQuery(&xmhsdk.ClaimQueryParam{
			ServiceOrderId: serviceOrderID,
			ClaimId:        claimResult.ClaimId,
		})
		if queryErr != nil {
			t.Errorf("pipeline claim query failed idx=%d claimId=%s err=%v", i, claimResult.ClaimId, queryErr)
			continue
		}
		if len(*claimItems) == 0 {
			t.Errorf("pipeline claim query returned no items idx=%d claimId=%s", i, claimResult.ClaimId)
			continue
		}
		for j, ci := range *claimItems {
			t.Logf("pipeline claim item[%d] idx=%d productId=%s itemId=%s itemName=%s itemUnitPrice=%s itemSumPrice=%s claimType=%d claimState=%d",
				j, i, ci.ProductId, ci.ItemId, ci.ItemName, ci.ItemUnitPrice, ci.ItemSumPrice, ci.ClaimType, ci.ClaimState)
		}

		t.Logf("pipeline success idx=%d orderId=%s serviceOrderId=%s claimId=%s claimIds=%v", i, shipResult.OrderId, serviceOrderID, claimResult.ClaimId, claimResult.ClaimIds)
		success++
	}

	if success == 0 {
		t.Fatalf("pipeline failed: no success in %d attempts", batchSize)
	}
}

func TestSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	//xmhsdk.AppId = "1000170"
	//xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	//xmhsdk.SignSecret = "vevor_beta"
	//xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "yujianfx@xcotton.cn",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
		OrderInfo: &xmhsdk.DOrder{
			OrderId: strconv.Itoa(int(time.Now().Unix())),
			//SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "310.00",
			TotalPrice:        "0.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "10.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "300.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Add(-8 * time.Hour).Format(time.RFC3339),
			//OrderCreateTime:   time.Now().In(time.UTC).Add(-8 * time.Hour).Format(time.RFC3339),
			OrderCreateTime: time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					ZipCode:  "10086",
					//City:     "亚特兰提斯",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "97760109-16ad-40c9-9385-caba381a26aa",
				SkuId:  "SKU001",
				//OrderGoodsId:      "SKU001000001",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "USD",
				UnitPrice:         "0",
				UnitNum:           "1",
				TotalPrice:        "150",
				PreferentialPrice: "0",
				TotalPayPrice:     "150.00",
			},
				{
					ItemId: "97760109-16ad-40c9-9385-caba381a26aa2",
					SkuId:  "SKU002",
					//OrderGoodsId:      "SKU001000002",
					ItemName:          "Durable Concrete Shirt2",
					Currency:          "USD",
					UnitPrice:         "0",
					UnitNum:           "2",
					TotalPrice:        "150",
					PreferentialPrice: "0",
					TotalPayPrice:     "150.00",
				},
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
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
		UserId:    "VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId: strconv.Itoa(int(time.Now().Unix())),
			//OrderId: "1772696426",
			//SubOrderId:    "SubOrderId6",
			TotalPayPrice: "3003.00",
			Currency:      "USD",
			OrderState:    xmhsdk.OERDER_STATE_PAID,
			//InsuredPayPrice:   "10",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					ZipCode:  "996",
				},
			},
			ItemList: []*xmhsdk.DItem{
				{
					ItemId:          "16GA5QRB5TO01_110",
					ItemName:        "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
					Currency:        "USD",
					UnitPrice:       "48.00",
					UnitNum:         "5",
					TotalPayPrice:   "4500.00",
					InsuredPayPrice: "20.0",
					SumInsured:      "1000.00",
					PpVariant: &xmhsdk.PPVariant{
						VariantID:   "",
						Name:        "3 Years-back",
						PriceString: "10.00",
						Currency:    "USD",
						Properties: map[string]string{
							"Reference": "16GA5QRB5TO01_110",
							"Plan ID":   "3 Years-back",
							"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
						},
					},
				},
				//{
				//	ItemId:          "PlatformItemId1112",
				//	ItemName:        "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				//	Currency:        "USD",
				//	UnitPrice:       "48.00",
				//	UnitNum:         "5",
				//	TotalPayPrice:   "4500.00",
				//	InsuredPayPrice: "80.0",
				//	SumInsured:      "1000.00",
				//	PpVariant: &xmhsdk.PPVariant{
				//		VariantID:   "",
				//		Name:        xmhsdk.PPPlanCodeFor1Years,
				//		PriceString: "20.00",
				//		Currency:    "USD",
				//		Properties: map[string]string{
				//			"Reference": "PlatformItemId1112",
				//			"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
				//			"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				//		},
				//	},
				//},
			},
		},
	}
	pr, err := New(params)
	if err != nil {
		t.Errorf("pr error: %s", xmhsdk.ToStr(err))
	}
	t.Logf("pr result: %s", xmhsdk.ToStr(pr))
}
func TestShopifyPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000146"
	xmhsdk.AppSecret = "oOBzwiEPLIx33dvQfbIS5iR0GN94mFv2"
	xmhsdk.SignSecret = ""
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
		UserId:    "VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId:       strconv.Itoa(int(time.Now().Unix())),
			SubOrderId:    "SubOrderId6",
			TotalPayPrice: "3003.00",

			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "10",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					ZipCode:  "996",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "7598670512227",
				//SkuId:             "SKU001",
				//OrderGoodsId: "16GA5QRB5TO01_110",
				ItemName:  "The Collection Snowboard: Liquid",
				Currency:  "USD",
				UnitPrice: "712.45",
				UnitNum:   "2",
				//TotalPrice:        "200.00",
				//PreferentialPrice: "50.00",
				TotalPayPrice:   "1424.90",
				InsuredPayPrice: "36.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "18.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "16GA5QRB5TO01_110",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
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

func TestPpOrders(t *testing.T) {
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
	//p3 := &xmhsdk.DItem{
	//	ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
	//	SkuId:             "SKU001",
	//	ItemName:          "Durable Concrete Shirt",
	//	Currency:          "USD",
	//	UnitPrice:         "100.00",
	//	UnitNum:           "2",
	//	TotalPrice:        "200.00",
	//	PreferentialPrice: "50.00",
	//	TotalPayPrice:     "150.00",
	//}
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
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
			ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"),
				//p3.AddThreeYearPp("10.00"),
				{
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
		UserEmail: "yujianfx@xcotton.cn",
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
		UserEmail:    "yujianfx@xcotton.cn",
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
				ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("12.00")},
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
				ItemList: []*xmhsdk.DItem{p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00")},
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
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "yujianfx@xcotton.cn",
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
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
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

func TestPpOrderBeta(t *testing.T) {
	xmhsdk.AppId = "1000170"
	xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	xmhsdk.SignSecret = "vevor_beta"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
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
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
			ItemList: []*xmhsdk.DItem{{
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
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "VEVOR_US_1.5KWTSSGJ0000001V1",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
					},
				},
			}, {
				ItemId: "VEVOR_US_1.5KWTSSGJ0000001V1",
				//SkuId:             "SKU001",
				OrderGoodsId: "7316476058016931841",
				ItemName:     "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency:     "USD",
				//UnitPrice:         "100.00",
				UnitNum: "1",
				//TotalPrice:        "200.00",
				//PreferentialPrice: "50.00",
				TotalPayPrice:   "150.00",
				InsuredPayPrice: "24.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "VEVOR_US_1.5KWTSSGJ0000001V1",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
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

func TestEuSpOrder(t *testing.T) {
	xmhsdk.AppId = "10001"
	xmhsdk.AppSecret = "sX6QltIfMJEobQqxsdRFt3w2Jr8jlZaM"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "yujianfx@xcotton.cn",
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

func TestEuPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
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
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "16GA5QRB5TO01_110",
				//SkuId:             "SKU001",
				OrderGoodsId: strconv.Itoa(int(time.Now().UnixNano()) + 3),
				ItemName:     "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency:     "USD",
				//UnitPrice:         "100.00",
				UnitNum: "3",
				//TotalPrice:        "200.00",
				//PreferentialPrice: "50.00",
				TotalPayPrice:   "150.00",
				InsuredPayPrice: "24.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "16GA5QRB5TO01_110",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
					},
				},
			}, {
				ItemId: "16GA5QRB5TO01_110",
				//SkuId:             "SKU001",
				OrderGoodsId: strconv.Itoa(int(time.Now().UnixNano()) + 5),
				ItemName:     "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
				Currency:     "USD",
				//UnitPrice:         "100.00",
				UnitNum: "2",
				//TotalPrice:        "200.00",
				//PreferentialPrice: "50.00",
				TotalPayPrice:   "150.00",
				InsuredPayPrice: "24.0",
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "16GA5QRB5TO01_110",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
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

func TestUsBetaSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000160"
	xmhsdk.AppSecret = "xzs6We8YcKdpHrGQn8XHpyGCNg0a7sd7"
	xmhsdk.SignSecret = "ex_xmh_test"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	//xmhsdk.AppId = "1000170"
	//xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	//xmhsdk.SignSecret = "vevor_beta"
	//xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail:    "yujianfx@xcotton.cn",
		UserId:       "VIPER3",
		DisComputeId: "xcp-00222580000000000508829558",
		OrderInfo: &xmhsdk.DOrder{
			OrderId: strconv.Itoa(int(time.Now().Unix())),
			//SubOrderId:        "SubOrderId6",
			TotalPayPrice:     "310.00",
			TotalPrice:        "0.00",
			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "10.00",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Add(-8 * time.Hour).Format(time.RFC3339),
			//OrderCreateTime:   time.Now().In(time.UTC).Add(-8 * time.Hour).Format(time.RFC3339),
			OrderCreateTime: time.Now().Format(time.RFC3339),
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					ZipCode:  "10086",
					//City:     "亚特兰提斯",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "97760109-16ad-40c9-9385-caba381a26aa",
				SkuId:  "SKU001",
				//OrderGoodsId:      "SKU001000001",
				ItemName:          "Durable Concrete Shirt",
				Currency:          "USD",
				UnitPrice:         "0",
				UnitNum:           "1",
				TotalPrice:        "150",
				PreferentialPrice: "0",
				TotalPayPrice:     "150.00",
			},
				{
					ItemId: "97760109-16ad-40c9-9385-caba381a26aa2",
					SkuId:  "SKU002",
					//OrderGoodsId:      "SKU001000002",
					ItemName:          "Durable Concrete Shirt2",
					Currency:          "USD",
					UnitPrice:         "0",
					UnitNum:           "2",
					TotalPrice:        "150",
					PreferentialPrice: "0",
					TotalPayPrice:     "150.00",
				},
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

func TestUsBetaPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000160"
	xmhsdk.AppSecret = "xzs6We8YcKdpHrGQn8XHpyGCNg0a7sd7"
	xmhsdk.SignSecret = "ex_xmh_test"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	params := &xmhsdk.PlatformOrderParam{
		UserEmail: "yujianfx@xcotton.cn",
		UserId:    "VIPER3",
		OrderInfo: &xmhsdk.DOrder{
			OrderId: strconv.Itoa(int(time.Now().Unix())),
			//SubOrderId:    "SubOrderId6",
			TotalPayPrice: "3003.00",

			Currency:          "USD",
			OrderState:        xmhsdk.OERDER_STATE_PAID,
			InsuredPayPrice:   "10",
			TaxPrice:          "0.00",
			ShipPrice:         "0.00",
			PreferentialPrice: "0.00",
			PayTime:           time.Now().Format(time.RFC3339),
			PaySn:             time.Now().Format(time.RFC3339),
			OrderModifyTime:   time.Now().Format(time.RFC3339),
			OrderCreateTime:   time.Now().Format(time.RFC3339),
			//OrderCreateTime: "2025-04-07T05:37:10.000Z",
			ReceiverInfoDto: &xmhsdk.ReceiverInfoDto{
				ReceiverShipAddress: &xmhsdk.ShipAddress{
					Country:  "CN",
					CityCode: "CN",
					ZipCode:  "996",
				},
			},
			ItemList: []*xmhsdk.DItem{{
				ItemId: "1FCALEORCDO01_100",
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
				//Properties: map[string]string{
				//	"Plan ID": "10802003",
				//},
				SumInsured: "1000.00",
				PpVariant: &xmhsdk.PPVariant{
					VariantID:   "",
					Name:        xmhsdk.PPPlanCodeFor1Years,
					PriceString: "12.00",
					Currency:    "USD",
					Properties: map[string]string{
						"Reference": "16GA5QRB5TO01_110",
						"Plan ID":   xmhsdk.PPPlanCodeFor1Years,
						"Product":   "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
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

func TestHmac(t *testing.T) {
	insuredEmail := "pkasamson@yahoo.com"
	platformOrderNumber := "6376650113083"
	serviceOrder := "20260324XPP30615F8B8C15"
	hMac := hmac.New(sha256.New, []byte(insuredEmail))
	hMac.Write([]byte(fmt.Sprintf("%s+%s", serviceOrder, platformOrderNumber)))
	toString := hex.EncodeToString(hMac.Sum(nil))
	t.Log(toString)
}
