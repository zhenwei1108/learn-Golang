package mine

import (
	"fmt"
	"testing"
)

// 测试指针
func Test_pointerOfGo(t *testing.T) {
	s := "good bye"
	//p为s的指针。
	var p = &s
	//p的类型为指针，*p代表取值
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	i := 12
	//& 代表取指针
	j := &i
	//* 代表取指针对应的值
	fmt.Printf("real i:%d", *j)

}
