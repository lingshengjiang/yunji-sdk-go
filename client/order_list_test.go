package yunji

import (
	"testing"
	"time"
	"yunji/test"
)

//
func TestOrderList1(t *testing.T) {
	// 请求参数
	request := new(OrderListRequest)
	// 不设置任何请求参数

	client, err := NewClient(test.ServerUrl, "test", "test")
	if err != nil{
		t.Error(err)
	}

	// 执行
	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":1,"desc":"request invalid!"}

	if response.Code != 1 {
		t.Error(response.Desc)
	}
}

//
func TestOrderList2(t *testing.T) {
	// 请求参数
	request := new(OrderListRequest)
	// 不设置任何请求参数

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	// 执行
	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:status"}

	if response.Code != 3{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList3(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:startTime"}

	if response.Code != 3{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList4(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().Format("2006-01-02 15:04:05")

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err !=nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:endTime"}

	if response.Code != 3{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList5(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2006-01-02 15:04:05")

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Invalid arguments:the value of pageNo can not be less than 1."}

	if response.Code != 3{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList6(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2006-01-02 15:04:05")
	request.PageNo = 1

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":0,"desc":"success","data":{"total":0,"list":[]}}

	if response.Code != 0{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList7(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().AddDate(-1, 0,0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2006-01-02 15:04:05")
	request.PageNo = 1

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":0,"desc":"success","data":{"total":969,"list":[]}}

	if response.Code != 0{
		t.Errorf(err.Error())
	}
}

//
func TestOrderList8(t *testing.T) {
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().AddDate(-1, 0,0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2006-01-02 15:04:05")
	request.PageNo = 1
	request.PageSize = 1

	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":0,"desc":"success","data":{"total":969,"list":[]}}

	if response.Code != 0{
		t.Errorf(err.Error())
	}
}
