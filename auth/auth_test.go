package auth

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"testing"
)

func TestNewAuth(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth, err := New()
	if err != nil {
		t.Errorf("NewAuth error: %v", err)
	}
	t.Logf("NewAuth success: %v", auth)
	token, err := GetToken()
	if err != nil {
		t.Errorf("GetToken error: %v", err)
	} else {
		t.Logf("GetToken success: %v", token)
	}
}
