package znet

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
	"io/ioutil"
	"net/http"
	"strings"
)

type ErrorStatusCode struct {
	Code int
	Msg  string
}

func (err ErrorStatusCode) Error() string {
	return z.String("code:", err.Code, ",msg:", err.Msg)
}

// 把结果转化为json的map
func getMap(data []byte, err1 error) (z.Map, error) {
	if result, err := gjson.Decode(data); err == nil {
		json := result.(map[string]interface{})
		return json, err1
	} else {
		return nil, err
	}

}

func DoRequest(method, url string, data string, header z.Map) ([]byte, error) {
	client := GetHttpClient()
	defer FreeHttpClient(client)
	body := (*strings.Reader)(nil)
	body = strings.NewReader(data)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		zlog.LogError(err, "getEmail", "NewRequest")
		return nil, err
	}
	if header != nil {
		for key, value := range header {
			request.Header.Add(key, gconv.String(value))
		}
	}
	response, err := client.Do(request)
	if err != nil && response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		zlog.LogError(err, "请求失败")
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		zlog.LogError(err, "WebRequest", "ioutil.ReadAll(response.Body)")
		return nil, err
	}
	if response.StatusCode != 200 {
		return result, ErrorStatusCode{response.StatusCode, response.Status}
	}

	//if response.StatusCode != 200 {
	//	zlog.Log("请求失败，status:"+response.Status, request.URL.String())
	//	resultA, _ := ioutil.ReadAll(response.Body)
	//	zlog.Log("url", url)
	//	zlog.Log("data", data)
	//	zlog.Log("header", header)
	//	zlog.Log("body", string(resultA))
	//
	//	return result, errors.New("请求失败，status:" + response.Status)
	//}

	return result, err
}

// Get请求
func GetJson(url string, body string, header z.Map) (z.Map, error) {
	data, err := DoRequest("GET", url, body, header)
	if err != nil {
		switch err.(type) {
		case ErrorStatusCode:
			return getMap(data, err)
		default:
			return nil, err
		}
	}
	return getMap(data, err)

}

// Post 请求
func PostJson(url string, body string, header z.Map) (z.Map, error) {
	data, err := DoRequest("POST", url, body, header)
	if err != nil {
		switch err.(type) {
		case ErrorStatusCode:
			return getMap(data, err)
		default:
			return nil, err
		}
	}
	return getMap(data, err)
}
