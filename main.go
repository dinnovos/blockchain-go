package main

import (
	"fmt"
	b "./blockchain"
)

func main(){
	fmt.Println("Iniciando Blockchain")

	blockchain := new(b.Block)

	//block.Data = "Mi Data"

	//block.SetValues("Jorge mi nombre", "m1-h45h", "previus-hash", time.Now())

	fmt.Println(blockchain)
}