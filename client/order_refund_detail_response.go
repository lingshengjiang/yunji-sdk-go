package client

// 售后详情响应
type OrderRefundDetailResponse struct {
	*BasicResponse
	Data []RefundOrderItem `json:"data"` // 必须,售后详情数据,见ReturnOrder描述
}

// 响应体
func (*OrderRefundDetailResponse) ResponseBody() (response string, err error){
	return "",nil
}
