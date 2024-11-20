package memorySupportCalculator

type MemoryCalc struct {
	memory float64
}

func (c *MemoryCalc) SetMemory(val float64) {
	c.memory = val
}
func (c *MemoryCalc) AddToMemory(val float64) float64 {
	return c.memory + val
}
func (c *MemoryCalc) SubtractFromMemory(val float64) float64 {
	return c.memory - val
}
func (c *MemoryCalc) ResetMemory() {
	c.memory = 0
}
