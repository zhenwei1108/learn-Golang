package _interface

type SumInt struct {
	One int
}

func (sumInt *SumInt) Sum(req2 *Req) (*Resp, error) {
	resp := new(Resp)
	resp.Sum = req2.A + req2.B
	sumInt.One = req2.A
	return resp, nil
}
