package ship

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
)

func TestShipSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvXmhAlpha)
	auth.New()
	params := &xmhsdk.ShipParam{
		OrderId:    "1770295473",
		SubOrderId: "",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:          strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode: "ups",
				ShipCompany:     "ups",
				//ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipTrackNumber: "YT2529200702482310",

				ShipStateString:    "已发货",
				ShipPrice:          "0.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: "2025-10-19T14:03:12+08:00",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:    "United States",
					CityCode:   "US",
					Province:   "New Jersey",
					PostalCode: "NJ",
					City:       "Union City",
					//CountryCode: "CN",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId: "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:  "SKU001",
					//OrderGoodsId:      "SKU001000001",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "0",
					UnitNum:           "1",
					TotalPrice:        "150",
					PreferentialPrice: "0",
					TotalPayPrice:     "150.00",
				}},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	item := &xmhsdk.DItem{
		ItemId: "16GA5QRB5TO01_110",
		//SkuId:             "SKU001",
		//OrderGoodsId:      "16GA5QRB5TO01_110",
		ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
		InsuredPayPrice:   "24.0",
	}
	item1 := &xmhsdk.DItem{
		ItemId: "16GA5QRB5TO01_110567",
		//SkuId:             "SKU001",
		OrderGoodsId:      "16GA5QRB5TO01_110",
		ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
		InsuredPayPrice:   "24.0",
	}
	intn := rand.New(rand.NewSource(time.Now().Unix()))

	params := &xmhsdk.ShipParam{
		OrderId:    "1772674982",
		SubOrderId: "",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: time.Date(2024, 2, 29, 11, 11, 11, 11, time.Local).Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					//City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{item},
			},
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: time.Date(2024, 2, 29, 11, 11, 11, 11, time.Local).Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					//City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{item1},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}
func TestShip(t *testing.T) {
	//xmhsdk.AppId = "1600205"
	//xmhsdk.AppSecret = "kalGfuKmjSOWkLz3RblBeNaACXT1tZMc"
	//xmhsdk.SignSecret = ""
	//xmhsdk.SetEnv(xmhsdk.EnvIdc)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:    "14031429304508",
		VariantId: "44700112978108",
		ProductId: "8765254533308",
		SkuId:     "",
		UnitNum:   "1",
	}

	params := &xmhsdk.ShipParam{
		OrderId:    "5817078907068",
		SubOrderId: "",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             "881474911169",
				ShipCompanyCode:    "Fedex",
				ShipCompany:        "Fedex",
				ShipTrackNumber:    "881474911169",
				ShipStateString:    "已发货",
				ShipPrice:          "0",
				ActualShipSendTime: "2025-05-22T00:00:00Z",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "United States",
					CityCode: "US",
					Province: "Texas",
					City:     "Houston",
				},
				ItemList: []*xmhsdk.DItem{p1},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipWithShipTime(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	params := &xmhsdk.ShipParam{
		OrderId:    "1741315736",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:          strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode: "SF",
				ShipTrackNumber: strconv.Itoa(int(time.Now().Unix())),
				ShipStateString: "已发货",
				ShipPrice:       "6.00",
				//ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:             "SKU001",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "100.00",
					UnitNum:           "2",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
					PpVariant: &xmhsdk.PPVariant{
						VariantID:   "",
						Name:        "3 Years",
						PriceString: "10.00",
						Currency:    "USD",
						Properties: map[string]string{
							"Reference": "97760109-16ad-40c9-9385-caba381a26aa",
							"Plan ID":   "10802001",
							"Product":   "Durable Concrete Shirt",
						},
					},
				}, {
					ItemId:            "666666",
					SkuId:             "SKU002",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "100.00",
					UnitNum:           "2",
					TotalPrice:        "200.00",
					PreferentialPrice: "50.00",
					TotalPayPrice:     "150.00",
				}, p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00"),
				//{
				//	ItemId:            "666666",
				//	SkuId:             "SKU002",
				//	ItemName:          "Durable Concrete Shirt",
				//	Currency:          "USD",
				//	UnitPrice:         "100.00",
				//	UnitNum:           "2",
				//	TotalPrice:        "200.00",
				//	PreferentialPrice: "50.00",
				//	TotalPayPrice:     "150.00",
				//}
				},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipMultiPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "b565f726-ff0f-434d-aa31-e3e3cee15c0c",
		SkuId:             "sku1",
		VariantId:         "var1",
		ItemName:          "Gorgeous Bronze Knife",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}

	params := &xmhsdk.ShipParam{
		OrderId:    "1741867486",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p1.AddOneYearPp("10.00"), p2.AddTwoYearPp("10.00"), p3.AddThreeYearPp("10.00")},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestConfuseShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()

	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	params := &xmhsdk.ShipParam{
		OrderId:    "1741779739",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p1, p2, p3},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestConfuseAllShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()

	p1 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p2 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	p3 := &xmhsdk.DItem{
		ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
		SkuId:             "SKU001",
		ItemName:          "Durable Concrete Shirt",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "1",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
	}
	params := &xmhsdk.ShipParam{
		OrderId:    "1741932705",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{p1, p2, p3},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestEuShipSpOrder(t *testing.T) {
	xmhsdk.AppId = "10001"
	xmhsdk.AppSecret = "sX6QltIfMJEobQqxsdRFt3w2Jr8jlZaM"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	params := &xmhsdk.ShipParam{
		OrderId: "1746698684",
		//SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:          strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode: "ups",
				ShipCompany:     "ups",
				ShipTrackNumber: strconv.Itoa(int(time.Now().Unix())),
				ShipStateString: "已发货",
				ShipPrice:       "6.00",
				//ActualShipSendTime: time.Now().Format(time.RFC3339),
				ActualShipSendTime: "2025-04-11T05:37:10.000Z",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:    "United States",
					CityCode:   "US",
					Province:   "New Jersey",
					PostalCode: "NJ",
					City:       "Union City",
					//CountryCode: "CN",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId:    "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:     "SKU001",
					ItemName:  "Durable Concrete Shirt",
					Currency:  "USD",
					UnitPrice: "0",
					UnitNum:   "2",
					//TotalPrice:        "0",
					//PreferentialPrice: "0",
					TotalPayPrice: "150.00",
				}},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestEuShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000168"
	xmhsdk.AppSecret = "tGPJZpnI9MGgIyyGxqXuazDRQXCtx2GW"
	xmhsdk.SignSecret = "vevor-alpha"
	xmhsdk.SetEnv(xmhsdk.EnvAlpha)
	auth.New()
	item := &xmhsdk.DItem{
		ItemId: "16GA5QRB5TO01_110",
		//SkuId:             "SKU001",
		OrderGoodsId:      "16GA5QRB5TO01_110",
		ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
		InsuredPayPrice:   "24.0",
	}

	params := &xmhsdk.ShipParam{
		OrderId:    "1746685138",
		SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: time.Date(2024, 2, 29, 11, 11, 11, 11, time.Local).Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					//City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{item},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestBatch(t *testing.T) {
	xmhsdk.AppId = "1600398"
	xmhsdk.AppSecret = "BTnKVWMdYgE37yyB3C2cU65WyyZXsEWR"
	xmhsdk.SignSecret = "vevor-idc"
	xmhsdk.SetEnv(xmhsdk.EnvIdc)
	auth.New()
	type OpenApiOrderShip struct {
		OrderSN               string `json:"order_sn"`
		LogisticsSN           string `json:"logistics_sn"`
		LogisticsBusinessCode string `json:"logistics_business_code"`
		DeliveryTimeBj        string `json:"delivery_time_bj"`
	}

	bytes, rErr := os.ReadFile("./output.json")
	if rErr != nil {
		xmhsdk.Logger.Errorf("read json file:%s", rErr.Error())
		return
	}
	var orderShips []OpenApiOrderShip
	uErr := json.Unmarshal(bytes, &orderShips)
	if uErr != nil {
		xmhsdk.Logger.Errorf("decode json data:%s", uErr.Error())
		return
	}
	for i, orderShip := range orderShips {
		beijingLoc, _ := time.LoadLocation("Asia/Shanghai")
		parse, dErr := time.ParseInLocation(time.DateTime, orderShip.DeliveryTimeBj, beijingLoc)
		if dErr != nil {
			xmhsdk.Logger.Errorf("parse time error:%s", dErr.Error())
			return
		}
		_, sErr := New(&xmhsdk.ShipParam{
			OrderId:    orderShip.OrderSN,
			SubOrderId: orderShip.OrderSN,
			ShipInfoList: []*xmhsdk.ShipInfo{
				{
					ShipId:             fmt.Sprintf("%v", orderShip.LogisticsSN),
					ActualShipSendTime: parse.Format(time.RFC3339),
					ShipTrackNumber:    orderShip.LogisticsSN,
					ShipCompany:        orderShip.LogisticsBusinessCode,
					ShipCompanyCode:    orderShip.LogisticsBusinessCode,
					ItemList:           []*xmhsdk.DItem{{}},
				},
			},
		})
		if sErr != nil {
			xmhsdk.Logger.Errorf("ship idx:%d error %s", i, sErr.Error())
			return
		}
	}
}

func TestShipUsBetaSpOrder(t *testing.T) {
	xmhsdk.AppId = "1000160"
	xmhsdk.AppSecret = "xzs6We8YcKdpHrGQn8XHpyGCNg0a7sd7"
	xmhsdk.SignSecret = "ex_xmh_test"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	params := &xmhsdk.ShipParam{
		OrderId:    "1770605779",
		SubOrderId: "",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:          strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode: "ups",
				ShipCompany:     "ups",
				//ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipTrackNumber: "YT2529200702482310",

				ShipStateString:    "已发货",
				ShipPrice:          "0.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: "2025-10-19T14:03:12+08:00",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:    "United States",
					CityCode:   "US",
					Province:   "New Jersey",
					PostalCode: "NJ",
					City:       "Union City",
					//CountryCode: "CN",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId: "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:  "SKU001",
					//OrderGoodsId:      "SKU001000001",
					ItemName:          "Durable Concrete Shirt",
					Currency:          "USD",
					UnitPrice:         "0",
					UnitNum:           "1",
					TotalPrice:        "150",
					PreferentialPrice: "0",
					TotalPayPrice:     "150.00",
				}},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}

func TestShipUsBetaPpOrder(t *testing.T) {
	xmhsdk.AppId = "1000160"
	xmhsdk.AppSecret = "xzs6We8YcKdpHrGQn8XHpyGCNg0a7sd7"
	xmhsdk.SignSecret = "ex_xmh_test"
	xmhsdk.SetEnv(xmhsdk.EnvUsBeta)
	auth.New()
	item := &xmhsdk.DItem{
		ItemId: "1FCALEORCDO01_100",
		//SkuId:             "SKU001",
		//OrderGoodsId:      "16GA5QRB5TO01_110",
		ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
		Currency:          "USD",
		UnitPrice:         "100.00",
		UnitNum:           "2",
		TotalPrice:        "200.00",
		PreferentialPrice: "50.00",
		TotalPayPrice:     "150.00",
		InsuredPayPrice:   "24.0",
	}
	//item1 := &xmhsdk.DItem{
	//	ItemId: "1FCALEORCDO01_100",
	//	//SkuId:             "SKU001",
	//	OrderGoodsId:      "16GA5QRB5TO01_110",
	//	ItemName:          "PawHut 2-tier Wood Rabbit Hutch Backyard Bunny Cage Small Animal House w/ Ramp and Outdoor Run",
	//	Currency:          "USD",
	//	UnitPrice:         "100.00",
	//	UnitNum:           "2",
	//	TotalPrice:        "200.00",
	//	PreferentialPrice: "50.00",
	//	TotalPayPrice:     "150.00",
	//	InsuredPayPrice:   "24.0",
	//}
	intn := rand.New(rand.NewSource(time.Now().Unix()))

	params := &xmhsdk.ShipParam{
		OrderId:    "1770608219",
		SubOrderId: "",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: time.Date(2024, 2, 29, 11, 11, 11, 11, time.Local).Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					//City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{item},
			},
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipCompanyCode:    "SF",
				ShipCompany:        "SF",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix()) + intn.Int()),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: time.Date(2024, 2, 29, 11, 11, 11, 11, time.Local).Format(time.RFC3339),
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:  "中国",
					Province: "浙江省",
					//City:     "杭州市",
				},
				ItemList: []*xmhsdk.DItem{item},
			},
		},
	}
	result, err := New(params)
	if err != nil {
		t.Errorf("Ship error: %v", err)
		return
	}
	t.Logf("Ship result: %v", result)

}
