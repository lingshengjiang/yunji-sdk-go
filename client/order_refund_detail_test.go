package client

import (
	"github.com/taadis/yunji-sdk-go/config"
	"testing"
	"time"
)

//
func TestOrderRefundDetail1(t *testing.T){
	request := new(OrderRefundDetailRequest)

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Error(err)
	}

	response, err := client.OrderRefundDetail(request)
	if err != nil{
		t.Error(err)
	}
	// response = {"code":3,"desc":"data error:Missing required arguments:orderId"}

	if response.Code != 3{
		t.Error(response.Desc)
	}
}

//
func TestOrderRefundDetail2(t *testing.T){
	orderRefundListRequest := new(OrderRefundListRequest)
	orderRefundListRequest.PageNo = 1
	orderRefundListRequest.PageSize = 1
	orderRefundListRequest.EndTime = time.Now().Format("2006-01-02 15:04:05")
	orderRefundListRequest.StartTime = time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05")
	orderRefundListRequest.ReturnStatus = "0,1,3,4,5"
	orderRefundListRequest.QueryType = 1

	client, err := NewClient(config.ServerUrl, config.AppKey, config.AppSecret)
	if err != nil{
		t.Error(err)
	}

	orderRefundListResponse, err := client.OrderRefundList(orderRefundListRequest)
	if err != nil{
		t.Error(err)
	}

	request := new(OrderRefundDetailRequest)
	request.OrderId = orderRefundListResponse.Data.RefundOrderList[0].OrderId

	response, err := client.OrderRefundDetail(request)
	if err != nil{
		t.Error(err)
	}
	// response = success

	if response.Code != 0{
		t.Error(response.Desc)
	}
}
