package mine

import "testing"

func Test_map(t *testing.T) {
	//key = String
	//value = int
	mineMap := map[string]int{"first": 0, "second": 2}
	t.Log(mineMap)

	mineMap["first"] = 1
	t.Log(mineMap)
	//会覆盖
	mineMap["first"] = -1
	t.Log(mineMap)

	//构造指定长度的空map
	intIntMap := make(map[int]int, 3)
	t.Log(intIntMap, len(intIntMap))

}

func Test_verifyIntValueNil(t *testing.T) {
	maps := map[int]int{}
	//当值不存在时，会返回0
	t.Log(maps[1])
	maps[2] = 0
	t.Log(maps[2])

	if value, result := maps[1]; result {
		t.Log(value)
	} else {
		t.Log("null")
	}

}

func Test_verifyStringValueNil(t *testing.T) {
	maps := map[int]string{}
	//当值不存在时，会返回 ""
	t.Log(maps[1])
	maps[2] = "0"
	t.Log(maps[2])
	//获取时，有两个返回值，第一个是值，第二个是boolean，表示是否有值
	if value, result := maps[1]; result {
		t.Log(value)
	} else {
		t.Log("null")
	}
}

func Test_RangeMap(t *testing.T) {
	maps := map[string]string{"1": "2", "3": "4"}
	//遍历key
	for s := range maps {
		t.Log(s)
	}
	//遍历键值对
	for key, value := range maps {
		t.Log(key, value)
	}

}
