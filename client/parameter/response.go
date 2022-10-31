package parameter

// 请求共享数据的返回结果
type GetResponse struct {
	Code		int					`json:"code"`
	Data		GetResponsePayload	`json:"data"`
	Timestamp	int64				`json:"timestamp"`
}
type GetResponsePayload struct {
	DataType    string 	`json:"data_type"`
	Key		  	string	`json:"key"`
	Value 	  	string	`json:"value"`
	TimeValid   string  `json:"time_valid"`
	BlockHash	string	`json:"block_hash"`
	MerklePath	MerklePath	`json:"merkle_path"`
}

// 同步区块头的数据结构
type BlockHeadInfo struct {
	// HashToHeight map[Hash]Height
	EndHeight	  int64	`json:"end_height"`
	HashToHeight  map[string]string		`json:"hash_to_height"`
	BlockHeadList map[string]BlockHead	`json:"block_head_list"`
}