package main

import (
	"colorout"
	"fmt"
)

func main() {
	fmt.Println(colorout.Yellow("This is a go btc implement."))
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
