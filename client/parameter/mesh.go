package parameter

type TxSharingData struct {
	// 打包进交易中的共享数据
	DataType    string
	Key   		string
	Value 		string
	TimeValid   string
}

//Block B
type BlockHead struct {
	Version       int64
	Height        int64
	TxCount       int
	Timestamp     int64
	Producer      string
	PrevBlockHash string
	Hash          string
	MerkleRoot	  []byte
}