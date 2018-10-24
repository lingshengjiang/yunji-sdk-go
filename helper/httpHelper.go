package helper

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// httpPost
func HttpPost(url string, param string)  (response string, err error){
	res, err := http.Post(url,"application/x-www-form-urlencoded", strings.NewReader(param))
	if err != nil{
		return
	}
	defer res.Body.Close()

	body , err := ioutil.ReadAll(res.Body)
	if err != nil{
		return
	}

	response = string(body)
	return
}

