package claim

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"testing"
)

func TestClaimItemsQuery(t *testing.T) {
	params := &xmhsdk.ClaimItemsQueryParam{
		OrderID:            "",
		ClaimInsuranceType: 1,
	}
	result, err := ClaimItemsQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimItemsQuery result: %v", result)
}

func TestClaimReport(t *testing.T) {
	params := &xmhsdk.OpenApiClaimReport{
		OrderID: "",
		//
	}
	result, err := ClaimReport(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimReport error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimReport result: %v", result)
}
