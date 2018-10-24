package client

import "encoding/json"

// 物流确认请求
type LogisticSyncRequest struct {
	OrderId string `json:"orderId"` // 必须,订单编号
	Status int `json:"status"` // 必须,物流发货状态:0=部分确认,1=确认完成
	LogisticsList []Logistics `json:"logisticsList"` // 必须,物流列表,见Logistics
}

// 物流
type Logistics struct {
	SkuNo string `json:"skuNo"` // 必须,商品编码
	LogisticsCompanyNo string `json:"logisticsCompanyNo"` // 必须,物流公司编号
	LogisticsNumber string `json:"logisticsNumber"` // 必须,运单号
	Qty int `json:"qty"` // 可选,该运单号发货数量,大于0的正整数,数量须和商品购买量一致(如多物流,累加值和购买量一致)
	UniqNos []string `json:"uniqNos"` // 可选,回传的货品sn码信息,须和商品购买量一致(如多物流,累加sn数量值和购买量一致),形如[“sn001”,”sn002”]
}

// API名称
func (*LogisticSyncRequest) GetApiName() (apiName string, err error){
	return "logistic.sync", nil
}

// 业务参数
func (r *LogisticSyncRequest) GetParam() (param string, err error){
	bytes, err := json.Marshal(r)
	if err != nil{
		return
	}
	param = string(bytes)
	return
}

// 物流确认
func (client *Client) LogisticSync (request *LogisticSyncRequest) (response *LogisticSyncResponse, err error)  {
	response = new(LogisticSyncResponse)
	err = client.Execute(request, response)
	if err != nil{
		return
	}
	return
}
