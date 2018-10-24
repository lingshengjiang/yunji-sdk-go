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

## Go Module Publish

每当需要发布新版本时,例如以下情况

- 新功能/特性
- 不兼容性
- bug修复
- ...

都应该发布一个新的版本来规避可能的问题,当然不要忘了测试

发布操作如下:

``` git
// 打个 tag, 注意版本语义
git tag v0.0.1

// push 上去
git push --tags
```
