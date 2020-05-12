package main

import (
	"fmt"
	b "./blockchain"
)

func main(){
	fmt.Println("Iniciando Blockchain")

	blockInstance := new(b.Block)

	genesisBlock := blockInstance.Genesis();
	
	fmt.Println(genesisBlock)
}