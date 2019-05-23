package stack

const arraySize = 100

type Stack struct {
	top  rune
	data [arraySize][4]rune
}

func (s *Stack) Push(i [4]rune) bool {
	if s.top == rune(len(s.data)) {
		return false
	}
	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack) Pop() (*[4]rune, bool) {
	if s.top == 0 {
		return nil, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return &i, true
}

func (s *Stack) Peek() [4]rune {
	return s.data[s.top-1]
}

func (s *Stack) Get() [][4]rune {
	return s.data[:s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Empty() {
	s.top = 0
}
