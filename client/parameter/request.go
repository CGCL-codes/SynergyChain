package parameter


// 请求共享数据结构体
type GetSharedDataRequest struct {
	MsgType 	int 					`json:"msg_type"`
	From    	string  				`json:"from"`
	Payload 	GetSharedDataPayload 	`json:"payload"`
	TimeStamp 	int64					`json:"time_stamp"`
}
type GetSharedDataPayload struct {
	DataType	string	`json:"data_type"`
	Key			string	`json:"key"`
	Signature 	string	`json:"signature"`
}