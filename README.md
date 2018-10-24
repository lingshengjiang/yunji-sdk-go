# yunji-sdk-go

云集SDK,以方便使用,基于[Go]实现.

## 接口文档

- [代销开放平台] - 接口文档

## 使用

``` go
// 获取订单列表
func GetOrderList(t *testing.T) {		
	// 客户端
	client, err := NewClient(test.ServerUrl, test.AppKey, test.AppSecret)
	if err != nil{
		panic(err.Error())
	}
	
	// 请求参数
	request := new(OrderListRequest)
	request.Status = "40,50"
	request.StartTime = time.Now().AddDate(-1, 0,0).Format("2006-01-02 15:04:05")
	request.EndTime = time.Now().Format("2006-01-02 15:04:05")
	request.PageNo = 1
	request.PageSize = 1

	// 响应
	response, err := client.OrderList(request)
	if err != nil{
		panic(err.Error())
	}
	
	// todo:...
	if response.Code != 0{
		//
	}
}
```

[Go]:https://golang.google.cn/
[代销开放平台]:http://api.yunjiweidian.com/showdoc/index.php?s=/1

