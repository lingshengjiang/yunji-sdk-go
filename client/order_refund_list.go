package client

import "encoding/json"

// 售后列表请求
type OrderRefundListRequest struct{
	StartTime string `json:"startTime"` // 必须,起始的修改时间,格式YYYY-MM-DD HH:MM:SS
	EndTime string `json:"endTime"` // 必须,结束的修改时间,格式YYYY-MM-DD HH:MM:SS
	PageSize int `json:"pageSize"` // 必须,每页的条数,默认值=10,最大=50
	PageNo int `json:"pageNo"` // 必须,页码,默认值=1
	ReturnStatus string `json:"returnStatus"` // 订单退款状态,0=无退款,1=申请退款,3=已退款,4=已关闭,5=确认收货,多个状态获取逗号隔开,默认值=0
	QueryType int `json:"queryType"` // 可选,请求时间类型,0=创建时间,1=修改时间,建议修改时间查询,默认值=0
}

// 获取接口名称
func (request *OrderRefundListRequest) GetApiName() (apiName string, err error)  {
	return "order.refund.list", nil
}

// 获取参数
func (request *OrderRefundListRequest) GetParam()(param string, err error)  {

	// 业务请求参数结构体转JSON字符串
	paramBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	param = string(paramBytes)

	return
}

// 响应体
func (response *OrderRefundListResponse) ResponseBody() (responseBody string, err error)  {
	return "",nil
}

// 售后列表响应
type OrderRefundListResponse struct {
	*BasicResponse
	Data OrderRefundListData `json:"data"`
}

// 售后列表数据
type OrderRefundListData struct {
	Total int `json:"total"` // 必须,总条数
	RefundOrderList []RefundOrderItem `json:"returnorderlist"` // 必须,售后订单列表
}

// 售后列表
func (client *Client) OrderRefundList(request *OrderRefundListRequest) (response *OrderRefundListResponse,err error)  {
	response = new(OrderRefundListResponse)
	err = client.Execute(request, response)
	if err != nil{
		return
	}
	return
}

