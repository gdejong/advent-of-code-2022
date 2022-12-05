package stack

type Stack struct {
	contents []string
}

func New() *Stack {
	return &Stack{
		contents: []string{},
	}
}

func (s *Stack) Push(data string) {
	s.contents = append(s.contents, data)
}

func (s *Stack) Unshift(data string) {
	s.contents = append([]string{data}, s.contents...)
}

func (s *Stack) Pop() string {
	if len(s.contents) == 0 {
		panic("cannot Pop an empty stack")
	}

	// Get top element
	data := s.contents[len(s.contents)-1]

	// Remove last element
	s.contents = s.contents[:len(s.contents)-1]

	return data
}

func (s *Stack) Peek() string {
	if len(s.contents) == 0 {
		panic("cannot Peek an empty stack")
	}

	return s.contents[len(s.contents)-1]
}

func (s *Stack) Len() int {
	return len(s.contents)
}
