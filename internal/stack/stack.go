package stack

import (
	"strconv"
)

type Stack struct {
	list []string
	len  int
	size int
}

func New(maxSize int) Stack {
	return Stack{
		list: make([]string, maxSize),
		len:  0,
		size: maxSize,
	}
}

func (s *Stack) IsFull() bool {
	return s.len == s.size
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Push(element string) bool {
	if s.IsFull() {
		return false
	}

	s.list[s.len] = element
	s.len += 1
	return true
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
