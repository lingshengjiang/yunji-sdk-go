package client

// 基础响应
type BasicResponse struct {
	Code int    `json:"code"` // 必须,响应码,0=成功,其他皆为失败
	Desc string `json:"desc"` // 必须,描述信息
}
