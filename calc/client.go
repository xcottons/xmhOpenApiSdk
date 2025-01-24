package calc

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const calcPriceURL = "CalcPrice"

func New(params *xmhsdk.CalcParams) (*xmhsdk.Calc, error) {
	result := &xmhsdk.Calc{}
	err := xmhsdk.MakeRequest(calcPriceURL, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("CalcPrice error: %v", err)
		return nil, err
	}
	return result, nil
}
