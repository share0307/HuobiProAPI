package models

type TimestampReturn struct {
	Status string `json:"status"` // 请求状态
	Data   int64  `json:"data"`   // 时间戳
}
