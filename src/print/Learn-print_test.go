package mine_test

import (
	"fmt"
	mine "learn-goLang/src/print"
	"testing"
)

// 文件名必须以 "_test" 结尾， 方法名，必须为 "Test" 开头

func Test_print(t *testing.T) {
	t.Log("this is mine test in Learn-print_test")
	mine.TestGetTime(t)
}

var (
	a int = 1
	b int = 1
)

func Test_Fib(t *testing.T) {
	c := 1
	print(c)
	var a = 1
	var b = 1
	fmt.Print(" ", a)
	fmt.Print(" ", b)
	for i := 0; i < 10; i++ {
		tmp := a + b
		a = b
		b = tmp
		print(" ", tmp)
	}

}
