package client

// 物流确认响应
type LogisticSyncResponse struct {
	*BasicResponse
	Data string `json:"data"` // none
}

// 响应体
func (*LogisticSyncResponse) ResponseBody() (response string, err error){
	return "", nil
}
