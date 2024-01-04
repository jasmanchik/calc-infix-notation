package stack

import (
	"strconv"
)

type Stack struct {
	list []string
	len  int
}

func New() *Stack {
	return &Stack{
		list: make([]string, 10),
		len:  0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Push(element string) {
	s.list[s.len] = element
	s.len += 1
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	top := s.list[s.len-1]
	s.len--
	result := top
	for !s.IsEmpty() {
		top = s.list[s.len-1]
		if top == " " || top == "" || !s.isDigit(top) {
			break
		}

		result = top + result
		s.len--
	}

	return result, true
}

func (s *Stack) isDigit(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}

func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	return s.list[s.len-1], true
}

func (s *Stack) Clear() {
	s.len = 0
}
