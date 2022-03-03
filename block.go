package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 简化的区块链数据结构

type Block struct {
	Version       string         // 版本号
	Timestamp     int64          // 时间戳
	Transactions  []*Transaction // 区块存储的实际有效信息，也就是交易
	PrevBlockHash []byte         // 前一个块的哈希
	Hash          []byte         // 当前块的哈希
	Nonce         int            // 随机值
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.HashTransactions(), timestamp}, []byte{})
	hash := sha256.Sum256(headers) // 计算哈希

	b.Hash = hash[:]
}

// 新建一个块

//func NewBlock(version string, data string, prevBlockHash []byte, nonce int) *Block {
//	block := &Block{version, time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, nonce}
//	block.SetHash()
//	return block
//}

func NewBlock(version string, transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{version, time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
