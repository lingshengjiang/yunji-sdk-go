# ...

## 注意事项

每一个请求需要实现以下 `interface`

``` go
// 请求器
type Requester interface {

	// API名称
	GetApiName() (apiName string, err error)

	// 业务参数
	GetParam() (param string, err error)
}
```

每一个响应需要实现以下 `interface` (感觉不大对劲...)

``` go
// 响应器
type Responser interface {

	// 响应体
	ResponseBody() (response string, err error)
}
```
