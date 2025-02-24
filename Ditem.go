package xmhOpenApiSdk

import "errors"

var (
	PPPlanCodeFor1Years = "1 Year"
	PPPlanCodeFor2Years = "2 Years"
	PPPlanCodeFor3Years = "3 Years"
)

type PpPlanYear int

const (
	PpPlanOneYear = iota + 1
	PpPlanTwoYear
	PpPlanThreeYear
)

var ppPlanYearMap = map[PpPlanYear]string{
	PpPlanOneYear:   PPPlanCodeFor1Years,
	PpPlanTwoYear:   PPPlanCodeFor2Years,
	PpPlanThreeYear: PPPlanCodeFor3Years,
}

var ppYearsMap = map[PpPlanYear]string{
	PpPlanOneYear:   "1 Year",
	PpPlanTwoYear:   "2 Year",
	PpPlanThreeYear: "3 Year",
}

type DItem struct {
	ItemId               string          `json:"itemId" validate:"required"`
	VariantId            string          `json:"variantId"`
	SkuId                string          `json:"skuId"`
	ProductId            string          `json:"productId"`
	ItemName             string          `json:"itemName" validate:"required"`
	ItemDesc             string          `json:"itemDesc"`
	ItemState            int32           `json:"itemState"`
	ItemTag              string          `json:"itemTag"`
	ItemType             string          `json:"itemType"`
	ItemOtherInfoJson    string          `json:"itemOtherInfoJson"`
	ItemPicInfoMain      []string        `json:"ItemPicInfoMain"`
	ItemPicInfoOther     []string        `json:"ItemPicInfoOther"`
	Currency             string          `json:"currency" validate:"required,currency"`
	UnitNum              string          `json:"unitNum" validate:"required,numeric"`
	IsPackageItem        bool            `json:"isPackageItem"`
	IsPPItem             bool            `json:"isPPItem"`
	UnitPrice            string          `json:"unitPrice" validate:"required,numeric"`
	UnitPriceInt         int64           `json:"unitPriceInt"`
	PreferentialPrice    string          `json:"preferentialPrice"`
	PreferentialPriceInt int64           `json:"preferentialPriceInt"`
	SubTotalPrice        string          `json:"subTotalPrice" validate:"omitempty,numeric"`
	SubTotalPriceInt     int64           `json:"subTotalPriceInt"`
	SubTotalTaxPrice     string          `json:"subTotalTaxPrice" validate:"omitempty,numeric"`
	TaxPrice             string          `json:"taxPrice" validate:"omitempty,numeric"`
	TotalPrice           string          `json:"totalPrice" validate:"required,numeric"`
	TotalPriceInt        int64           `json:"totalPriceInt"`
	TotalPayPrice        string          `json:"totalPayPrice"`
	TotalPayPriceInt     int64           `json:"totalPayPriceInt"`
	InsuredPayPrice      string          `json:"insuredPayPrice,omitempty" validate:"omitempty,numeric"`
	Properties           DItemProperties `json:"Properties"`
	PpVariant            *PPVariant      `json:"ppVariant"`
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

// to Add a product protection to this order item you should add a ppVariant object to the item. which come from calc api when you sync order to xmh
func (d *DItem) AddPp(pPrice string, py PpPlanYear) (*DItem, error) {
	if pId, ok := ppPlanYearMap[py]; !ok {
		return nil, errors.New("invalid pp plan year")
	} else {
		d.Properties = DItemProperties{
			"Plan ID": pId,
		}
		d.InsuredPayPrice = pPrice
		d.PpVariant = &PPVariant{
			VariantID:      "",
			Name:           ppYearsMap[py],
			PriceString:    pPrice,
			Currency:       d.Currency,
			CurrencySymbol: "",
			Properties: map[string]string{
				"Reference": d.ItemId,
				"Plan ID":   pId,
				"Product":   d.ItemName,
			},
		}
	}
	return d, nil
}

func (d *DItem) AddOneYearPp(pPrice string) *DItem {
	pp, _ := d.AddPp(pPrice, PpPlanOneYear)
	return pp
}
func (d *DItem) AddTwoYearPp(pPrice string) *DItem {
	pp, _ := d.AddPp(pPrice, PpPlanTwoYear)
	return pp
}
func (d *DItem) AddThreeYearPp(pPrice string) *DItem {
	pp, _ := d.AddPp(pPrice, PpPlanThreeYear)
	return pp
}

type ProductItemsParam struct {
	BusiId  string         `json:"busiId"`
	BatchId uint64         `json:"batchId"`
	Items   []*ProductItem `json:"items"`
	FileUrl string         `json:"fileUrl"`
}

type ProductItem struct {
	PlatformItemId string                              `json:"platformItemId"`
	ItemState      int32                               `json:"itemState"`
	ItemName       string                              `json:"itemName"`
	Currency       string                              `json:"currency"`
	PicLink        []string                            `json:"picLink,omitempty"`
	ItemLink       string                              `json:"itemLink,omitempty"`
	Variants       []OpenApiPlatformProductItemVariant `json:"variants,omitempty"`
	ItemPriceExt   string                              `json:"itemPriceExt,omitempty"`
	ItemDetailExt  string                              `json:"itemDetailExt,omitempty"`
	VariantId      string                              `json:"variantId"`
	SkuId          string                              `json:"skuId"`
	Name           string                              `json:"Name"`
	Price          string                              `json:"price"`
}
type OpenApiPlatformProductItemVariant struct {
	VariantId string `json:"variantId"`
	SkuId     string `json:"skuId"`
	Name      string `json:"Name"`
	Price     string `json:"price"`
	ImageUrl  string `json:"imageUrl,omitempty"`
}
