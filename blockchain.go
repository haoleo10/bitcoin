package main

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

//定义区块链结构（使用数组模拟区块链）
type BlockChain struct {
	//Blocks []*Block
	db *bolt.DB //用于存储数据

	tail [] byte //最后一个区块的哈希值
}
const lastBlockHashKey = "lastBlockHashKey"//用于访问bolt数据库，得到最后一个区块的哈希值
const bucketBlock = "buckeetBlock"//装block的桶
const genInfo = "The first block"
const blockchainDBFile = "blockchain.db"
//创建区块链，从无到有，这个函数仅执行一次
func CreatBlockChain() error {

	//区块链不存在，则创建
	db, err := bolt.Open(blockchainDBFile,0600,nil)
	if err != nil{
		return err
	}
	//不要db。close，后需要使用这个句柄
	defer db.Close()

	//--2.开始创建
	err = db.Update(func (tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte(bucketBlock))
		//如果bucket为空，说明不存在
		if bucket == nil{
			//创建bucket
			bucket, err := tx.CreateBucket([]byte(bucketBlock))
			if err != nil{
				return err
			}
			//写入创世块
			genBlock := NewBlock(genInfo, nil)
			//key是区块的哈希值，value是block的字节流，将block序列化
			bucket.Put(genBlock.Hash, genBlock.Serialize())
			//更新最后一个区块的哈希值到数据库
			bucket.Put([]byte(lastBlockHashKey),genBlock.Hash)

		}
		return nil
	})
	return err

}
//获取区块链实例，用于后续操作，每次有业务时都会调用
func GetBlockChainInstance() (*BlockChain, error) {

	var lastHash [] byte //内存中最后一个区块的哈希值
	//该函数两个功能：
	//--1.如果区块链不存在，则创建，同事返回blockchain的示例
	db, err := bolt.Open(blockchainDBFile,0400,nil)
	if err != nil{
		return nil, err
	}
	//不要db。close，后需要使用这个句柄
	//--2.如果区块链存在，则直接返回blockchain示例
	db.View(func (tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte(bucketBlock))
		//如果bucket为空，说明不存在
		if bucket == nil{
			return errors.New("bucket不应为空")
		}else{
			//直接读取特定的key，得到最后一个区块的哈希值
			lastHash = bucket.Get([]byte(lastBlockHashKey))
		}
		return nil
	})

	bc := BlockChain{db, lastHash}
	return &bc, nil
}

//向区块链中添加区块

func (bc *BlockChain)AddBlock(data string)error{
	//区块链中最后一个hash
	lastblockhash := bc.tail
	//1.创建区块
	newblock := NewBlock(data,lastblockhash)

	//2/写入数据库
	err := bc.db.Update(func(tx *bolt.Tx)error{
		bucket := tx.Bucket([]byte(bucketBlock))
		if bucket == nil{
			return errors.New("addblock时bucket不应为空")
		}
		//key是新区块的哈希值，value是这个区块的字节流
		bucket.Put(newblock.Hash, newblock.Serialize())
		bucket.Put([]byte(lastBlockHashKey),newblock.Hash)

		//更新bc的tail，这样后续的addblock才会基于我们newBlock追加
		bc.tail = newblock.Hash
		return nil
	})
	return err
	//3/更新lastBlockhashkey







	// //通过下标得到最后一个区块
	// lastBlock := bc.Blocks[len(bc.Blocks) - 1]
	// //最后一个区块的哈希值是新区块的前哈希
	// prehash := lastBlock.Hash 
	// //创建block（需要数据
	// newBlock := NewBlock(data, prehash)

	// //添加bc中
	// bc.Blocks = append(bc.Blocks, newBlock)

}

//+++++++++++++++++++++++迭代器
//定义迭代器
type Iterator struct{
	db *bolt.DB
	currentHash []byte//不断移动的哈希值，遍历所有区块

}

//将迭代器绑定到blockchain
func (bc *BlockChain) NewIterator()Iterator{
	it := Iterator{
		db:bc.db,
		currentHash: bc.tail,
	}
	return it
}
//给Iterator 绑定一个方法：Next
//``1 返回当前指向的区块
//  2.向左移动
func (it *Iterator) Next()(block *Block){
	//1. 读取bucket当前哈希的数据
	err := it.db.View(func (tx *bolt.Tx) error{
		
		//读取bucket
		bucket := tx.Bucket([]byte(bucketBlock))
		if bucket == nil{
			return errors.New("Iterator next时bucket不应为nil")
		}
		//得到一个block的字节流
		blockTmpInfo := bucket.Get(it.currentHash)
		block = Deseriaalize(blockTmpInfo)
		//游标左移
		it.currentHash = block.PreHash
		return nil
	})
	if err != nil{
		fmt.Println("iterator err:",err)
		return nil
	}
	return 
}

