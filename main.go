package main


// 打印
func main() {
	cli := CLI{}
	cli.Run()
}

// func main1() {

// 	err := CreatBlockChain()
// 	fmt.Println("err:", err)

// 	//获取区块链实例,相当于打开数据库了，返回一个bc
// 	//bc里面已经存储了最后一个区块的哈希
// 	bc, err := GetBlockChainInstance()
// 	defer bc.db.Close()
// 	if err != nil {
// 		fmt.Println("GetBlockChainInstance err:", err)
// 		return
// 	}
// 	err = bc.AddBlock("hello world !!!!!!!")
// 	if err != nil {
// 		fmt.Println("addBlock err:", err)
// 		return

// 	}
// 	err = bc.AddBlock("hello hao !!!!!!!")
// 	if err != nil {
// 		fmt.Println("addBlock err:", err)
// 		return

// 	}
// 	//调用迭代器，输出blockchain
// 	it := bc.NewIterator()
// 	for {
// 		//调用next方法，获取区块，游标左移
// 		block := it.Next()

// 		fmt.Printf("\n+++++++++++++++++++++++++\n")
// 		fmt.Printf("Version : %d\n", block.Version)
// 		fmt.Printf("前哈希：%x\n", block.PreHash)
// 		fmt.Printf("MerkleRoot : %d\n", block.MerkleRoot)
// 		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
// 		fmt.Printf("Bits : %d\n", block.Bits)
// 		fmt.Printf("Nonce : %d\n", block.Nonce)
// 		fmt.Printf("当前节点哈希:%x\n", block.Hash)
// 		fmt.Printf("当前节点数据:%s\n", string(block.Data))

// 		pow := NewProofOfWok(block)

// 		fmt.Printf("有效ma: %v\n", pow.isValid())

// 		//退出条件
// 		if block.PreHash == nil {
// 			fmt.Println("区块链遍历结束")
// 			break
// 		}

// 	}

	//newbc := NewBlockChain()
	// time.Sleep(1 * time.Second)
	// newbc.AddBlock("区块1号")
	// time.Sleep(1 * time.Second)
	// newbc.AddBlock("区块2号")
	//time.Sleep(1 * time.Second)
	//遍历区块
	// for i, block := range newbc.Blocks{
	// 	fmt.Printf("\n++++++当前区块：%d +++++++++\n",i)
	// 	fmt.Printf("Version : %d\n", block.Version)
	// 	fmt.Printf("前哈希：%x\n",block.PreHash)
	// 	fmt.Printf("MerkleRoot : %d\n", block.MerkleRoot)
	// 	fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
	// 	fmt.Printf("Bits : %d\n", block.Bits)
	// 	fmt.Printf("Nonce : %d\n", block.Nonce)
	// 	fmt.Printf("当前节点哈希:%x\n",block.Hash)
	// 	fmt.Printf("当前节点数据:%s\n",string(block.Data))

	// 	pow := NewProofOfWok(block)

	// 	fmt.Printf("有效ma: %v\n", pow.isValid())
	//}

//}
