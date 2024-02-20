package main

import (
	"learn-goLang/src/dir"
	_interface "learn-goLang/src/interface"
	"log"
)

func main() {
	sumInt := _interface.SumInt{}
	req := _interface.Req{A: 3, B: 5}
	sum, _ := sumInt.Sum(&req)
	log.Print("the sum = ", sum.Sum)
	log.Println("one of data is : ", sumInt.One)

	//获取当前用户信息
	dir.GetUserMsg()

}
