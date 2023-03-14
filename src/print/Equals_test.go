package mine

import "testing"

func Test_equals(t *testing.T) {
	a := [...]int{1, 2, 3}

	b := [...]int{2, 3, 4}

	c := [...]int{1, 2, 3}
	//两个长度不同的数组不能比较
	t.Log(a == b)
	t.Log(a == c)

	//&^ 按位清零
	d := 3 //0011
	t.Log(d &^ 0)
	t.Log(d &^ 1)
	t.Log(d &^ 2)
}
