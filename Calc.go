package xmhOpenApiSdk

type CalcParams struct {
	UserID                   string  `json:"userId"`
	UserEmail                string  `json:"userEmail"`
	OrderInfo                *DOrder `json:"orderInfo"`
	BuyerIP                  string  `json:"buyerIp"`
	CartToken                string  `json:"cartToken"`
	IsShippingProtectionOpen bool    `json:"isShippingProtectionOpen"`
}

type Calc struct {
	RcResult        RcResult        `json:"rcResult"`
	SpComputeResult SpComputeResult `json:"spComputeResult"`
	PpComputeResult PpComputeResult `json:"ppComputeResult"`
}

type RcResult struct {
	DisRcId  string      `json:"disRcId"`
	RcState  int         `json:"rcState"`
	RcAction string      `json:"rcAction"`
	RcMsg    string      `json:"rcMsg"`
	RcInfo   interface{} `json:"rcInfo"`
}

type SpComputeResult struct {
	DisComputeId   string `json:"disComputeId"`
	ComputeState   int    `json:"computeState"`
	ComputeMsg     string `json:"computeMsg"`
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currencySymbol"`
	TotalPrice     string `json:"totalPrice"`
	ExtId          string `json:"extId"`
}

type PpComputeResult struct {
	DisComputeId string     `json:"disComputeId"`
	ComputeState int        `json:"computeState"`
	ComputeMsg   string     `json:"computeMsg"`
	ExtId        string     `json:"extId"`
	OVariants    []OVariant `json:"oVariants"`
}

type OVariant struct {
	OItemId        string          `json:"oItemId"`
	OVariantId     string          `json:"oVariantId"`
	OSku           string          `json:"oSku"`
	PpVariantItems []PpVariantItem `json:"ppVariantItems"`
}

type PpVariantItem struct {
	VariantId      string            `json:"variantId"`
	Name           string            `json:"name"`
	PriceString    string            `json:"priceString"`
	Currency       string            `json:"currency"`
	CurrencySymbol string            `json:"currencySymbol"`
	Properties     VariantProperties `json:"properties"`
}

type VariantProperties struct {
	PlanID    string `json:"Plan ID"`
	Product   string `json:"Product"`
	Reference string `json:"Reference"`
}
