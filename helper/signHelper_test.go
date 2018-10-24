package helper

import "testing"

func TestGetSign(t *testing.T){
	paramMap := make(map[string]string)
	paramMap["param"] = "=param&"
	paramMap["appkey"] = "=appkey&"
	paramMap["timestamp"] = "=timestamp&"
	paramMap["v"] = "1.0"
	paramMap["method"] = "=method"

	sign, err := GetSign(paramMap, "=appSecret")
	if err != nil{
		t.Errorf(err.Error())
	}
	t.Logf("sign=%s", sign)
}
