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

func Test_share(t *testing.T) {
	test_array := []int{}
	for i := 0; i < 11; i++ {
		test_array = append(test_array, i)
		t.Log(len(test_array), cap(test_array))
	}
	ints := test_array[2:8]
	//里面有6个值，但是内存空间是14(到结尾)
	t.Log("从2截取后：", len(ints), cap(ints))
	ints[3] = 0

	t.Log("test_array 5 和 ints 3 共享内存空间，指向同一个地址", test_array[5], ints[3])

}
