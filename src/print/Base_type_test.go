package mine

import (
	"fmt"
	"math"
	"testing"
)

// 自定义类型
type myInt int16

func Test_type(t *testing.T) {
	var a int = 1
	var b myInt = 1.0

	fmt.Print(myInt(a) == b)
	fmt.Print(a == int(b))
	t.Log("强制类型转换")
	t.Log(math.MaxInt64)
}

func Test_Point(t *testing.T) {
	//定义a
	a := 1
	//a的指针
	aPtr := &a
	t.Log(a, aPtr)
	a = 3
	t.Log(*aPtr)
}

func Test_String(t *testing.T) {
	//默认被初始化为空串
	var a string
	t.Log(a)
	t.Log(a == "")
}
