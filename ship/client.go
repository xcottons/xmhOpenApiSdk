package ship

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformShipURL = "SyncPlatformShip"

func New(params *xmhsdk.ShipParam) (*xmhsdk.ShipResult, error) {
	result := &xmhsdk.ShipResult{}
	err := xmhsdk.MakeRequest(platformShipURL, params, result)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
