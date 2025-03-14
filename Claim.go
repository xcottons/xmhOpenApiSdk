package xmhOpenApiSdk

type OpenApiClaimReport struct {
	OrderID            string              `json:"orderId" validate:"required"`
	ClaimItems         []*OpenApiClaimItem `json:"claimItems" validate:"required"`
	ClaimType          int                 `json:"claimType" validate:"required"`
	FileLinks          []string            `json:"fileLinks" validate:"required"`
	Comments           string              `json:"comments" validate:"lte=1000"`
	ClaimPaymentObj    int                 `json:"claimPaymentObj" validate:"required"`
	PaymentMethod      int                 `json:"paymentMethod" validate:"required"`
	AccountInfo        AccountInfo         `json:"accountInfo" validate:"required"`
	ClaimInsuranceType int32               `json:"claimInsuranceType"`
}

type AccountInfo struct {
	AccountName   string `json:"accountName" validate:"required,gte=1,lte=30"`
	AccountNumber string `json:"accountNumber" validate:"required,gte=1,lte=20,checkInteger"`
	RoutingNumber string `json:"routingNumber" validate:"required,gte=1,lte=20,checkInteger"`
	Address       string `json:"address" validate:"required,gte=1,lte=100"`
}

type OpenApiClaimReportResult struct {
	ClaimIds []string `json:"claimIds"`
}

type ClaimItemsQueryParam struct {
	OrderID            string `json:"orderId" validate:"required"`
	ClaimInsuranceType int32  `json:"claimInsuranceType"`
}
type ClaimItems struct {
	Items []*OpenApiClaimItem `json:"claimItems"`
}

type OpenApiClaimItem struct {
	ItemId           string `json:"itemId"`
	SkuId            string `json:"skuId"`
	ItemNum          int32  `json:"itemNum"`
	MaxItemNum       int32  `json:"maxItemNum"`
	Pic              string `json:"pic"`
	GoodsName        string `json:"goodsName"`
	ItemUnitPrice    string `json:"itemUnitPrice"`
	ItemSumPrice     string `json:"itemSumPrice"`
	PriceCurrency    string `json:"priceCurrency"`
	ClaimApplyMoney  string `json:"claimApplyMoney"`
	ClaimReviewMoney string `json:"claimReviewMoney"`
	ClaimType        int32  `json:"claimType"`
	ClaimState       int32  `json:"claimState"`
	XmhServiceId     string `json:"xmhServiceId"`
}
