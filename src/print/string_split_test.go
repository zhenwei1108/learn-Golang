package mine

import (
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
