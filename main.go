package main

import (
	"fmt"
	b "./blockchain"
)

func main(){
	fmt.Println("\n===================== Iniciando Blockchain =====================\n")

	//blockInstance := new(b.Block)

	genesisBlock := b.Genesis();
	newBlock1 := b.Mine(genesisBlock, "data 1");
	newBlock2 := b.Mine(newBlock1, "data 2");
	newBlock3 := b.Mine(newBlock2, "data 3");

	fmt.Println(genesisBlock.ToString())
	fmt.Println(newBlock1.ToString())
	fmt.Println(newBlock2.ToString())
	fmt.Println(newBlock3.ToString())
}