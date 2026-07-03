package serviceorder

import xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"

const serviceOrderQueryUrl = "ServiceOrderQuery"

// Query 保单分页查询
func Query(params *xmhsdk.ServiceOrderQueryParam) (*xmhsdk.ServiceOrderQueryResult, error) {
	result := &xmhsdk.ServiceOrderQueryResult{}
	err := xmhsdk.MakeRequest(serviceOrderQueryUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("ServiceOrderQuery error: %v", err)
		return nil, err
	}
	return result, nil
}
