package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//实现挖矿功能：pow

//字段
//--区块：block
//--目标值：target
//方法：
//--run计算
//--功能：找到nonce，满足哈希值<target

type ProofOfWork struct{

	block *Block

	target *big.Int //结构，提供了方法：比较，把哈希值设为big.Int类型

}

//创建proofofwork
//block由用户提供
//target目标值由系统提供
func NewProofOfWok(block *Block) *ProofOfWork{
	pow := ProofOfWork{
		block: block,
	}

	//难度值先写死,后面再补充
	targetStr := "0100000000000000000000000000000000000000000000000000000000000000"

	tmpBigInt := new(big.Int)
	//将难度值赋值给Big.Int
	tmpBigInt.SetString(targetStr,16)

	pow.target = tmpBigInt

	return &pow
}
//写工作量证明，挖矿函数，不断变nonce，使得sha256（数据+nonce）《难度值，这个方法是属于pow的
func (pow *ProofOfWork)Run()([]byte,uint64){
	var nonce uint64
	var hash [32]byte

	fmt.Println("开始挖矿。。。")

	for{
		fmt.Printf("%x\r", hash[:])
		
		//1.拼接字符踹+nonce
		data := pow.PrepareData(nonce)
		//2.哈希值=sha256（字符串+nonce）
		//这个hash是一个32位的数组
		hash = sha256.Sum256(data)


		//将hash转换为bigint类型，便于比较
		tmpInt := new(big.Int)
		tmpInt.SetBytes(hash[:])

		//3.比较当前哈希与难度值
		//if 当前哈希《难度值{
			//break
		//}
		//当前计算的哈希值。CMP（难度值）
		if tmpInt.Cmp(pow.target) == -1{
			fmt.Printf("挖矿成功,hash:%x, nonce:%d\n", hash[:],nonce)
			break
		}

		nonce ++
	}

	//return 哈希， nonce
	return hash[:],nonce

}
//拼接nonce和block数据
func (pow *ProofOfWork)PrepareData(nonce uint64)[]byte{

	b :=pow.block

	tmp := [][]byte{
		uintToByte(b.Version),
		b.PreHash,
		b.MerkleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(nonce),
		//b.Hash,
		b.Data,
	}
	//使用join方法将二维切片数组转为1维切片
	data := bytes.Join(tmp, []byte{})

	return data

}

func (pow *ProofOfWork) isValid() bool{
	//1.获取区块
	//2.拼接数据
	data := pow.PrepareData(pow.block.Nonce)
	//3.计算sha256
	hash := sha256.Sum256(data)
	//4.与难度值比较
	tmpInt := new(big.Int)
	tmpInt.SetBytes(hash[:])


	return tmpInt.Cmp(pow.target) == -1
}