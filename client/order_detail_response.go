package client

// 订单详情响应
type OrderDetailResponse struct {
	*BasicResponse
	Data Order `json:"data"` // 必须,订单数据,见结构体Order
}

// 响应体
func (*OrderDetailResponse) ResponseBody() (response string, err error){
	return "",nil
}
