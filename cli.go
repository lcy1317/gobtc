package main

import (
	"color"
	"colorout"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// 创建交互接口

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println(colorout.Cyan("Usage:"))
	fmt.Println(colorout.Cyan("TODO:createWallet 创建钱包"))                                  //TODO:
	fmt.Println(colorout.Cyan("TODO:listAddresses 显示所有地址"))                               //TODO:
	fmt.Println(colorout.Cyan("TODO:getBalance -address 根据地址查询金额"))                       //TODO:
	fmt.Println(colorout.Cyan("TODO:createBlockChain 根据地址创建区块链"))                         //TODO:
	fmt.Println(colorout.Cyan("TODO:send -from FROM_ADDR -to TO_ADDR -amount AMOUNT 转账")) //TODO:
	fmt.Println(colorout.Cyan("addBlock 向区块链增加区块"))
	fmt.Println(colorout.Cyan("showBlockChain 显示区块链"))
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
func (cli *CLI) Run() {
	cli.validateArgs()
	//nodeID := os.Getenv("NODE_ID")
	nodeID := NodeID
	if nodeID == "" {
		fmt.Printf("NODE_ID env. var is not set!")
		os.Exit(1)
	}
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	createBlockChainData := createBlockChainCmd.String("address", "", "The address to send genesis block reward to")

	switch os.Args[1] {
	case "showBlockChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(colorout.Red("输出链信息参数读取错误！"))
		}
	case "createBlockChain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(colorout.Red("创建链参数读取错误！"))
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createBlockChainCmd.Parsed() {
		if *createBlockChainData == "" {
			createBlockChainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockChain(*createBlockChainData, nodeID)
	}
}

// TODO: NodeID
func (cli *CLI) createBlockChain(address string, nodeID string) {
	bc := CreateBlockChain(address, nodeID)
	defer bc.db.Close()
}

func (cli *CLI) printChain() {
	bc := NewBlockchain(NodeID)
	defer bc.db.Close()
	bci := bc.Iterator()
	color.Blueln("Show the Block Chain!")

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Transaction: %s\n", block.HashTransactions())
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
