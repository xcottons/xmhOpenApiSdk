package xmhOpenApiSdk

// ServiceOrderQueryParam 保单分页查询入参
type ServiceOrderQueryParam struct {
	OrderID      string `json:"orderId,omitempty"`      // 电商平台订单号
	FrontOrderID string `json:"frontOrderId,omitempty"` // 电商平台前端订单号
	BegTime      string `json:"begTime,omitempty"`      // 保单更新时间-开始，RFC3339，须与 EndTime 同时填写
	EndTime      string `json:"endTime,omitempty"`      // 保单更新时间-截止，RFC3339，须与 BegTime 同时填写
	PageSize     int    `json:"pageSize,omitempty"`     // 每页返回条数 1-1000，时间范围查询时必填
	Cursor       string `json:"cursor,omitempty"`       // 游标值，首次传 "0"，时间范围查询时必填
	Sort         int    `json:"sort,omitempty"`         // 0-正序（默认），1-逆序
	Region       string `json:"region,omitempty"`       // ISO 3166-1 ALPHA2，默认 CN
}

// ServiceOrderQueryResult 保单分页查询返回数据
type ServiceOrderQueryResult struct {
	TotalCount int64               `json:"totalCount"` // 当前页总记录数
	NextCursor string              `json:"nextCursor"` // 下次游标，值为空表示已到最后一页
	List       []*ServiceOrderInfo `json:"list"`       // 保单列表
}

// ServiceOrderInfo 保单信息
type ServiceOrderInfo struct {
	OrderId          string `json:"orderId"`          // 电商平台订单号
	FrontOrderId     string `json:"frontOrderId"`     // 电商平台前端订单号
	ServiceOrderId   string `json:"serviceOrderId"`   // 小棉花的服务单号
	GuaranteeBegTime string `json:"guaranteeBegTime"` // 保险起始时间（权益开始时间）RFC3339
	GuaranteeEndTime string `json:"guaranteeEndTime"` // 保险结束时间（权益结束时间）RFC3339
	PolicyStatus     int32  `json:"policyStatus"`     // 权益状态，见 XMH保单的状态 枚举
	PolicyStatusDesc string `json:"policyStatusDesc"` // 权益状态描述（中文）
	PolicyUpdateTime string `json:"policyUpdateTime"` // 保单更新时间 RFC3339
	ClaimState       int32  `json:"claimState"`       // 理赔状态，见 理赔状态 枚举
}
