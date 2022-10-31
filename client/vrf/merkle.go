package vrf

import (
	"ShareChain-Client/block"
	"ShareChain-Client/parameter"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/merkletree"
)

func VerifyData(response parameter.GetResponsePayload) bool {
	// 1. 由数据组装Tx
	txSharedData := parameter.TxSharingData{
		DataType:  response.DataType,
		Key:       response.Key,
		Value:     response.Value,
		TimeValid: response.TimeValid,
	}
	// 2. 由TX所在的区块头中获取 Merkle 根
	merkleRoot := block.GetMerkleRoot(response.BlockHash)
	if merkleRoot == nil {
		return false
	}
	// 3. 验证 TX 与 Merkle 路径中的 Index = 0 元素是否一致
	bTx, _ := json.Marshal(txSharedData)
	fmt.Println(string(response.MerklePath.ProofSet[0]))
	fmt.Println(string(bTx))

	//// todo 删除DEBUG代码
	//bhash_proof,_ := utils.GetHashFromBytes(response.MerklePath.ProofSet[0])
	//bhash_bTx,_ := utils.GetHashFromBytes(bTx)
	//hash_proof := hex.EncodeToString(bhash_proof)
	//hash_bTx := hex.EncodeToString(bhash_bTx)
	//fmt.Println("hash_proof : ", hash_proof)
	//fmt.Println("hash_bTx : ", hash_bTx)
	//
	//// todo 删除DEBUG代码

	if bytes.Equal(response.MerklePath.ProofSet[0], bTx) {
		fmt.Println("交易存在于验证路径中，开始进入验证过程")
	} else {
		return false
	}
	// 4. 提取 Response 中的验证数据
	path := response.MerklePath.ProofSet
	proofIndex := response.MerklePath.ProofIndex
	numLeaves := response.MerklePath.NumLeaves
	// 5. 执行 Merkle 验证
	isVrf := merkletree.VerifyProof(sha256.New(), merkleRoot, path, proofIndex, numLeaves)
	// 6. 返回结果
	if !isVrf {
		fmt.Println("验证失败!")
		return false
	}
	return true
}