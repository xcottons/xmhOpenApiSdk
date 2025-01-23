package xmhOpenApiSdk

type DOrderDto struct {
	OrderId                    string                 `json:"orderId"`
	SubOrderId                 string                 `json:"subOrderId,omitempty"`
	SubOrderId2                string                 `json:"subOrderId2,omitempty"`
	FrontId1                   string                 `json:"frontId1,omitempty"`
	Currency                   string                 `json:"currency" validate:"required,currency"`
	TotalPrice                 string                 `json:"totalPrice" validate:"required,numeric"`                  //总金额，总金额(含税费，含运费) - 优惠金额 = 实际支付金额，以元为单位
	PreferentialPrice          string                 `json:"preferentialPrice" validate:"required,numeric"`           //优惠金额，以元为单位，订单的优惠金额。
	TaxPrice                   string                 `json:"taxPrice" validate:"omitempty,numeric"`                   //税费，可能是增值税可能是其他税，以元为单位
	ShipPrice                  string                 `json:"shipPrice" validate:"omitempty,numeric"`                  //运费，以元为单位
	TotalPayPrice              string                 `json:"totalPayPrice" validate:"omitempty,numeric"`              //实际支付金额，以元为单位
	InsuredPayPrice            string                 `json:"insuredPayPrice,omitempty" validate:"omitempty,numeric"`  //邮包险支付金额，以元为单位 （邮包险）
	InsuredPayPriceForExtended string                 `json:"insuredPayPriceForExtended" validate:"omitempty,numeric"` //延保险支付金额，以元为单位  （延保险）
	PayTime                    string                 `json:"payTime" validate:"omitempty,datatimefrfc3339"`           //支付时间
	OtherPriceInfo             map[string]interface{} `json:"otherPriceInfo"`                                          //其他价格信息
	OrderOtherInfo             map[string]interface{} `json:"orderOtherInfo"`                                          //其他订单信息
	SubOrderList               []*DOrderDto           `json:"subOrderList"`
	ItemList                   []*DItem               `json:"itemList"`
	ItemValidNum               int                    `json:"itemValidNum"`                                          //有效商品数量，除了邮包险以外的商品数量。无需发货的商品，也算在内。在shopify中有用到。
	OrderState                 string                 `json:"orderState"`                                            //同步平台订单时需要判断是否为空
	SenderInfo                 *SenderInfoDto         `json:"senderInfo"`                                            //validate:"required" 试算的时候不要求填写，同步订单数据的时候才要求。
	ReceiverInfo               *ReceiverInfoDto       `json:"receiverInfo"`                                          //validate:"required" 同上。
	OrderCreateTime            string                 `json:"orderCreateTime" validate:"omitempty,datatimefrfc3339"` //同步平台订单时需要判断是否为空
	OrderModifyTime            string                 `json:"orderModifyTime" validate:"omitempty,datatimefrfc3339"` //同步平台订单时需要判断是否为空
	IsHistoryOrder             bool                   `json:"isHistoryOrder"`                                        //是否历史订单，默认为false
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
	Country      string `json:"country" // validate:"required"`
	CountryCode  string `json:"countryCode"`
	Province     string `json:"province" // validate:"required"`
	ProvinceCode string `json:"provinceCode"`
	City         string `json:"city" // validate:"required"`
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
