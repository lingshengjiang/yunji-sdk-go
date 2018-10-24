package yunji

import (
	"testing"
	"time"
	"yunji/test"
)

//
func TestLogisticSync1(t *testing.T)  {
	request := new(LogisticSyncRequest)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:orderId"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync2(t *testing.T)  {
	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:logisticsList"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync3(t *testing.T)  {
	logistics := new(Logistics)

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:skuNo"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync4(t *testing.T)  {
	logistics := new(Logistics)
	logistics.SkuNo = "test"

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:logisticsCompanyNo"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync5(t *testing.T)  {
	logistics := new(Logistics)
	logistics.SkuNo = "test"
	logistics.LogisticsCompanyNo = "ems"

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:logisticsNumber"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync6(t *testing.T)  {
	logistics := new(Logistics)
	logistics.SkuNo = "test"
	logistics.LogisticsCompanyNo = "test"
	logistics.LogisticsNumber = "test"

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Invalid arguments:the value of qty can not be less than 1."}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync7(t *testing.T)  {
	logistics := new(Logistics)
	logistics.SkuNo = "test"
	logistics.LogisticsCompanyNo = "test"
	logistics.LogisticsNumber = "test"
	logistics.Qty = 1

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":1,"desc":"发货失败 失败原因如下:[\"test物流公司编码不存在\"]"}
	// response = {"code":1,"desc":"发货失败 失败原因如下:[\"test物流公司编码不存在\"]"}

	if response.Code != 1{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync8(t *testing.T)  {
	logistics := new(Logistics)
	logistics.SkuNo = "test"
	logistics.LogisticsCompanyNo = "EMS"
	logistics.LogisticsNumber = "test"
	logistics.Qty = 1

	request := new(LogisticSyncRequest)
	request.OrderId = "46546416"
	request.LogisticsList = append(request.LogisticsList, *logistics)

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":1,"desc":"发货失败 失败原因如下:[\"46546416订单不存在\"]"}

	if response.Code != 1{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync9(t *testing.T)  {
	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	// 先获取待发货的订单数据
	orderListRequest := new(OrderListRequest)
	orderListRequest.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	orderListRequest.EndTime = time.Now().Format("2006-01-02 15:04:05")
	orderListRequest.PageSize = 1
	orderListRequest.PageNo = 1
	orderListRequest.Status = "40"
	orderListRequest.QueryType = 1

	orderListResponse, err := client.OrderList(orderListRequest)
	if err != nil{
		t.Error(err)
	}

	// 物流确认
	logistics := new(Logistics)
	logistics.SkuNo = "test"
	logistics.LogisticsCompanyNo = "EMS"
	logistics.LogisticsNumber = "test"
	logistics.Qty = 1

	request := new(LogisticSyncRequest)
	request.OrderId = orderListResponse.Data.List[0].OrderId
	request.LogisticsList = append(request.LogisticsList, *logistics)

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":1,"desc":"发货失败 失败原因如下:[\"skuNo不存在 订单号与skuNo无法正确匹

	if response.Code != 1{
		t.Error(response.Desc)
	}
}

//
func TestLogisticSync10(t *testing.T)  {
	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Error(err)
	}

	// 先获取待发货的订单数据
	orderListRequest := new(OrderListRequest)
	orderListRequest.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	orderListRequest.EndTime = time.Now().Format("2006-01-02 15:04:05")
	orderListRequest.PageSize = 1
	orderListRequest.PageNo = 1
	orderListRequest.Status = "40"
	orderListRequest.QueryType = 1

	orderListResponse, err := client.OrderList(orderListRequest)
	if err != nil{
		t.Error(err)
	}

	// 物流确认
	logistics := new(Logistics)
	logistics.SkuNo = orderListResponse.Data.List[0].OrderDetail[0].SkuNo
	logistics.LogisticsCompanyNo = "EMS"
	logistics.LogisticsNumber = "test"
	logistics.Qty = orderListResponse.Data.List[0].OrderDetail[0].Qty

	request := new(LogisticSyncRequest)
	request.OrderId = orderListResponse.Data.List[0].OrderId
	request.LogisticsList = append(request.LogisticsList, *logistics)

	response, err := client.LogisticSync(request)
	if err != nil{
		t.Error(err)
	}
	// response = success

	if response.Code != 0{
		t.Error(response.Desc)
	}
}
