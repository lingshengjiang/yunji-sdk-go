package client

// 请求器
type Requester interface {

	// API名称
	GetApiName() (apiName string, err error)

	// 业务参数
	GetParam() (param string, err error)
}
