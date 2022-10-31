package block

import (
	"ShareChain-Client/client"
	"ShareChain-Client/parameter"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var Local_BlockChain_Height int64

var BlockChainDB *leveldb.DB



func InitBlockChainDB() {
	Local_BlockChain_Height = 0
	BlockChainDB,_ = leveldb.OpenFile("./BlockChain", nil)
}


func SaveBlockHead(blockHead parameter.BlockHead) {
	bBlockHead, _ := json.Marshal(blockHead)
	fmt.Println("存储区块头 Hash : ", blockHead.Hash)
	err := BlockChainDB.Put([]byte(blockHead.Hash), bBlockHead, nil)
	if err !=  nil {
		fmt.Println("未能正确存储区块头信息: ", err)
	}
}

func GetMerkleRoot(hash string) []byte {
	bBlockHead, err := BlockChainDB.Get([]byte(hash), nil)
	if err != nil {
		fmt.Println("通过 Hash 访问区块头失败：", err)
		return nil
	}
	var blockHead parameter.BlockHead
	json.Unmarshal(bBlockHead, &blockHead)
	return blockHead.MerkleRoot
}

func SyncBlockHead() {
	newHeight, blockHeadList := client.GetBlockHead(Local_BlockChain_Height)

	if newHeight == -1 {
		fmt.Println("同步失败，请求数据没有得到正确回应")
		return
	}

	Local_BlockChain_Height = newHeight
	fmt.Println("Local_BlockChain_Height = ", Local_BlockChain_Height)
	for _,v := range blockHeadList {
		bBlockHead,_ := json.MarshalIndent(v,""," ")
		fmt.Println(string(bBlockHead))
		SaveBlockHead(v)
	}

}