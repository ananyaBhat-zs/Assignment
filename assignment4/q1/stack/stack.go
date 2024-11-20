// This is stack implementation
package stack

type Stack struct {
	stk []int
	top int
}

func NewStack() *Stack {
	return &Stack{
		stk: make([]int, 0),
		top: -1,
	}
}

func (s *Stack) Push(v int) {
	s.stk = append(s.stk, v)
	s.top += 1
}
func (s *Stack) Pop() (bool, int) {
	if s.top != -1 {
		val := s.stk[s.top]
		s.stk = s.stk[:s.top]
		s.top--
		return true, val
	}
	return false, -1
}
