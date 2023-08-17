package main

import (
	_interface "learn-goLang/src/interface"
	"log"
)

func main() {
	sumInt := _interface.SumInt{}
	req := _interface.Req{A: 3, B: 5}
	sum, _ := sumInt.Sum(&req)
	log.Print("the sum = ", sum.Sum)
	log.Println("one of data is : ", sumInt.One)
}
