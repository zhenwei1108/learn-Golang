package extends

import "testing"

type base struct {
	traceID string
	appID   string
}

type People struct {
	base
	name string
	age  int
}

type Habit struct {
	play string
}

type Bob struct {
	//匿名属性等于继承
	People
	Habit
}

func Test_game(t *testing.T) {
	b := new(Bob)
	b.name = "bob"
	b.play = "game"
	b.age = 18
	b.appID = "app_123"
	b.traceID = "123456"
	t.Log(&b)
	t.Log(b)
}
