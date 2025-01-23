package xmhOpenApiSdk

import (
	"net/http"
	"sync"
)

type EnvMark int

const (
	EnvIdc  = EnvMark(200)
	EnvBeta = EnvMark(600)
)

var (
	AppId     string
	AppSecret string

	ApiVersion    string  = "20250122"
	ClientVersion string  = "20250122"
	AlphaBaseUrl  string  = "https://api.stripe.com"
	BetaBaseUrl   string  = "https://api.stripe.com"
	IdcBaseUrl    string  = "https://api.stripe.com"
	CurrentEnv    EnvMark = EnvBeta
)

// SetEnvIdc set current env to idc
func SetEnvIdc() {
	SetEnv(EnvIdc)
}

// SetEnvBeta set current env to beta
func SetEnvBeta() {
	SetEnv(EnvBeta)
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
	default:
		return BetaBaseUrl

	}
}

type SdkClient struct {
	httpClient    *http.Client
	leveledLogger LeveledLoggerInterface
	baseUrl       string
}

var ClientOnce = sync.Once{}

func GetClient() *SdkClient {
	var client *SdkClient
	ClientOnce.Do(func() {
		client = &SdkClient{
			httpClient:    &http.Client{},
			leveledLogger: DefaultLeveledLogger,
			baseUrl:       GetBaseurl(),
		}
	})
	return client
}
