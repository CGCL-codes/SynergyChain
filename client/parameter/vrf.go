package parameter

// Merkle 验证路径
type MerklePath struct {
	// ProofSet 就是验证用的路径 其中ProofSet[0] 就是需要验证的数据的 []byte 形式
	ProofSet  	[][]byte	`json:"proof_set"`
	ProofIndex 	uint64		`json:"proof_index"`
	NumLeaves	uint64		`json:"num_leaves"`
}
