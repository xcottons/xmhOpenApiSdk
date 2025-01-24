package xmhOpenApiSdk

type OrderState string

const (
	// OERDER_STATE_UNPAID indicates that the order is unpaid, and the customer has not completed the payment process.
	OERDER_STATE_UNPAID OrderState = "UNPAID"

	// OERDER_STATE_PAID indicates that the order is paid, and the customer has successfully completed the payment.
	OERDER_STATE_PAID OrderState = "PAID"

	// OERDER_STATE_UNSHIPPED indicates that the order is not shipped yet, but the payment has been completed.
	OERDER_STATE_UNSHIPPED OrderState = "UNSHIPPED"

	// OERDER_STATE_SHIPPED indicates that the order has been shipped but not yet delivered to the customer.
	OERDER_STATE_SHIPPED OrderState = "SHIPPED"

	// OERDER_STATE_DELIVERED indicates that the order has been delivered to the customer successfully.
	OERDER_STATE_DELIVERED OrderState = "DELIVERED"

	// OERDER_STATE_FINISHED indicates that the order is completed, and the transaction is closed.
	OERDER_STATE_FINISHED OrderState = "FINISHED"
)

type DOrder struct {
	OrderId           string         `json:"orderId"`
	SubOrderId        string         `json:"subOrderId,omitempty"`
	Currency          string         `json:"currency" validate:"required,currency"`
	TotalPrice        string         `json:"totalPrice" validate:"required,numeric"`                 //总金额，总金额(含税费，含运费) - 优惠金额 = 实际支付金额，以元为单位
	PreferentialPrice string         `json:"preferentialPrice" validate:"required,numeric"`          //优惠金额，以元为单位，订单的优惠金额。
	TaxPrice          string         `json:"taxPrice" validate:"omitempty,numeric"`                  //税费，可能是增值税可能是其他税，以元为单位
	ShipPrice         string         `json:"shipPrice" validate:"omitempty,numeric"`                 //运费，以元为单位
	TotalPayPrice     string         `json:"totalPayPrice" validate:"omitempty,numeric"`             //实际支付金额，以元为单位
	InsuredPayPrice   string         `json:"insuredPayPrice,omitempty" validate:"omitempty,numeric"` //邮包险支付金额，以元为单位 （邮包险）
	PayTime           string         `json:"payTime" validate:"omitempty"`                           //支付时间
	ItemList          []*DItem       `json:"itemList"`
	OrderState        OrderState     `json:"orderState"`                           //同步平台订单时需要判断是否为空
	SenderInfo        *SenderInfoDto `json:"senderInfo"`                           //validate:"required" 试算的时候不要求填写，同步订单数据的时候才要求。
	*ReceiverInfoDto  `json:"receiverInfo"`                                        //validate:"required" 同上。
	OrderCreateTime   string         `json:"orderCreateTime" validate:"omitempty"` //同步平台订单时需要判断是否为空
	OrderModifyTime   string         `json:"orderModifyTime" validate:"omitempty"` //同步平台订单时需要判断是否为空
	IsHistoryOrder    bool           `json:"isHistoryOrder"`                       //是否历史订单，默认为false
}

type SenderInfoDto struct {
	SenderShipAddress *ShipAddress           `json:"senderShipAddress"`
	ExtendInfo        map[string]interface{} `json:"extendInfo"`
}

type ReceiverInfoDto struct {
	StoreName           string                 `json:"storeName"`
	StoreCode           string                 `json:"storeCode"`
	ReceiverShipAddress *ShipAddress           `json:"receiverShipAddress" validate:"required"`
	ExtendInfo          map[string]interface{} `json:"extendInfo"`
}

type ShipAddress struct {
	Country      string `json:"country"`
	CountryCode  string `json:"countryCode"`
	Province     string `json:"province" `
	ProvinceCode string `json:"provinceCode"`
	City         string `json:"city" `
	CityCode     string `json:"cityCode"`
	District     string `json:"district"`
	Street       string `json:"street"`
	Address      string `json:"address"`
	PostalCode   string `json:"postalCode"`
	Name         string `json:"name"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CompanyName  string `json:"companyName"`
	Mobile       string `json:"mobile"`
	Cert         string `json:"cert"`
	Email        string `json:"email" validate:"omitempty,email"`
	LinkName     string `json:"linkName"`
	LinkMobile   string `json:"linkMobile"`
}

//to add ship protection for this order  you should call this method and give the orderTotalPayFee when you sync order to xmh
func (d *DOrder) AddSp(orderTotalPayFee string) *DOrder {
	d.InsuredPayPrice = orderTotalPayFee
	return d
}

type PlatformOrderResult struct {
	OrderId           string   `json:"orderId"`           //your orderId
	DisXmhShopOrderId string   `json:"disXmhShopOrderId"` //xmh saved your orderId
	DisXmhLDealIds    []string `json:"disXmhLDealIds"`    //xmh service deal id list,  one of your order maybe gen multiple deal
}

type PlatformOrderParam struct {
	UserId            string  `json:"userId" validate:"required"`
	UserEmail         string  `json:"userEmail" validate:"required,email"`
	OrderInfo         *DOrder `json:"orderInfo" validate:"required"`
	BuyerIp           string  `json:"buyerIp"`
	DisRcId           string  `json:"disRcId"`
	DisComputeId      string  `json:"disComputeId"`
	ShopId            string  `json:"shopId"`
	DisXmhShopOrderId string  `json:"disXmhShopOrderId"`
}

type CancelOrderParam struct {
	OrderId           string   `json:"orderId"`
	SubOrderId        string   `json:"subOrderId,omitempty"`
	DisXmhShopOrderId string   `json:"disXmhShopOrderId,omitempty"`
	CancelId          string   `json:"cancelId"`
	CancelReasonType  string   `json:"cancelReasonType"`
	CancelReason      string   `json:"cancelReason"`
	Currency          string   `json:"currency,omitempty"`
	TotalRefundPrice  string   `json:"totalRefundPrice,omitempty"`
	CancelItems       []*DItem `json:"cancelItems,omitempty"`
}
type InsuredOrderParam struct {
	UserId            string  `json:"userId" validate:"required"`
	UserEmail         string  `json:"userEmail" validate:"required,email"`
	OrderInfo         *DOrder `json:"orderInfo" validate:"required"`
	BuyerIp           string  `json:"buyerIp"`
	DisRcId           string  `json:"disRcId"`
	DisComputeId      string  `json:"disComputeId"`
	ShopId            string  `json:"shopId"`
	DisXmhShopOrderId string  `json:"disXmhShopOrderId"`
}

type InsuredOrderResult struct {
	DisXmhShopOrderId string `json:"disXmhShopOrderId"`
}

type CancelOrderResult struct {
	DisXmhShopOrderId string `json:"disXmhShopOrderId"`
}
