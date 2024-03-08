package genericstack

import (
	"errors"
)

// -> It holds elements of any type.

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(el interface{}) {
	s.data = append(s.data, el)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	lastEl := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return lastEl, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
