package xmhOpenApiSdk

type ShipParam struct {
	OrderId      string      `json:"orderId"`      // 订单ID，电商平台的订单号
	SubOrderId   string      `json:"subOrderId"`   // 子订单ID
	ShipInfoList []*ShipInfo `json:"shipInfoList"` // 货运信息列表
}

type ShipInfo struct {
	ShipId             string       `json:"shipId"`                    // 货运ID，若没有发货的实体ID，可以直接使用货运的运单号
	Currency           string       `json:"currency,omitempty"`        // 价格币种（可选）
	ShipPrice          string       `json:"shipPrice,omitempty"`       // 货运价格（可选）
	ActualShipSendTime string       `json:"actualShipSendTime"`        // 货运-实际发货时间
	ShipCompany        string       `json:"shipCompany"`               // 货运公司
	ShipCompanyCode    string       `json:"shipCompanyCode"`           // 货运公司代码(简称)
	ShipTrackNumber    string       `json:"shipTrackNumber"`           // 货运-运单号
	ShipStateString    string       `json:"shipStateString,omitempty"` // 货运状态（可选）
	ShipOtherInfo      *ShipAddress `json:"shipOtherInfo,omitempty"`   // 货运其他信息（可选）
	ItemList           []*DItem     `json:"itemList,omitempty"`        // 商品信息（可选）
}

type ShipResult struct {
	OrderId               string   `json:"orderId"`               // 电商平台的订单ID
	ShipId                []string `json:"shipId"`                // 电商平台的发货ID
	DisXmhShopOrderId     string   `json:"disXmhShopOrderId"`     // XMH侧的平台订单ID
	DisXmhShopOrderShipId []string `json:"disXmhShopOrderShipId"` // XMH侧的shipId
}
