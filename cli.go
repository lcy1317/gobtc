package main

import (
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

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(colorout.Red("增加区块参数读取错误！"))
		}
	case "showBlockChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(colorout.Red("输出链信息参数读取错误！"))
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
