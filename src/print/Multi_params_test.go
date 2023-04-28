package mine

import (
	"testing"
)

// 方法返回多个返回值， 第一个为入参。 第二个为返回值的类型
func double_return(i int) (int, int) {
	return i + 1, i + 2
}

func Test_main(t *testing.T) {
	first, second := double_return(1)

	one, _ := double_return(1)
	t.Log(first, second)
	t.Log("忽略另外一个：", one)
}
