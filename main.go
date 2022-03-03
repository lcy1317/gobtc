package main

import (
	"color"
)

func main() {
	color.Yellowln("This is a go btc implement.")
	//bc := NewBlockchain(NodeID)
	//defer bc.db.Close()

	cli := CLI{}
	cli.Run()
}
