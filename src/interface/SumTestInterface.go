package _interface

// 首字母大写，表示 public  都可以访问
type Req struct {
	A int
	B int
}

type Resp struct {
	Sum int
}

// 先定义一个接口，作为接收者。 然后实现指定接口的方法即可
type demoInterface interface {
	Sum(req2 *Req) (*Resp, error)
}
