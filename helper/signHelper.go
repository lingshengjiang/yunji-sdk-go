package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func GetSign(paramMap map[string]string, appSecret string) (sign string, err error) {

	// 定义排序切片,以便间接对paramMap进行按键排序
	sortSlice := make([]string, len(paramMap))
	index := 0
	for  key, _ := range paramMap {
		sortSlice[index] = key
		index++
	}

	// 排序
	sort.Strings(sortSlice)

	// 参数字符串拼接,头尾各加一次 appSecret
	joinStr := fmt.Sprintf("%s%s", appSecret, "")
	for _, value := range sortSlice  {
		joinStr = fmt.Sprintf("%s%s%s", joinStr, value, paramMap[value])
	}
	joinStr = fmt.Sprintf("%s%s", joinStr, appSecret)

	// MD5 & 16进制小写
	md5 := md5.New()
	md5.Write([]byte(joinStr))
	sign = strings.ToLower(hex.EncodeToString(md5.Sum(nil)))

	return
}
