package client

import (
	"ShareChain-Client/parameter"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	JOIN = iota
	SUBMIT
	GET_BLOCKHEAD
	GET_DATA
)

// 请求共享数据
func GetSharedData(url string, key string, address string) *parameter.GetResponsePayload {
	// 1. 组装请求结构体
	// todo 概念验证实验，省略签名
	requestBody := parameter.GetSharedDataRequest{
		MsgType:   GET_DATA,
		From:      address,
		TimeStamp: time.Now().UTC().UnixNano(),
	}
	payload := parameter.GetSharedDataPayload{
		DataType:  "",
		Key:       key,
		Signature: "",
	}
	requestBody.Payload = payload

	// 2. 序列化请求结构体
	bRequestBody, _ := json.Marshal(requestBody)

	// 3. 建立客户端
	client := &http.Client{}
	req := bytes.NewBuffer(bRequestBody)
	request, _ := http.NewRequest("GET", url, req)
	request.Header.Set("Content-type", "application/json")

	// 4. 执行请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("请求出错", err)
		return nil
	}
	respBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(respBody))
	var getResponse parameter.GetResponse
	err = json.Unmarshal(respBody, &getResponse)
	if err != nil {
		fmt.Println(err)
	}
	return &getResponse.Data
}

