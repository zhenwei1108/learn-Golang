package _interface

type Req struct {
	A int
	B int
}

type Resp struct {
	Sum int
}
type demoInterface interface {
	Sum(req2 *Req) (*Resp, error)
}
