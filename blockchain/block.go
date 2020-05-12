package blockchain

import (
	"time" 
	"sync"
)

type Block struct {
	Timestamp 		time.Time
	PreviousHash	string
    Hash			string
    Data 			string
}

var instance *Block
var once sync.Once

func (this *Block) Genesis() *Block{
   instance = &Block{
    	Timestamp: time.Now(),
    	PreviousHash: "empty",
    	Hash: "g3n3s1s-h4sh",
    	Data: "g3n3s1s-data",
    }

    return instance
}