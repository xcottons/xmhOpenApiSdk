package claim

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	orderapi "github.com/cjay-shouhui/xmhOpenApiSdk/order"
	shipapi "github.com/cjay-shouhui/xmhOpenApiSdk/ship"
)

const claimPipelineBatchDefaultSize = 10
const claimQueryRetryCount = 6
const claimQueryRetryInterval = 2 * time.Second

func initAlphaEnvForClaim(t *testing.T) {
	t.Helper()
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	if _, err := auth.New(); err != nil {
		t.Fatalf("auth.New error: %v", err)
	}
}

func getClaimBatchSizeFromEnv() int {
	raw := os.Getenv("XMH_CLAIM_BATCH_SIZE")
	if raw == "" {
		return claimPipelineBatchDefaultSize
	}
	size, err := strconv.Atoi(raw)
	if err != nil || size <= 0 {
		return claimPipelineBatchDefaultSize
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

func queryClaimableItemByOrder(orderID string) (*xmhsdk.OpenApiClaimItem, error) {
	var lastErr error
	for i := 0; i < claimQueryRetryCount; i++ {
		items, err := ClaimItemsQuery(&xmhsdk.ClaimItemsQueryParam{
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

func queryClaimableItemByServiceOrder(serviceOrderID string) (*xmhsdk.OpenApiClaimItem, error) {
	var lastErr error
	for i := 0; i < claimQueryRetryCount; i++ {
		items, err := ServiceClaimItemsQuery(&xmhsdk.ServiceClaimItemsQueryParam{
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
		CompensationReason: 5,
	}
}

func TestBatchOrderShipClaimPipeline(t *testing.T) {
	initAlphaEnvForClaim(t)

	batchSize := getClaimBatchSizeFromEnv()
	success := 0
	for i := 0; i < batchSize; i++ {
		orderID := strconv.FormatInt(time.Now().UnixNano()+int64(i), 10)
		itemID := fmt.Sprintf("claim-batch-item-%d", i)
		orderParam := buildClaimPipelineOrderParam(orderID, itemID)

		orderResult, orderErr := orderapi.New(orderParam)
		if orderErr != nil {
			t.Errorf("pipeline order failed idx=%d orderId=%s err=%v", i, orderID, orderErr)
			continue
		}

		shipResult, shipErr := shipapi.New(buildClaimPipelineShipParam(orderResult, orderParam, i))
		if shipErr != nil {
			t.Errorf("pipeline ship failed idx=%d orderId=%s err=%v", i, orderResult.OrderId, shipErr)
			continue
		}

		serviceOrderID := orderResult.Insurance.SPInsureDetail.ServiceOrderID
		if serviceOrderID == "" {
			claimableItemByOrder, claimableErr := queryClaimableItemByOrder(shipResult.OrderId)
			if claimableErr != nil {
				t.Errorf("pipeline serviceOrderId empty and order claim query failed idx=%d orderId=%s err=%v", i, shipResult.OrderId, claimableErr)
				continue
			}
			serviceOrderID = claimableItemByOrder.XmhServiceId
		}
		if serviceOrderID == "" {
			t.Errorf("pipeline serviceOrderId still empty idx=%d orderId=%s", i, shipResult.OrderId)
			continue
		}

		claimableItem, claimableErr := queryClaimableItemByServiceOrder(serviceOrderID)
		if claimableErr != nil {
			t.Errorf("pipeline service claim item query failed idx=%d orderId=%s serviceOrderId=%s err=%v", i, shipResult.OrderId, serviceOrderID, claimableErr)
			continue
		}

		claimResult, claimErr := ClaimReport(buildClaimReportParam(serviceOrderID, claimableItem, orderParam.OrderInfo.ItemList[0]))
		if claimErr != nil {
			t.Errorf("pipeline claim report failed idx=%d orderId=%s serviceOrderId=%s err=%v", i, shipResult.OrderId, serviceOrderID, claimErr)
			continue
		}

		t.Logf("pipeline success idx=%d orderId=%s serviceOrderId=%s claimId=%s claimIds=%v", i, shipResult.OrderId, serviceOrderID, claimResult.ClaimId, claimResult.ClaimIds)
		success++
	}

	if success == 0 {
		t.Fatalf("pipeline failed: no success in %d attempts", batchSize)
	}
}

func TestClaimItemsQuery(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ClaimItemsQueryParam{
		OrderID:            "1770265268",
		ClaimInsuranceType: 4,
	}
	result, err := ClaimItemsQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimItemsQuery result: %s", xmhsdk.ToStr(result))
}

func TestServiceClaimItemsQuery(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ServiceClaimItemsQueryParam{
		ServiceOrderId: "20260205XPP1167613859465",
	}
	result, err := ServiceClaimItemsQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("TestServiceClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("TestServiceClaimItemsQuery result: %s", xmhsdk.ToStr(result))
}

func TestClaimReport(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.OpenApiClaimReport{
		OrderGoodsId:   "123",
		RefundId:       "456",
		ServiceOrderId: "20260702XSP11676138A5B65",
		ClaimItems: []*xmhsdk.OpenApiClaimItem{{
			ItemId:        "16GA5QRB5TO01_110",
			ItemName:      "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
			PriceCurrency: "USD",
			ItemNum:       2,
			//ItemUnitPrice:   "100.00",
			//ItemSumPrice:    "200.00",
			Claimpayout: "1000",
		}},
		ClaimType:       1,
		FileLinks:       []string{"https://example.com"},
		Comments:        "some comments",
		ClaimPaymentObj: 1,
		PaymentMethod:   1,
		AccountInfo: xmhsdk.AccountInfo{
			AccountNumber: "1234",
			AccountName:   "123",
			RoutingNumber: "124435",
			Address:       "fwhrughryug",
		},
		ClaimReportTime:    time.Now().Format(time.RFC3339),
		LossOccurrenceTime: time.Now().Format(time.RFC3339),
		Describe:           "some describe",
	}
	result, err := ClaimReport(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimReport error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimReport result: %v", result)
}

func TestClaimQuery(t *testing.T) {
	xmhsdk.AppId = "1000170"
	xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	xmhsdk.SignSecret = "vevor_beta"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	auth.New()
	params := &xmhsdk.ClaimQueryParam{
		ServiceOrderId: "20251212XSP83015378DC36",
		ClaimId:        "claim-00232360000000000117705236",
	}
	result, err := ClaimQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimQuery result: %s", xmhsdk.ToStr(result))
}

func TestClaimQueryByOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ClaimQueryByOrderParam{
		OrderID:            "1782982705036830808",
		ClaimInsuranceType: 3,
		Region:             "CN",
	}
	result, err := ClaimQueryByOrder(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimQueryByOrder error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimQueryByOrder result: %s", xmhsdk.ToStr(result))
}
