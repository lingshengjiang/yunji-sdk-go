package helper

import "testing"

func TestHttpPost(t *testing.T)  {
	response, err :=  HttpPost("http://t.yunjiglobal.com/yunjiesb/api/esb/insteadsale/service/", "")
	if err != nil{
		t.Errorf("请求失败=%s", err.Error())
	}
	t.Logf("请求成功=%s", response)
}
