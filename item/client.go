package item

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformItemURL = "SyncPlatformProduct"

func Sync(params *xmhsdk.ProductItemsParam) error {

	result := &xmhsdk.PlatformOrderResult{}
	if err := xmhsdk.MakeRequest(platformItemURL, params, result); err != nil {
		return err
	}
	return nil
}
