package client

import "encoding/json"

// 订单详情请求
type OrderDetailRequest struct {
	OrderId string `json:"orderId"` // 必须,订单编号
}

// API名称
func (*OrderDetailRequest) GetApiName() (apiName string, err error){
	return "order.detail",nil
}

// 业务参数
func (r *OrderDetailRequest) GetParam() (param string, err error){
	bytes, err := json.Marshal(r)
	if err != nil{
		return
	}
	param = string(bytes)
	return
}

//
func (client *Client) OrderDetail(request *OrderDetailRequest) (response *OrderDetailResponse, err error) {
	response = new(OrderDetailResponse)
	err = client.Execute(request, response)
	if err != nil{
		return
	}
	return
}
