package main

type Stack struct {
	data []interface{}
	top  int
}

func (s *Stack) push(data interface{}) {
	s.data = append(s.data, data)
	s.top++
}

func (s *Stack) pop() interface{} {
	var last interface{}
	if s.top > 0 {
		s.data = s.data[:s.top-1]
		last = s.data[s.top]
		s.top--
		return last
	}
	return nil
}
