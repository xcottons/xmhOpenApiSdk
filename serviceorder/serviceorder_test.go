package serviceorder

import (
	"testing"
	"time"

	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
	"github.com/cjay-shouhui/xmhOpenApiSdk/auth"
)

func initBetaEnvForServiceOrder(t *testing.T) {
	t.Helper()
	xmhsdk.AppId = "1000170"
	xmhsdk.AppSecret = "HsmqvWKfKajgFgPvZBRuwfmf9IcWQFkw"
	xmhsdk.SignSecret = "vevor_beta"
	xmhsdk.SetEnv(xmhsdk.EnvBeta)
	if _, err := auth.New(); err != nil {
		t.Fatalf("auth.New error: %v", err)
	}
}

func TestServiceOrderQueryByOrderID(t *testing.T) {
	initBetaEnvForServiceOrder(t)
	params := &xmhsdk.ServiceOrderQueryParam{
		OrderID:  "1770265268",
		PageSize: 100,
		Cursor:   "0",
	}
	result, err := Query(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ServiceOrderQuery by orderId error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ServiceOrderQuery by orderId result: %s", xmhsdk.ToStr(result))
}

func TestServiceOrderQueryByFrontOrderID(t *testing.T) {
	initBetaEnvForServiceOrder(t)
	params := &xmhsdk.ServiceOrderQueryParam{
		FrontOrderID: "1770265268",
		PageSize:     100,
		Cursor:       "0",
	}
	result, err := Query(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ServiceOrderQuery by frontOrderId error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ServiceOrderQuery by frontOrderId result: %s", xmhsdk.ToStr(result))
}

func TestServiceOrderQueryByTimeRange(t *testing.T) {
	initBetaEnvForServiceOrder(t)
	end := time.Now()
	begin := end.AddDate(0, 0, -7)
	params := &xmhsdk.ServiceOrderQueryParam{
		BegTime:  begin.Format(time.RFC3339),
		EndTime:  end.Format(time.RFC3339),
		PageSize: 100,
		Cursor:   "0",
		Sort:     1,
	}
	result, err := Query(params)
	if err != nil {
		xmhsdk.Logger.Errorf("ServiceOrderQuery by timeRange error: %v", err)
		return
	}
	xmhsdk.Logger.Infof("ServiceOrderQuery by timeRange result: %s", xmhsdk.ToStr(result))
	xmhsdk.Logger.Infof("totalCount=%d nextCursor=%s listLen=%d", result.TotalCount, result.NextCursor, len(result.List))

	// 翻页示例：当 nextCursor 非空时继续查询下一页
	if result.NextCursor != "" {
		params.Cursor = result.NextCursor
		nextResult, nextErr := Query(params)
		if nextErr != nil {
			xmhsdk.Logger.Errorf("ServiceOrderQuery page2 error: %v", nextErr)
		} else {
			xmhsdk.Logger.Infof("ServiceOrderQuery page2 result: %s", xmhsdk.ToStr(nextResult))
		}
	}
}
