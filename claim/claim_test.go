package claim

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
	"time"
)

func TestClaimItemsQuery(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ClaimItemsQueryParam{
		OrderID:            "1744182466",
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
		ServiceOrderId: "20250407XPP11676136C5565",
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
		ServiceOrderId: "20250407XPP11676136C6665",
		ClaimItems: []*xmhsdk.OpenApiClaimItem{{
			ItemId:          "16GA5QRB5TO01_110",
			ItemName:        "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
			PriceCurrency:   "USD",
			ItemNum:         1,
			ItemUnitPrice:   "100.00",
			ItemSumPrice:    "200.00",
			ClaimApplyMoney: "100",
			PlanId:          xmhsdk.PPPlanCodeFor1Years,
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
		Claimpayout:        "100",
	}
	result, err := ClaimReport(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimReport error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimReport result: %v", result)
}

func TestClaimQuery(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ClaimQueryParam{
		ServiceOrderId: "20250407XPP11676136C6665",
	}
	result, err := ClaimQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimQuery result: %s", xmhsdk.ToStr(result))
}
