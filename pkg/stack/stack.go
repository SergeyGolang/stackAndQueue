package stack

import (
	"errors"
)

/*
A stack implements LIFO (Last In, First Out).
A stack has four methods:
 1. Push - append a value
 2. Pop - return and remove the last element
 3. Peek - retrieve the last element
 4. IsEmpty - check stack if the stack is empty
*/

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]

	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0

}
