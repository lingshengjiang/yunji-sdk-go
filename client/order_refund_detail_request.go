package client

import "encoding/json"

// 售后详情请求
type OrderRefundDetailRequest struct {
	RefundId string `json:"refundId"` // 必须,售后编号,orderId和refundId必须传入一个
	OrderId string `json:"orderId"`	// 必须,订单号,可能返会多笔售后,orderId和refundId必须传入一个
}

// API名称
func (r *OrderRefundDetailRequest) GetApiName() (apiName string, err error){
	return "order.refund.detail", nil
}

// 业务参数
func (r *OrderRefundDetailRequest) GetParam() (param string, err error){
	bytes, err := json.Marshal(r)
	if err !=nil{
		return
	}
	param = string(bytes)
	return
}

// 售后详情
func (client *Client) OrderRefundDetail (request *OrderRefundDetailRequest) (response *OrderRefundDetailResponse, err error)  {
	response = new(OrderRefundDetailResponse)
	err = client.Execute(request, response)
	if err != nil{
		return
	}
	return
}
