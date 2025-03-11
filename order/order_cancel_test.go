package order

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"testing"
)

func TestTestCancelSpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{ServiceOrderId: "20250310XSP11675436AC758"},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))
}

func TestCancelPpServiceOrder(t *testing.T) {
	xmhsdk.AppId = "1000161"
	xmhsdk.AppSecret = "1WWyTSmQwCLWSFRt1WIEWbAmutfe4nbc"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	param := &xmhsdk.CancelServiceOrderParam{
		CancelServiceOrderItems: []xmhsdk.CancelServiceOrderItem{
			{ServiceOrderId: "20250310XPP11675436AC958"},
		},
	}
	result, err := CancelServiceOrder(param)
	if err != nil {
		t.Errorf("TestCancelSpServiceOrder failed,err:%v", err)
	}
	t.Logf("TestCancelSpServiceOrder result:%s", xmhsdk.ToStr(result))
}
