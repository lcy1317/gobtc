package main

import (
	bolt "bbolt" // 这是boltDB
	"colorout"
	"fmt"
)

const dbFile = "blockchain.db"
const blocksBucket = "luochengyu"
const genesisData = "回忆这东西就像做倒行马车，远处的风景看的很清晰，身旁的风景却很模糊。等我们过了很多年，再看现在的时光，记忆力留下的也都是美好的 --梦逸云"

type Blockchain struct {
	//tip意为"末梢", 这里记录链中最新一个区块的hash
	tip []byte
	db  *bolt.DB
}

// 创建一个带有创始区块的区块链。
//func NewBlockchain() *Blockchain {
//	return &Blockchain{[]*Block{NewGenesisBlock()}}
//}
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		fmt.Println(colorout.Red("区块链数据读取错误！"))
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				fmt.Println(colorout.Red("创建存储错误！"))
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func NewGenesisBlock() *Block {
	return NewBlock("0.0-beta0", genesisData, []byte{})
}

//TODO: Version 处理
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		fmt.Println(colorout.Red("区块链数据读取错误！"))
	}

	newBlock := NewBlock("0.0-beta1", data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			fmt.Println(colorout.Red("区块链数据序列化错误！"))
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}
