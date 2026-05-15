package ship

import (
	"encoding/json"
	"os"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformShipURL = "SyncPlatformShip"

func New(params *xmhsdk.ShipParam) (*xmhsdk.ShipResult, error) {
	result := &xmhsdk.ShipResult{}
	err := xmhsdk.MakeRequest(platformShipURL, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Batch(params []*xmhsdk.ShipParam) error {
	for i, param := range params {
		_, err := New(param)
		if err != nil {
			xmhsdk.Logger.Errorf("batch ship idx:%d error:%s", i, err.Error())
			return err
		}
	}
	return nil
}
func BatchFromJsonFile(filePath string) error {
	bytes, rErr := os.ReadFile(filePath)
	if rErr != nil {
		xmhsdk.Logger.Errorf("read json file:%s", rErr.Error())
		return rErr
	}
	var params []*xmhsdk.ShipParam
	uErr := json.Unmarshal(bytes, &params)
	if uErr != nil {
		xmhsdk.Logger.Errorf("decode json file error:%s", uErr.Error())
		return uErr
	}
	return Batch(params)
}
