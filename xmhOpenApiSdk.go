package xmhOpenApiSdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

type EnvMark int

const (
	EnvIdc   = EnvMark(200)
	EnvBeta  = EnvMark(600)
	EnvAlpha = EnvMark(500)
	EnvLocal = EnvMark(100)
)

var (
	AlphaBaseUrl    string = "https://alphaxcottons.xinhulu.com"
	BetaBaseUrl     string = "https://betaxcottons.xinhulu.com"
	IdcBaseUrl      string = "https://seller.xcottons.com"
	LocalBaseUrl    string = ""
	SignSecretIdc   string = "ex_xmh_idc"
	SignSecretBeta  string = "ex_xmh_test"
	SignSecretAlpha string = "ex_xmh_test"
	SignSecretLocal string = "ex_xmh_test"
)
var (
	AppId     string
	AppSecret string

	ApiVersion         string                 = "20250122"
	ApiVersionDesc     string                 = "init std openapi"
	ClientVersion      string                 = "20250122"
	ClientVersionDesc  string                 = "init std openapi sdk"
	CurrentEnv         EnvMark                = EnvBeta
	BasePath           string                 = "go/packageOpenApiWeb/OpenApiV2"
	Logger             LeveledLoggerInterface = DefaultLeveledLogger
	KvStorage          Storage                = &DefaultMemoryStorage{}
	XMHTokenStorageKey string                 = "xmh_token_storage_key"
	SignSecret         string                 = SignSecretBeta
)

// SetEnvIdc set current env to idc
func SetEnvIdc() {
	SetEnv(EnvIdc)
	SignSecret = SignSecretIdc
	Logger.Infof("set env to idc")
}

// SetEnvBeta set current env to beta
func SetEnvBeta() {
	SetEnv(EnvBeta)
	SignSecret = SignSecretBeta
	Logger.Infof("set env to beta")
}
func SetEnv(env EnvMark) {
	CurrentEnv = env
	switch env {
	case EnvIdc:
		GetClient().baseUrl = IdcBaseUrl
	case EnvBeta:
		GetClient().baseUrl = BetaBaseUrl
	}
}
func GetBaseurl() string {
	switch CurrentEnv {
	case EnvIdc:
		return IdcBaseUrl
	case EnvBeta:
		return BetaBaseUrl
	case EnvAlpha:
		return AlphaBaseUrl
	default:
		return BetaBaseUrl

	}
}

func SetLogger(logger LeveledLoggerInterface) {
	Logger = logger
}

func SetKvStorage(storage Storage) {
	KvStorage = storage
}

func SetXmhTokenStorageKey(key string) {
	XMHTokenStorageKey = key
}

type SdkClient struct {
	http.Client
	baseUrl string
}

func (sc *SdkClient) GetBaseUrl() string {
	return fmt.Sprintf("%s/%s", sc.baseUrl, BasePath)
}

func (sc *SdkClient) Do(req *http.Request) (*http.Response, error) {
	if !strings.Contains(req.URL.Path, "applyToken") {
		get, err := KvStorage.Get(XMHTokenStorageKey)
		if err != nil {
			Logger.Errorf("get token error: %v", err)
			return nil, err
		}
		if get == nil {
			Logger.Errorf("get token error: %v", "auth is nil")
			return nil, fmt.Errorf("auth is nil")
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", get.(string)))
	}
	req.Header.Set("X-Client-Version", ClientVersion)
	req.Header.Set("X-Client-Version-Desc", ClientVersionDesc)
	req.Header.Set("X-Api-Version", ApiVersion)
	req.Header.Set("X-Api-Version-Desc", ApiVersionDesc)
	req.Header.Set("X-App-Id", AppId)
	if req.Method == http.MethodPost && req.Body != nil { //if post request, add signature
		hmacSha256 := hmac.New(sha256.New, []byte(SignSecret))
		bodyData, err := io.ReadAll(req.Body)
		if err != nil {
			Logger.Errorf("read request body error: %v", err)
			return nil, err
		}
		defer req.Body.Close()
		hmacSha256.Write(bodyData)
		signature := base64.StdEncoding.EncodeToString(hmacSha256.Sum(nil))
		Logger.Debugf("signature: %s", signature)
		req.Header.Set("Signature", signature)
		req.Body = io.NopCloser(bytes.NewReader(bodyData)) // reset body
	}

	Logger.Debugf("request url: %s", req.URL.String())
	return sc.Client.Do(req)
}
func (sc *SdkClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return sc.Do(req)
}

var ClientOnce = sync.Once{}
var client *SdkClient

func GetClient() *SdkClient {
	ClientOnce.Do(func() {
		client = &SdkClient{
			baseUrl: GetBaseurl(),
		}
	})
	return client
}
func MakeRequest(urlPath string, params interface{}, resultData interface{}) error {
	sdkClient := GetClient()
	jsonParam, err := json.Marshal(params)
	if err != nil {
		Logger.Errorf("%s json marshal error: %v", urlPath, err)
		return err
	}
	realUrl := fmt.Sprintf("%s/%s", sdkClient.GetBaseUrl(), urlPath)
	Logger.Debugf("%s request url: %s", urlPath, realUrl)
	Logger.Infof("%s request params: %s", urlPath, string(jsonParam))
	resp, apiErr := sdkClient.Post(realUrl, "application/json", bytes.NewBuffer(jsonParam))
	if apiErr != nil {
		Logger.Errorf("%s request error: %v", urlPath, apiErr)
		return apiErr
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Errorf("%s read response error: %v", urlPath, err)
		return err
	}
	Logger.Debugf("%s response body: %s", urlPath, string(bodyData))
	if resp.StatusCode != 200 {
		Logger.Errorf("%s request error http status: %v", urlPath, resp.Status)
		return fmt.Errorf("request error http status: %v", resp.Status)
	}
	apiRes := &OpenApiRes{
		Data: resultData,
	}
	err = json.Unmarshal(bodyData, apiRes)
	if err != nil {
		Logger.Errorf("%s decode response error: %v", urlPath, err)
		return err
	}
	if apiRes.ErrCode != 0 {
		Logger.Errorf("%s request error: %v", urlPath, apiRes.ErrMsg)
		return fmt.Errorf("request error: %v", apiRes.ErrMsg)
	}
	//if apiRes.Data == nil {
	//	Logger.Errorf("%s request error: data is nil", urlPath)
	//	return fmt.Errorf("request error: data is nil")
	//}
	return nil
}
