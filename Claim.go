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

// ClaimQueryByOrderParam 按订单查询理赔入参
type ClaimQueryByOrderParam struct {
	OrderID            string `json:"orderId"`            // 电商平台订单号，与 FrontOrderID 二选一
	FrontOrderID       string `json:"frontOrderId"`       // 电商平台前端订单号，与 OrderID 二选一
	ClaimInsuranceType string `json:"claimInsuranceType"` // 3-邮包险 4-理赔险
	Region             string `json:"region,omitempty"`   // ISO 3166-1 ALPHA2，默认 CN
}

// ClaimQueryByOrderResult 按订单查询理赔返回数据
type ClaimQueryByOrderResult struct {
	OrderID      string       `json:"orderId"`      // 长订单号
	FrontOrderID string       `json:"frontOrderId"` // 短订单号
	Claims       []*ClaimInfo `json:"claims"`       // 理赔单数组，无理赔时为空数组
}

// ClaimInfo 理赔单信息
type ClaimInfo struct {
	ClaimType              int32   `json:"claimType"`              // 索赔类型
	ClaimTypeDesc          string  `json:"claimTypeDesc"`          // 索赔类型描述（中文）
	ClaimState             int32   `json:"claimState"`             // 理赔状态
	ClaimStateDesc         string  `json:"claimStateDesc"`         // 理赔状态描述（中文）
	CompensationMethod     int32   `json:"compensationMethod"`     // 赔付方式
	CompensationMethodDesc string  `json:"compensationMethodDesc"` // 赔付方式描述（中文）
	ClaimApplyTime         string  `json:"claimApplyTime"`         // 理赔申请时间 RFC3339
	ClaimFinishTime        string  `json:"claimFinishTime"`        // 理赔完成时间 RFC3339
	ClaimMoney             float64 `json:"claimMoney"`             // 赔付金额（元）
	ClaimMoneyCurrency     string  `json:"claimMoneyCurrency"`     // 赔付金额币种
	Comments               string  `json:"comments"`               // 备注
}
