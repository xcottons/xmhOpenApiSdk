package xmhOpenApiSdk

type OrderState string

const (
	// OERDER_STATE_UNPAID indicates that the order is unpaid, and the customer has not completed the payment process.
	OERDER_STATE_UNPAID OrderState = "UNPAIED"

	// OERDER_STATE_PAID indicates that the order is paid, and the customer has successfully completed the payment.
	OERDER_STATE_PAID OrderState = "PAIED"

	// OERDER_STATE_UNSHIPPED indicates that the order is not shipped yet, but the payment has been completed.
	OERDER_STATE_UNSHIPPED OrderState = "UNSHIPED"

	// OERDER_STATE_SHIPPED indicates that the order has been shipped but not yet delivered to the customer.
	OERDER_STATE_SHIPPED OrderState = "SHIPED"

	// OERDER_STATE_DELIVERED indicates that the order has been delivered to the customer successfully.
	OERDER_STATE_DELIVERED OrderState = "DELIVERED"

	// OERDER_STATE_FINISHED indicates that the order is completed, and the transaction is closed.
	OERDER_STATE_FINISHED OrderState = "FINISHED"
)

type DOrder struct {
	OrderId           string         `json:"orderId"`
	SubOrderId        string         `json:"subOrderId,omitempty"`
	Currency          string         `json:"currency" validate:"required,currency"`
	TotalPrice        string         `json:"totalPrice" validate:"required,numeric"`
	PreferentialPrice string         `json:"preferentialPrice" validate:"required,numeric"`
	TaxPrice          string         `json:"taxPrice" validate:"omitempty,numeric"`
	ShipPrice         string         `json:"shipPrice" validate:"omitempty,numeric"`
	TotalPayPrice     string         `json:"totalPayPrice" validate:"omitempty,numeric"`
	InsuredPayPrice   string         `json:"insuredPayPrice,omitempty" validate:"omitempty,numeric"`
	PaySn             string         `json:"paySn"`
	PayTime           string         `json:"payTime" validate:"omitempty"`
	ItemList          []*DItem       `json:"itemList"`
	OrderState        OrderState     `json:"orderState"`
	SenderInfo        *SenderInfoDto `json:"senderInfo"`
	*ReceiverInfoDto  `json:"receiverInfo"`
	OrderCreateTime   string `json:"orderCreateTime" validate:"omitempty"`
	OrderModifyTime   string `json:"orderModifyTime" validate:"omitempty"`
	IsHistoryOrder    bool   `json:"isHistoryOrder"`
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
	OrderId        string    `json:"orderId"`        //your orderId
	XmhShopOrderId string    `json:"xmhShopOrderId"` //xmh saved your orderId
	Insurance      Insurance `json:"insurance"`
}

type PlatformOrderParam struct {
	UserId            string      `json:"userId" validate:"required"`
	UserEmail         string      `json:"userEmail" validate:"required,email"`
	OrderInfo         *DOrder     `json:"orderInfo" validate:"required"`
	ShipInfoList      []*ShipInfo `json:"shipInfoList"`
	BuyerIp           string      `json:"buyerIp"`
	DisRcId           string      `json:"disRcId"`
	DisComputeId      string      `json:"disComputeId"`
	ShopId            string      `json:"shopId"`
	DisXmhShopOrderId string      `json:"disXmhShopOrderId"`
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

type Insurance struct {
	SPInsureDetail SPInsureDetail `json:"spInsureDetail"`
	PPInsureDetail PPInsureDetail `json:"ppInsureDetail"`
}

type SPInsureDetail struct {
	ServiceOrderID string `json:"serviceOrderId"`
}

type PPInsureDetail struct {
	InsureItemList []*DItem `json:"insureItemList"`
}
