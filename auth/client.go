package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"io"
)

const authUrl = "applyToken"

func New() (*xmhsdk.Auth, error) {
	sdkClinet := xmhsdk.GetClient()
	params := &xmhsdk.AuthParam{
		AppId:     xmhsdk.AppId,
		AppSecret: xmhsdk.AppSecret,
	}
	jsonParam, err := json.Marshal(params)
	if err != nil {
		xmhsdk.Logger.Errorf("new auth json marshal error: %v", err)
		return nil, err
	}
	realUrl := fmt.Sprintf("%s/%s", sdkClinet.GetBaseUrl(), authUrl)
	xmhsdk.Logger.Debugf("new auth  request url: %s", realUrl)
	xmhsdk.Logger.Infof("new auth  request params: %s", string(jsonParam))
	resp, apiErr := sdkClinet.Post(realUrl, "application/json", bytes.NewBufferString(string(jsonParam)))
	if apiErr != nil {
		xmhsdk.Logger.Errorf("new auth  request error: %v", apiErr)
		return nil, apiErr
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		xmhsdk.Logger.Errorf("new auth  read response error: %v", err)
		return nil, err
	}
	xmhsdk.Logger.Debugf("new auth  response body: %s", string(bodyData))
	if resp.StatusCode != 200 {
		xmhsdk.Logger.Errorf("new auth  request error http status : %v", resp.Status)
		return nil, fmt.Errorf("request error http status : %v", resp.Status)
	}
	authOpenApiRes := &xmhsdk.OpenApiRes{
		Data: &xmhsdk.Auth{},
	}
	decodeErr := json.Unmarshal(bodyData, authOpenApiRes)
	if decodeErr != nil {
		xmhsdk.Logger.Errorf("new auth  decode response error: %v", decodeErr)
		return nil, decodeErr
	}
	if authOpenApiRes.ErrCode != 0 {
		xmhsdk.Logger.Errorf("new auth  request error: %v", authOpenApiRes.ErrMsg)
		return nil, fmt.Errorf("request error: %v", authOpenApiRes.ErrMsg)
	}
	if authOpenApiRes.Data == nil {
		xmhsdk.Logger.Errorf("new auth  request error: %v", "data is nil")
		return nil, fmt.Errorf("request error: %v", "data is nil")
	}
	auth, ok := authOpenApiRes.Data.(*xmhsdk.Auth)
	if !ok {
		xmhsdk.Logger.Errorf("new auth  got wrong response error: %v", "data is not *Auth")
		return nil, fmt.Errorf("got wrong response data type not auth")
	}
	if auth == nil {
		xmhsdk.Logger.Errorf("new auth  request error: %v", "data is nil")
		return nil, fmt.Errorf("request error: %v", "data is nil")
	}
	storage := xmhsdk.KvStorage
	err = storage.Save(xmhsdk.XMHTokenStorageKey, auth.AccessToken)
	if err != nil {
		xmhsdk.Logger.Errorf("new auth  save token error: %v", err)
		return nil, err
	}
	return auth, nil
}

func GetToken() (string, error) {
	storage := xmhsdk.KvStorage
	auth, err := storage.Get(xmhsdk.XMHTokenStorageKey)
	if err != nil {
		xmhsdk.Logger.Errorf("get token error: %v", err)
		return "", err
	}
	if auth == nil {
		xmhsdk.Logger.Errorf("get token error: %v", "auth is nil")
		return "", fmt.Errorf("auth is nil")
	}
	if token, ok := auth.(string); ok {
		return token, nil
	} else {
		xmhsdk.Logger.Errorf("XMHTokenStorageKey:%s exsist but got %v", xmhsdk.XMHTokenStorageKey, auth)
		return "", fmt.Errorf("XMHTokenStorageKey:%s exsist but got %v", xmhsdk.XMHTokenStorageKey, auth)
	}
}
