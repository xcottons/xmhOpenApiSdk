package item

import (
	"fmt"
	xmhsdk "github.com/cjay-shouhui/xmhOpenApiSdk"
)

const platformItemURL = "SyncPlatformProduct"
const MaxItemBatchSize = 500

var batchId = uint64(0)

func Sync(params *xmhsdk.ProductItemsParam) (*struct{}, error) {
	baseParams := *params
	items := baseParams.Items
	total := len(items)
	if total == 0 {
		return &struct{}{}, nil
	}

	batchSize := MaxItemBatchSize
	batchId := 0
	successCount := 0
	for start := 0; start < total; {
		end := start + batchSize
		if end > total {
			end = total
		}
		batchParams := baseParams
		batchParams.Items = items[start:end]
		batchParams.BatchId = uint64(batchId)
		result := &xmhsdk.PlatformOrderResult{}
		if err := xmhsdk.MakeRequest(platformItemURL, &batchParams, result); err != nil {
			return nil, fmt.Errorf("batch %d failed: %w (success:%d/%d)",
				batchId, err, successCount, total)
		}
		successCount += len(batchParams.Items)
		batchId++
		start = end
	}
	return &struct{}{}, nil
}
