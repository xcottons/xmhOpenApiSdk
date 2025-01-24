package auth

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"testing"
)

func TestNewAuth(t *testing.T) {
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	param := &xmhsdk.AuthParam{
		AppId:     "1000151",
		AppSecret: "3j0k9TkrkiPbvfl2eLjqfNHUBaTOvR1p",
	}
	auth, err := New(param)
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
