package blockchain

import (
	"time"
	"encoding/hex"
	"crypto/sha256"
	"encoding/json"
)

type Block struct {
	Timestamp		string
	PreviousHash	string
    Hash 			string
    Data 			string
}

func (this *Block) ToString() string {
	str, _ := json.Marshal(this)
	return string(str)
}

func Genesis() *Block{
    return GenerateBlock(time.Now(), "3mp7y", "g3n3s1s-h4sh", "g3n3s1s-data")
}

func GenerateBlock(timestamp time.Time, previousHash string, hash string, data string) *Block {

	newBlock := &Block{Timestamp: timestamp.String(), PreviousHash: previousHash, Hash: hash, Data: data}

    return newBlock
}

func Mine(previousBlock *Block, data string) *Block{
	timestamp := time.Now()
	previousHash := previousBlock.Hash

	hash := calculateHash(timestamp, previousHash, data)

	return GenerateBlock(timestamp, previousHash, hash, data)
}

func calculateHash(timestamp time.Time, previousHash string, data string) string {
	record := timestamp.String() + previousHash + string(data)

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}