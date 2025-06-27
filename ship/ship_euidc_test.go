package ship

import (
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
	"strconv"
	"testing"
	"time"
)

func TestEuIdcShipSpOrder(t *testing.T) {
	xmhsdk.AppId = "1001"
	xmhsdk.AppSecret = "WkC2LaMgs7ckzxsBtDNSJR9YZb9B0wU3"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvEuIdc)
	auth.New()
	params := &xmhsdk.ShipParam{
		OrderId: "1747125774",
		//SubOrderId: "SubOrderId6",
		ShipInfoList: []*xmhsdk.ShipInfo{
			{
				ShipId:             strconv.Itoa(int(time.Now().Unix())),
				ShipCompanyCode:    "ups",
				ShipCompany:        "ups",
				ShipTrackNumber:    strconv.Itoa(int(time.Now().Unix())),
				ShipStateString:    "已发货",
				ShipPrice:          "6.00",
				ActualShipSendTime: time.Now().Format(time.RFC3339),
				//ActualShipSendTime: "2025-04-11T05:37:10.000Z",
				ShipOtherInfo: &xmhsdk.ShipAddress{
					Country:    "United States",
					CityCode:   "US",
					Province:   "New Jersey",
					PostalCode: "NJ",
					City:       "Union City",
					//CountryCode: "CN",
				},
				ItemList: []*xmhsdk.DItem{{
					ItemId:            "97760109-16ad-40c9-9385-caba381a26aa",
					SkuId:             "SKU001",
					ItemName:          "Ulanzi 40W Portable LED Video Light Bundle - L024 40W RGB Light Bundle / US Plu",
					Currency:          "USD",
					UnitPrice:         "229.95",
					UnitNum:           "1",
					TotalPrice:        "202.36",
					PreferentialPrice: "27.59",
					TotalPayPrice:     "202.36",
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

func TestEuIdcShipPpOrder(t *testing.T) {
	xmhsdk.AppId = "1001"
	xmhsdk.AppSecret = "WkC2LaMgs7ckzxsBtDNSJR9YZb9B0wU3"
	xmhsdk.SignSecret = "yujianfx-eu"
	xmhsdk.SetEnv(xmhsdk.EnvEuIdc)
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

	params := &xmhsdk.ShipParam{
		OrderId:    "1747125542",
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
