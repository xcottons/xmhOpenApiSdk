package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformOrderURL = "SyncPlatformOrder"
const cancelOrderURl = ""
const insureOrderURl = ""

func New(params *xmhsdk.PlatformOrderParam) (*xmhsdk.PlatformOrderResult, error) {
	result := &xmhsdk.PlatformOrderResult{}
	err := xmhsdk.MakeRequest(platformOrderURL, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Cancel(params *xmhsdk.CancelOrderParam) (*xmhsdk.CancelOrderResult, error) {
	result := &xmhsdk.CancelOrderResult{}
	err := xmhsdk.MakeRequest(cancelOrderURl, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Insured(params *xmhsdk.InsuredOrderParam) (*xmhsdk.InsuredOrderResult, error) {
	result := &xmhsdk.InsuredOrderResult{}
	err := xmhsdk.MakeRequest(insureOrderURl, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
