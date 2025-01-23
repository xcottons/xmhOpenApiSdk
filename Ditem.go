package xmhOpenApiSdk

import "encoding/json"

type DItem struct {
	ItemId            string   `json:"itemId" validate:"required"` //商品ID
	LineId            string   `json:"lineId"`                     //商品行ID或者购物车行ID
	VariantId         string   `json:"variantId"`                  //变体ID
	SkuId             string   `json:"skuId"`                      //SKUID
	ProductId         string   `json:"productId"`                  //产品ID
	UnitId            string   `json:"unitId"`                     //单位ID，一般都是空。目前只有延保险2个商品的情况下，要拆分为2个DitemDto，会加上这个UnitId的标识。
	ItemName          string   `json:"itemName" validate:"required"`
	ItemDesc          string   `json:"itemDesc"` //商品描述
	ItemState         int32    `json:"itemState"`
	ItemTag           string   `json:"itemTag"`           //商品标签
	ItemType          string   `json:"itemType"`          //商品类别
	ItemOtherInfoJson string   `json:"itemOtherInfoJson"` //商品的其他信息，放JSON中。一般用不上。。。
	ItemPicInfoMain   []string `json:"ItemPicInfoMain"`
	ItemPicInfoOther  []string `json:"ItemPicInfoOther"`
	Currency          string   `json:"currency" validate:"required,currency"`
	UnitNum           string   `json:"unitNum" validate:"required,numeric"` //商品数量

	IsPackageItem bool `json:"isPackageItem"` //是否邮包险商品。默认为false
	IsPPItem      bool `json:"isPPItem"`      //是否延保险商品。默认为false

	UnitPrice    string `json:"unitPrice" validate:"required,numeric"` //单价
	UnitPriceInt int64  `json:"unitPriceInt"`                          //非前端输入参数，只是后端转存了一下以分为单位的价格。

	PreferentialPrice    string `json:"preferentialPrice"`    //优惠价格，商品的优惠金额。如果有些优惠金额落在商品上，有些落在订单上，但最终应该都可以将优惠金额折算到商品上面。
	PreferentialPriceInt int64  `json:"preferentialPriceInt"` //同 UnitPriceInt

	SubTotalPrice    string `json:"subTotalPrice" validate:"omitempty,numeric"` //小结的总价，不是最终总价。有可能是折扣之前的价格。
	SubTotalPriceInt int64  `json:"subTotalPriceInt"`                           //同 UnitPriceInt

	SubTotalTaxPrice string `json:"subTotalTaxPrice" validate:"omitempty,numeric"` //小结税费
	TaxPrice         string `json:"taxPrice" validate:"omitempty,numeric"`         //税费

	TotalPrice    string `json:"totalPrice" validate:"required,numeric"` //最终总价
	TotalPriceInt int64  `json:"totalPriceInt"`                          //同 UnitPriceInt

	TotalPayPrice    string `json:"totalPayPrice"`    //商品实际支付总价
	TotalPayPriceInt int64  `json:"totalPayPriceInt"` //同UnitPriceInt

	ItemWeightUnit  string `json:"itemWeightUnit"`
	ItemWeight      string `json:"itemWeight" validate:"omitempty,numeric"`
	ItemVolumeUnit  string `json:"itemVolumeUnit"`
	ItemVolume      string `json:"itemVolume" validate:"omitempty,numeric"`
	ItemCountry     string `json:"itemCountry"`
	ItemCountryCode string `json:"itemCountryCode"`

	AppliedClaimNum int32           `json:"appliedClaimNum"` // 已申请理赔的数量
	ClaimState      int32           `json:"claimState"`      // 理赔状态
	ClaimType       int32           `json:"claimType"`       // 理赔类型
	RequireShip     bool            `json:"requireShip"`     // 是否需要发货
	GiftCard        bool            `json:"giftCard"`        // 是否礼品卡
	Properties      DItemProperties `json:"Properties"`
	PpVariant       *PPVariant      `json:"ppVariant"`
	Ref             *DItem          `json:"ref"` // 引用的商品 非入参 中间变量
}
type PPVariant struct {
	VariantID      string          `json:"variantId,omitempty"` // 延保险商品的变体ID
	Name           string          `json:"name,omitempty"`      // 延保商品的名称，例如：1 Years
	PriceString    string          `json:"price"`               // 延保的价格
	Currency       string          `json:"currency"`            // 延保的币种
	CurrencySymbol string          `json:"currencySymbol"`      // 延保币种符号
	Properties     DItemProperties `json:"properties"`          // 保障商品ID、计划ID和商品名等属性
}

type DItemProperties map[string]string

func (dp DItemProperties) GetReference() (string, bool) {
	if v, ok := dp["Reference"]; ok {
		return v, true
	}
	return "", false
}
func (dp DItemProperties) GetPlanId() (string, bool) {
	if v, ok := dp["Plan ID"]; ok {
		return v, true
	}
	return "", false
}

func (dp DItemProperties) GetProduct() (string, bool) {
	if v, ok := dp["Product"]; ok {
		return v, true
	}
	return "", false
}
func (h *DItem) Clone() *DItem {
	bs, _ := json.Marshal(h)
	var cloned DItem
	_ = json.Unmarshal(bs, &cloned)

	return &cloned
}
