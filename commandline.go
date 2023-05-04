package main

import "fmt"

func (cli *CLI) addBlock(data string) {
	fmt.Printf("添加区块被调用")
		bc , _:= GetBlockChainInstance()
		 
		err := bc.AddBlock(data)
		if err != nil{
			fmt.Println("AddBlock failed", err)
			return 
		}
		fmt.Println("添加区块成功")

}

func (cli *CLI) CreatBlockChain() {
	err := CreatBlockChain()
	if err != nil {
		fmt.Println("Create Blockchain failed", err)
		return
	}
	fmt.Printf("创建区块链成功")

}

func (cli *CLI) print() {
	bc , _:= GetBlockChainInstance()
	//调用迭代器，输出blockchain
	it := bc.NewIterator()
	for {
		//调用next方法，获取区块，游标左移
		block := it.Next()

		fmt.Printf("\n+++++++++++++++++++++++++\n")
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("前哈希：%x\n", block.PreHash)
		fmt.Printf("MerkleRoot : %d\n", block.MerkleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Bits : %d\n", block.Bits)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("当前节点哈希:%x\n", block.Hash)
		fmt.Printf("当前节点数据:%s\n", string(block.Data))

		pow := NewProofOfWok(block)

		fmt.Printf("有效ma: %v\n", pow.isValid())

		//退出条件
		if block.PreHash == nil {
			fmt.Println("区块链遍历结束")
			break
		}

	}

}