package mine

import "testing"

const (
	one = iota + 1
	two
	three
)
const (
	first = 1 << iota
	second
	third
	forth
)

const (
	a = 1
	b
	c
	d
	e = iota
)

func Test_link_int(t *testing.T) {
	t.Log(e)
	t.Log(forth)

}
