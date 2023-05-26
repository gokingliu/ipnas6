package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var client = http.Client{
	Timeout: time.Duration(30) * time.Second, //超时时间
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   5,   //单个路由最大空闲连接数
		MaxConnsPerHost:       100, //单个路由最大连接数
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

func handleUriValues(args map[string]interface{}) string {
	if args != nil && len(args) > 0 {
		params := url.Values{}
		for k, v := range args {
			params.Set(k, fmt.Sprintf("%v", v))
		}
		return params.Encode()
	}

	return ""
}

func handleRequest(method string, uri string, headers map[string]string, body *bytes.Reader) *http.Request {
	request, _ := http.NewRequest(method, uri, body)
	request.Header.Add("Accept", "application/json")
	for header := range headers {
		request.Header.Add(header, headers[header])
	}

	return request
}

func handleResponse(request *http.Request) (map[string]interface{}, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	res := buf.Bytes()

	data := make(map[string]interface{})
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func Get(uri string, headers map[string]string, args map[string]interface{}) (map[string]interface{}, error) {
	uri = uri + handleUriValues(args)
	request := handleRequest(http.MethodGet, uri, headers, nil)

	res, err := handleResponse(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Post(uri string, headers map[string]string, args map[string]interface{}) (map[string]interface{}, error) {
	marshal, _ := json.Marshal(args)
	request := handleRequest(http.MethodPost, uri, headers, bytes.NewReader(marshal))

	res, err := handleResponse(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Put(uri string, headers map[string]string, args map[string]interface{}) (map[string]interface{}, error) {
	marshal, _ := json.Marshal(args)
	request := handleRequest(http.MethodPut, uri, headers, bytes.NewReader(marshal))

	res, err := handleResponse(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
