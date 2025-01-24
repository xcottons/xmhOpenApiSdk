package xmhOpenApiSdk

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
func (d *DItem) AddPp() *DItem {
	return d
}

type ProductItemsParam struct {
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
}
type OpenApiPlatformProductItemVariant struct {
	VariantId string `json:"variantId"`
	SkuId     string `json:"skuId"`
	Name      string `json:"Name"`
	Price     string `json:"price"`
	ImageUrl  string `json:"imageUrl,omitempty"`
}
