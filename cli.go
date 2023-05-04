package main

import (
	"fmt"
	"os"
)

//处理用户输入的命令，完成具体函数的调用
//cli :command line命令行

type CLI struct{
	//不需要字段

}
//使用说明，帮助用户正确使用

const Usage = `
Usage:
	./blockchain-v1 create "创建区块链"
	./blockchain-v1 addBlock <需要写入的数据> "添加区块"
	./blockchain-v1 print "打印区块链"

`

//负责解析命令的方法
func (cli *CLI) Run(){
	cmds := os.Args

	//用户至少输入两个参数
	if len(cmds) < 2{
		fmt.Println("输入参数无效")
		fmt.Printf(Usage)
		return
	}

	switch cmds[1]{
	case "create":
		fmt.Printf("创建区块被调用")
		cli.CreatBlockChain()
		
	case "addBlock":
		if len(cmds) != 3{
			fmt.Printf("输入参数无效")
			return
		}
		data := cmds[2]//需要校验个数
		cli.addBlock(data)
		
	case "print":
		fmt.Printf("打印区块被调用")
		cli.print()
	
	default:
		fmt.Printf("输入参数无效")
		fmt.Printf(Usage)
	}	

}