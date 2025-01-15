package stack

type Stack struct {
	Top  int32
	Data []int32
}

func NewSeqStack(n int32) *Stack {
	return &Stack{
		Top:  -1,
		Data: make([]int32, n),
	}
}

func (s *Stack) Push(val int32) {
	s.Top += 1
	s.Data[s.Top] = val
}

func (s *Stack) Pop() int32 {
	val := s.Data[s.Top]
	s.Top -= 1
	return val
}
