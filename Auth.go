package xmhOpenApiSdk

type AuthParam struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type Auth struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   uint64 `json:"expiresIn"`
}
