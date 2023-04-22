package mine

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_split(t *testing.T) {
	splitString := "q,w,e,r,t"
	//拆分
	allString := strings.Split(splitString, ",")

	for index, data := range allString {
		t.Log(index, data)
	}
	//拼接
	join := strings.Join(allString, "")
	t.Log(join)
}

func Test_type_conv(t *testing.T) {
	one := "1"
	t.Log(one)
	//获取实际类型
	t.Log(reflect.TypeOf(one))
	//类型转换  int -> string
	itoa := strconv.Itoa(10)
	t.Log(reflect.TypeOf(itoa))
	//类型转化 string -> int
	atoi, _ := strconv.Atoi(itoa)
	t.Log(reflect.TypeOf(atoi))
	t.Log(atoi + 1)
}
