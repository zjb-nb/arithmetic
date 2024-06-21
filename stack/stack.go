package stack

type StrStack struct {
	stack []string
}

func MakeStack() StrStack {
	return StrStack{}
}

func (s *StrStack) Push(val string) {
	s.stack = append(s.stack, val)
}

func (s *StrStack) Pop() string {
	res := s.Top()
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *StrStack) Top() string {
	return s.stack[len(s.stack)-1]
}

func (s *StrStack) Len() int {
	return len(s.stack)
}
