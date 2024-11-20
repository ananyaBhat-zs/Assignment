package calculator

type Calc struct {
	n1, n2 float64
}

func NewCalc(v1, v2 float64) *Calc {
	return &Calc{
		n1: v1,
		n2: v2,
	}
}
func (c *Calc) Add() float64 {
	return c.n1 + c.n2
}
func (c *Calc) Sub() float64 {
	return c.n1 - c.n2
}
func (c *Calc) Mul() float64 {
	return c.n1 * c.n2
}
func (c *Calc) Div() (float64, bool) {
	if c.n2 != 0 {
		return c.n1 / c.n2, true
	}
	return 0.0, false
}
