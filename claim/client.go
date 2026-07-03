package claim

import xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"

var claimItemsQueryUrl = "ClaimItemsQuery"
var claimUrl = "ClaimReport"
var serviceClaimItemsQueryUrl = "ServiceClaimItemQuery"
var claimQueryUrl = "ClaimQuery"
var claimQueryByOrderUrl = "ClaimQueryByOrder"

func ClaimItemsQuery(params *xmhsdk.ClaimItemsQueryParam) (*xmhsdk.ClaimItems, error) {
	result := &xmhsdk.ClaimItems{}
	err := xmhsdk.MakeRequest(claimItemsQueryUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return nil, err
	}
	return result, nil
}

func ServiceClaimItemsQuery(params *xmhsdk.ServiceClaimItemsQueryParam) (*xmhsdk.ClaimItems, error) {
	result := &xmhsdk.ServiceClaimItemsQueryResult{}
	err := xmhsdk.MakeRequest(serviceClaimItemsQueryUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return nil, err
	}
	return &result.ClaimItems, nil
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

func ClaimQuery(params *xmhsdk.ClaimQueryParam) (*xmhsdk.ClaimItems, error) {
	result := &xmhsdk.ClaimItems{}
	err := xmhsdk.MakeRequest(claimQueryUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return nil, err
	}
	return result, nil
}

// ClaimQueryByOrder 按订单号查询该订单下所有保单的理赔记录
func ClaimQueryByOrder(params *xmhsdk.ClaimQueryByOrderParam) (*xmhsdk.ClaimQueryByOrderResult, error) {
	result := &xmhsdk.ClaimQueryByOrderResult{}
	err := xmhsdk.MakeRequest(claimQueryByOrderUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimQueryByOrder error: %v", err)
		return nil, err
	}
	return result, nil
}
