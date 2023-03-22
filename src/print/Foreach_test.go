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

		if a := 1; n > 2 {
			a++
			t.Log(n)
		}

		fmt.Print(n)
	}

}
