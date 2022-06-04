package stack

import "errors"

type stack []int


func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() bool {
	if s.isEmpty() {
		return false
	} else {
		top := len(*s) - 1
		*s = (*s)[:top]
		return true
	}
}

func (s *stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty...")
	} else {
		top := len(*s) - 1
		val := (*s)[top]
		return val, nil
	}
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) show() stack {
	return *s
}
