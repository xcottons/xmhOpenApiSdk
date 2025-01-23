package xmhOpenApiSdk

type AuthParam struct {
	AppId     uint64 `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type Auth struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   uint64 `json:"expiresIn"`
}
