package client

import (
	"testing"
	"time"
	"github.com/taadis/yunji-sdk-go/config"
)

//
func TestOrderDetail1(t *testing.T) {
	request := new(OrderDetailRequest)

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil {
		t.Error(err)
	}

	response, err := client.OrderDetail(request)
	if err != nil {
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:orderId"}

	if response.Code != 3 {
		t.Error(response.Desc)
	}
}

//
func TestOrderDetail2(t *testing.T)  {
	orderListRequest := new(OrderListRequest)
	orderListRequest.QueryType = 1
	orderListRequest.EndTime = time.Now().Format("2006-01-02 15:04:05")
	orderListRequest.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	orderListRequest.PageSize = 1
	orderListRequest.PageNo = 1
	orderListRequest.Status = "40,50"

	client,err := NewClient(config.ServerUrl,config.AppKey, config.AppSecret)
	if err != nil{
		t.Error(err)
	}

	orderListResponse, err := client.OrderList(orderListRequest)
	if err != nil{
		t.Error(err)
	}

	request := new(OrderDetailRequest)
	request.OrderId = orderListResponse.Data.List[0].OrderId

	response, err := client.OrderDetail(request)
	if err != nil{
		t.Error(err)
	}
	// response = success

	if response.Code != 0{
		t.Error(response.Desc)
	}
}
