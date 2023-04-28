package mine

import (
	"fmt"
	"testing"
	"time"
)

// 测试入口
func TestGetTime(t *testing.T) {
	//调用 calculate_time 传入另外一个函数 doJob, 返回一个函数
	calculateTime := calculate_time(doJob)
	//调用返回的函数
	s := calculateTime(100)
	t.Log(s)
}

// 函数 calculate_time
//
//		方法入参 是一个函数，此函数的入参是字符串，出参是int
//	 方法出参 是一个函数，此函数的入参是int， 出参是字符串
func calculate_time(inner func(s string) int) func(i int) string {
	return func(i int) string {
		start := time.Now()
		result := inner(string(i))
		fmt.Print(result)
		end := time.Now().Nanosecond() - start.Nanosecond()
		return string(rune(end))
	}
}

// 构造 用于输入的inner方法
func doJob(start string) int {
	fmt.Print("input: " + start)
	return 0
}
