package mine

import (
	"fmt"
	"testing"
)

func Test_foreach1(t *testing.T) {
	n := 0
	for n < 9 {
		n++
		fmt.Print(n)
	}

}

func Test_foreach2(t *testing.T) {
	n := 0
	for {
		n++

		if a := 1; n > a {
			t.Log(n)
		}

		fmt.Print(n)
	}

}

// 默认会有break
func Test_switch(t *testing.T) {
	var a = "aaa"
	switch a {
	case "a":
		t.Log("a")
	case "aa", "aaa":
		t.Log("aa...")

	}

}
