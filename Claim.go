package xmhOpenApiSdk

type OpenApiClaimReport struct {
	ServiceOrderId     string              `json:"serviceOrderId" validate:"required"`
	ClaimItems         []*OpenApiClaimItem `json:"claimItems" validate:"required"`
	ClaimType          int                 `json:"claimType" validate:"required"`
	FileLinks          []string            `json:"fileLinks" validate:"required"`
	Comments           string              `json:"comments" validate:"lte=1000"`
	ClaimPaymentObj    int                 `json:"claimPaymentObj" validate:"required"`
	PaymentMethod      int                 `json:"paymentMethod" validate:"required"`
	AccountInfo        AccountInfo         `json:"accountInfo" validate:"required"`
	ClaimReportTime    string              `json:"claimReportTime"`
	LossOccurrenceTime string              `json:"lossOccurrenceTime"`
	Describe           string              `json:"describe"`
}

type AccountInfo struct {
	AccountName   string `json:"accountName" validate:"required,gte=1,lte=30"`
	AccountNumber string `json:"accountNumber" validate:"required,gte=1,lte=20,checkInteger"`
	RoutingNumber string `json:"routingNumber" validate:"required,gte=1,lte=20,checkInteger"`
	Address       string `json:"address" validate:"required,gte=1,lte=100"`
}

type OpenApiClaimReportResult struct {
	ClaimId  string   `json:"claimId"`
	ClaimIds []string `json:"claimIds"`
}

type ClaimItemsQueryParam struct {
	OrderID            string `json:"orderId" validate:"required"`
	ClaimInsuranceType int32  `json:"claimInsuranceType"`
}
type ServiceClaimItemsQueryParam struct {
	ServiceOrderId string `json:"serviceOrderId"`
}
type ClaimItems []*OpenApiClaimItem
type ServiceClaimItemsQueryResult struct {
	ClaimItems ClaimItems `json:"claimItems"`
}

type OpenApiClaimItem struct {
	ProductId        string `json:"productId"`
	ItemId           string `json:"itemId"`
	SkuId            string `json:"skuId"`
	OrderGoodsId     string `json:"orderGoodsId"`
	PlanId           string `json:"planId"` //延保方案id
	VariantId        string `json:"variantId"`
	ItemNum          int32  `json:"itemNum"`     // 理赔申请的数量
	MaxClaimNum      int32  `json:"maxClaimNum"` // 理赔商品最大可申请数量
	Pic              string `json:"pic"`
	ItemName         string `json:"itemName"`
	ItemUnitPrice    string `json:"itemUnitPrice"`
	ItemSumPrice     string `json:"itemSumPrice"`
	PriceCurrency    string `json:"priceCurrency"`
	ClaimApplyMoney  string `json:"claimApplyMoney"`
	ClaimReviewMoney string `json:"claimReviewMoney"`
	ClaimType        int32  `json:"claimType"`
	ClaimState       int32  `json:"claimState"`
	XmhServiceId     string `json:"xmhServiceId"`
	Claimpayout      string `json:"claimpayout"`
}

type ClaimQueryParam struct {
	ServiceOrderId string `json:"serviceOrderId"`
	ClaimId        string `json:"claimId"`
}
