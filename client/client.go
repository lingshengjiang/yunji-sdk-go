package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
	"github.com/taadis/yunji-sdk-go/helper"
)

//
type Client struct {
	serverUrl string
	appKey string
	appSecret string
	//httpClient *http.Client
}

//
func NewClient(serverUrl string, appKey string, appSecret string) (client *Client, err error) {
	// 参数验证
	if serverUrl == ""{
		err = errors.New("argument serverUrl equal null string")
		return
	}
	if appKey == ""{
		err = errors.New("argument appKey equal null string")
		return
	}
	if appSecret == ""{
		err = errors.New("argument appKey equal null string")
		return
	}

	client = new(Client)
	client.serverUrl = fmt.Sprintf("%s/%s", serverUrl, appKey)
	client.appKey = appKey
	client.appSecret = appSecret
	return
}

//
func (client *Client) Execute(request Requester, response Responser)(err error)  {
	paramMap := make(map[string]string)
	paramMap["param"], _ = request.GetParam()
	paramMap["appkey"] = client.appKey
	paramMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	paramMap["v"] = "1.0"
	paramMap["method"], _ = request.GetApiName()

	// 签名
	paramMap["sign"], err = helper.GetSign(paramMap, client.appSecret)
	if err != nil {
		return
	}

	// 发起 HTTP POST 请求
	postValues := url.Values{}
	for key, value := range paramMap {
		postValues.Set(key, value)
	}
	postDataString :=postValues.Encode()
	res, err := helper.HttpPost(client.serverUrl, postDataString)
	if err != nil {
		return
	}

	// JSON 字符串响应数据转结构体
	err = json.Unmarshal([]byte(res), &response)
	if err != nil {
		return
	}

	return
}
