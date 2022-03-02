package main

import (
	"colorout"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(colorout.Yellow("This is a go btc implement."))
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf(colorout.Green("PoW: " + strconv.FormatBool(pow.Validate()) + "\n"))
		fmt.Println()
		fmt.Println()
	}

}
