package mine

import "testing"

var array []int

func Test_array(t *testing.T) {
	array[1] = 1
	array_b := [2]int{1}
	t.Log(array_b)
	array_a := [...][3]int{{1, 2, 3}, {1}}
	t.Log(array_a)

	//foreach 遍历
	for _, each := range array_a {
		t.Log(each)
	}

}

// 截取数组
func Test_split_array(t *testing.T) {
	//向数组填充
	array := append(array, 1, 2, 3, 4)
	tail := array[1:]
	t.Log(tail)

	t.Log(array[2:len(array)])
	//含1不含2
	t.Log(array[1:2])

	array = make([]int, 3, 5)
	t.Log(len(array), cap(array))
}
