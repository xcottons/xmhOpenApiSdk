package calc

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const calcPriceUrl = "CalcPrice"
const ppCalcPriceUrl = "pPCalcPrice"

func Calc(params *xmhsdk.CalcParams) (*xmhsdk.Calc, error) {
	result := &xmhsdk.Calc{}
	err := xmhsdk.MakeRequest(calcPriceUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("CalcPrice error: %v", err)
		return nil, err
	}
	return result, nil
}
func PpCalc(params *xmhsdk.PpCalcParams) (*xmhsdk.PpCalc, error) {
	result := &xmhsdk.PpCalc{}
	err := xmhsdk.MakeRequest(ppCalcPriceUrl, params, result)
	if err != nil {
		xmhsdk.Logger.Errorf("CalcPrice error: %v", err)
		return nil, err
	}
	return result, nil
}
