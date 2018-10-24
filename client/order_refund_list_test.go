package client

import (
	"github.com/taadis/yunji-sdk-go/config"
	"testing"
	"time"
)

//
func TestOrderRefundListRequest1(t *testing.T) {

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	request := new(OrderRefundListRequest)

	response,err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:returnStatus"}

	if response.Code != 3{
		t.Errorf(response.Desc)
	}
}

//
func TestOrderRefundListRequest2(t *testing.T){

	request := new(OrderRefundListRequest)
	request.ReturnStatus = "0,1,3,4,5"

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:startTime"}

	if response.Code != 3{
		t.Errorf(response.Desc)
	}
}

//
func TestOrderRefundListRequest3(t *testing.T){

	request := new(OrderRefundListRequest)
	request.ReturnStatus = "0,1,3,4,5"
	request.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:endTime"}

	if response.Code != 3{
		t.Errorf(response.Desc)
	}
}

//
func TestOrderRefundListRequest4(t *testing.T){

	request := new(OrderRefundListRequest)
	request.ReturnStatus = "0,1,3,4,5"
	request.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2016-01-02 15:04:05")

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:pageSize"}

	if response.Code != 3{
		t.Errorf(response.Desc)
	}
}

//
func TestOrderRefundListRequest5(t *testing.T){

	request := new(OrderRefundListRequest)
	request.ReturnStatus = "0,1,3,4,5"
	request.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2016-01-02 15:04:05")
	request.PageSize = 1

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = {"code":3,"desc":"data error:Invalid arguments:the value of pageNo can not be less than 1."}

	if response.Code != 3{
		t.Errorf(response.Desc)
	}
}

//
func TestOrderRefundListRequest6(t *testing.T){

	request := new(OrderRefundListRequest)
	request.ReturnStatus = "0,1,3,4,5"
	request.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2016-01-02 15:04:05")
	request.PageSize = 1
	request.PageNo = 1

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Errorf(err.Error())
	}

	response, err := client.OrderRefundList(request)
	if err != nil{
		t.Errorf(err.Error())
	}
	// response = success

	if response.Code != 0{
		t.Errorf(response.Desc)
	}
}