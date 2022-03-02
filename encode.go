package main

import (
	"bytes"
	"colorout"
	"encoding/gob"
	"fmt"
)

// 序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		fmt.Println(colorout.Red("序列化错误！"))
	}
	return result.Bytes()
}

// 反序列化成block
func DeserializeBlock(d []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println(colorout.Red("反序列化错误！"))
	}
	return &block
}
