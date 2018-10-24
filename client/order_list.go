package client

import "encoding/json"

// 订单列表请求
type OrderListRequest struct {
	StartTime string `json:"startTime"` // 必须,起始时间,格式:YYYY-MM-DD HH:MM:SS
	EndTime   string `json:"endTime"`   // 必须,起始时间,格式:YYYY-MM-DD HH:MM:SS
	PageSize  int    `json:"pageSize"`  // 必须,每页的条数,默认值=10,最大值=50
	PageNo    int    `json:"pageNo"`    // 必须,页码,默认值=1
	Status    string `json:"status"`    // 必须,订单状态:40=待发货(默认值),50=已发货,多个状态获取逗号隔开
	QueryType int    `json:"queryType"` // 可选,请求时间类型:0=创建时间(默认值),1=修改时间
}

// API名称
func (r *OrderListRequest) GetApiName() (apiName string, err error){
	return "order.list", nil
}

// 业务参数
func (r *OrderListRequest) GetParam() (param string, err error){

	bytes, err := json.Marshal(r)
	if err!= nil{
		return
	}

	param = string(bytes)
	return
}

// 响应体
func (r *OrderListResponse)ResponseBody() (response string, err error){
	return "", nil
}

// 订单列表响应
type OrderListResponse struct {
	*BasicResponse
	Data OrderListData `json:"data"` // 必须,订单列表数据
}

// 订单列表数据
type OrderListData struct {
	Total int `json:"total"` // 必须,总条数
	List []Order `json:"list"` // 必须,订单列表
}

// 接口方法名称
func (o *OrderListRequest) ApiName() string {
	return "order.list"
}

// 订单列表
func (client *Client) OrderList(request *OrderListRequest) (response *OrderListResponse, err error) {
	response = new(OrderListResponse)
	err = client.Execute(request, response)
	if err != nil {
		return
	}
	return
}
