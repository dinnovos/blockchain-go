package blockchain

import (
	"sync"
	"fmt"
)

type Blockchain struct {
    Blocks []*Block
}

var mutex = &sync.Mutex{}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{GenesisBlock()}}
}

func (this *Blockchain) AddBlock(data string) {

	previousBlock := this.Blocks[len(this.Blocks)-1]

	newBlock := Mine(previousBlock, data)

	this.Blocks = append(this.Blocks, newBlock)
}

func (this *Blockchain) ReplaceBlocks(newBlocks []*Block) {
	mutex.Lock()

	// Si la cadena de bloque es mayor a la que posee la instancia actual
	if len(newBlocks) > len(this.Blocks) {

		// Si la cadena de bloques es valida
		if IsBlockchainValid(newBlocks){
			this.Blocks = newBlocks
		}	
	}

	mutex.Unlock()
}

func (this *Blockchain) PrintBlocks() {

	for _, block := range this.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Printf("Prev: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println();
	}
}

func IsBlockchainValid(blocks []*Block) bool{

	genesisBlockOriginal := GenesisBlock();
	genesisBlock := blocks[0]

	// Compara que el genesis de la blockchain sea igual al original
	if genesisBlock.ToString() != genesisBlockOriginal.ToString(){
		return false
	}

	for index, block := range blocks{
		if index > 0{

			// Verifica que el PreviousHash del bloque actual sea igual al hash del bloque anterior
			if block.PreviousHash != blocks[index-1].Hash{
				return false
			}

			if CalculateHash(block.Index, block.Timestamp, block.PreviousHash, block.Data) !=  block.Hash{
				return false
			}
 
		}
	}

	return true
}
