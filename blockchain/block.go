package blockchain

import "time"

type Block struct {
	Timestamp 		time.Time
    Hash			string
    PreviousHash	string
    Data 			string
}

func (this *Block) SetValues(data string, hash string, previousHash string, timestamp time.Time){
	this.Timestamp 		= timestamp
	this.Hash 			= hash
	this.PreviousHash 	= previousHash
	this.Data 			= data
}