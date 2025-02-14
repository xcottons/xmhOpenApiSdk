package claim

import xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"

var claimItemsQueryUrl = "ClaimItemsQuery"
var claimUrl = "ClaimReport"

func ClaimItemsQuery(params *xmhsdk.ClaimItemsQueryParam) (*xmhsdk.ClaimItems, error) {
	result := &xmhsdk.ClaimItems{}
	err := xmhsdk.MakeRequest(claimItemsQueryUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return nil, err
	}
	return result, nil
}

func ClaimReport(params *xmhsdk.OpenApiClaimReport) (*xmhsdk.OpenApiClaimReportResult, error) {
	result := &xmhsdk.OpenApiClaimReportResult{}
	err := xmhsdk.MakeRequest(claimUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimReport error: %v", err)
		return nil, err
	}
	return result, nil
}
