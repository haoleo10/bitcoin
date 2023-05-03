package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// 定义区块结构
// 1.基础字段：前哈希
// 2.版本号，时间戳，难度值
type Block struct {
	//版本号
	Version uint64

	//前区块哈希值
	PreHash []byte

	//交易的根哈希值
	MerkleRoot []byte

	//时间戳
	TimeStamp uint64

	//难度值，系统一个数值，用于计算出一个哈希值
	Bits uint64

	//随机数，挖矿要求的数值
	Nonce uint64


	//当前区块哈希值（为了方便把当前区块哈希放入Block中）
	Hash []byte

	//数据
	Data []byte
}

// 创建以一个区块
func NewBlock(data string, preHash []byte) *Block {
	b := Block{
		Version: 0,
		PreHash: preHash,
		MerkleRoot: nil,
		TimeStamp: uint64(time.Now().Unix()),
		Bits: 0,
		Nonce: 0,
		Hash:    nil,
		Data:    []byte(data),
	}
	//计算哈希值

	//b.setHash()

	//将工作量证明集成到block中
	pow := NewProofOfWok(&b)

	hash, nonce := pow.Run()

	b.Hash = hash
	b.Nonce = nonce

	return &b

}

//绑定Serialize方法,gob编码
func (b *Block)Serialize()[]byte{
	var buffer bytes.Buffer

	//编码器
	encoder := gob.NewEncoder(&buffer)
	//编码
	err := encoder.Encode(b)
	if err != nil{
		fmt.Println("Encode err:", err)
		return nil
	}
	return buffer.Bytes()

}

//反序列化，输入字节流，返回block
func Deseriaalize(src []byte) *Block{
	var block Block
	//解码器
	decoder := gob.NewDecoder(bytes.NewReader(src))
	//解码
	err := decoder.Decode(&block)
	if err != nil{
		fmt.Println("decode err", err)
		return nil
	}
	return &block
}





// // 提供计算区块哈希值的方法
// func (b *Block) setHash() {
// 	//data是区块各个字段拼成的字节流

// 	//拼接block里的三个切片，他可以接收一个二维的切片，使用一维切片拼接
// 	//部分将uint64转换为byte切片
// 	tmp := [][]byte{
// 		uintToByte(b.Version),
// 		b.PreHash,
// 		b.MerkleRoot,
// 		uintToByte(b.TimeStamp),
// 		uintToByte(b.Bits),
// 		uintToByte(b.Nonce),
// 		b.Hash,
// 		b.Data,
// 	}
// 	//使用join方法将二维切片数组转为1维切片
// 	data := bytes.Join(tmp, []byte{})

// 	hash := sha256.Sum256(data)
// 	b.Hash = hash[:]

// }
