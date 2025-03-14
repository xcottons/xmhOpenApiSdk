package claim

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestClaimItemsQuery(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	params := &xmhsdk.ClaimItemsQueryParam{
		OrderID:            "1741948760",
		ClaimInsuranceType: 4,
	}
	result, err := ClaimItemsQuery(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimItemsQuery error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimItemsQuery result: %s", xmhsdk.ToStr(result))
}

func TestClaimReport(t *testing.T) {
	params := &xmhsdk.OpenApiClaimReport{
		OrderID: "",
		//
	}
	result, err := ClaimReport(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ClaimReport error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ClaimReport result: %v", result)
}
