package client

// 响应器
type Responser interface {

	// 响应体
	ResponseBody() (response string, err error)
}
