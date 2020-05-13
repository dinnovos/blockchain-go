package main

import (
	"fmt"
	bc "./blockchain"
)

func main(){
	fmt.Println("\n===================== Iniciando Blockchain =====================\n")

	blockchain := bc.NewBlockchain();
	blockchain.AddBlock("data 1");
	blockchain.AddBlock("data 2");
	blockchain.AddBlock("data 3");

	fmt.Println("> Blockchain 1:")
	blockchain.PrintBlocks()

	if bc.IsBlockchainValid(blockchain.Blocks) {
		fmt.Println("Blockchain is valid...");
	}else{
		fmt.Println("Blockchain is invalid...");
	}

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println()

	fmt.Println("> Blockchain 2:")
	blockchain2 := bc.NewBlockchain();
	blockchain2.ReplaceBlocks(blockchain.Blocks)
	blockchain2.PrintBlocks()

	//blockInstance := new(b.Block)

	/*
	genesisBlock := b.Genesis();
	newBlock1 := b.Mine(genesisBlock, "data 1");
	newBlock2 := b.Mine(newBlock1, "data 2");
	newBlock3 := b.Mine(newBlock2, "data 3");

	fmt.Println(genesisBlock.ToString())
	fmt.Println(newBlock1.ToString())
	fmt.Println(newBlock2.ToString())
	fmt.Println(newBlock3.ToString())
	*/


}