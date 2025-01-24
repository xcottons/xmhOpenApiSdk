package item

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformItemURL = "SyncPlatformProduct"

func Sync(params *xmhsdk.ProductItemsParam) (*struct{}, error) {
	result := &xmhsdk.PlatformOrderResult{}
	err := xmhsdk.MakeRequest(platformItemURL, params, result)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
