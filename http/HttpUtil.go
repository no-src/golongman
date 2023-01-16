package _http

import (
	"errors"
	"io"
	"net/http"
)

var (
	NotSuccessError = errors.New("HTTP响应非200")
	BodyNilError    = errors.New("响应Body为nil")
	BodyReadError   = errors.New("读取Body数据异常")
	StringEmpty     = ""
)

// Get 发起Get请求，并返回body
func Get(url string) (body io.ReadCloser, err error) {
	resp, err := http.Get(url)
	//defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, NotSuccessError
	}
	return resp.Body, nil
}

// GetString 发起Get请求，返回响应的字符结果
func GetString(url string) (result string, err error) {
	body, err := Get(url)
	if err != nil {
		return StringEmpty, err
	}
	if body == nil {
		return StringEmpty, BodyNilError
	}
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return StringEmpty, BodyReadError
	}
	respContent := string(bodyBytes)
	return respContent, nil
}
