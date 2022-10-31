package client

import (
	"ShareChain-Client/parameter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetBlockHead(localHeight int64) (Height int64, blockHead map[string]parameter.BlockHead){
	// 组装 URL
	sLocalHeight := strconv.FormatInt(localHeight, 10)
	url := "http://127.0.0.1:7999/SyncBlockHead/" + sLocalHeight
	
	// 新建客户端请求
	client := http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	
	// 执行请求
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Println("同步区块头请求出错", err)
		return -1, nil
	}
	responseBody, _ := ioutil.ReadAll(response.Body)

	var blockHeadList parameter.BlockHeadInfo
	json.Unmarshal(responseBody, &blockHeadList)

	return blockHeadList.EndHeight, blockHeadList.BlockHeadList
}