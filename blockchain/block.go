package blockchain

import (
	"time"
	"encoding/hex"
	"crypto/sha256"
	"encoding/json"
	"strconv"
)

type Block struct {
	Index			uint64
	Timestamp		int64
	PreviousHash	string
    Hash 			string
    Data 			string
}

func (this *Block) ToString() string {
	str, _ := json.Marshal(this)
	return string(str)
}

func NewBlock(index uint64, timestamp int64, previousHash string, hash string, data string) *Block {
	newBlock := &Block{Index: index, Timestamp: timestamp, PreviousHash: previousHash, Hash: hash, Data: data}
    return newBlock
}

func GenesisBlock() *Block{
    return NewBlock(0, time.Date(2000, 01, 01, 01, 01, 01, 01, time.UTC).Unix(), "3mp7y", "g3n3s1s-h4sh", "g3n3s1s-data")
}

func Mine(previousBlock *Block, data string) *Block{
	timestamp := time.Now().Unix()
	previousHash := previousBlock.Hash

	hash := CalculateHash(previousBlock.Index + 1, timestamp, previousHash, data)

	return NewBlock(previousBlock.Index + 1, timestamp, previousHash, hash, data)
}

func CalculateHash(index uint64, timestamp int64, previousHash string, data string) string {
	record := strconv.FormatUint(index, 10) + strconv.FormatInt(timestamp, 10) + previousHash + string(data)

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}