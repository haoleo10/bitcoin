package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//这种工具类的函数都可以放在utils里面

func uintToByte(num uint64) []byte{
	var buffer bytes.Buffer
	//使用二进制编码
	//num以小段对齐的方式 写给buffer
	err := binary.Write(&buffer,binary.LittleEndian, &num)
	if err != nil{
		fmt.Println("binary.Write err:", err)
		return nil
	}
	return buffer.Bytes()
}