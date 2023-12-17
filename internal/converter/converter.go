package converter

import (
	"CalcInfixNotation/internal/stack"
	"errors"
)

var (
	fullStackError = errors.New("stack is full")
)

type Converter struct {
	stack *stack.Stack
}

func NewConverter(stack *stack.Stack) *Converter {
	return &Converter{stack: stack}
}

func (c *Converter) ConvertInfixToPostfix(infix string) (string, error) {

	postfix := ""

	for i, ch := range infix {
		if c.isOperand(ch) {
			postfix += string(ch)
			if len(infix)-1 == i || !c.isOperand(rune(infix[i+1])) {
				postfix += " "
			}
		} else if ch == '(' {
			if pushed := c.stack.Push(string(ch)); !pushed {
				return "", fullStackError
			}
		} else if ch == ')' {
			for top, exist := c.stack.Top(); exist && top != "("; top, exist = c.stack.Top() {
				op, _ := c.stack.Pop()
				postfix += op + " "
			}
			c.stack.Pop() // remove '(' from stack
		} else {
			if lastOperator, exist := c.stack.Top(); !exist {
				if pushed := c.stack.Push(string(ch)); !pushed {
					return "", fullStackError
				}
			} else if c.getOperatorPriority(lastOperator) < c.getOperatorPriority(string(ch)) {
				if pushed := c.stack.Push(string(ch)); !pushed {
					return "", fullStackError
				}
			} else {
				for top, exist := c.stack.Top(); exist && c.getOperatorPriority(top) >= c.getOperatorPriority(string(ch)); top, exist = c.stack.Top() {
					op, _ := c.stack.Pop()
					postfix += string(op) + " "
				}
				if pushed := c.stack.Push(string(ch)); !pushed {
					return "", fullStackError
				}
			}
		}
	}

	for !c.stack.IsEmpty() {
		op, _ := c.stack.Pop()
		postfix += string(op) + " "
	}

	return postfix, nil
}

func (c *Converter) getOperatorPriority(operator string) int {
	if operator == "+" || operator == "-" {
		return 1
	} else if operator == "*" || operator == "/" {
		return 2
	} else if operator == "^" {
		return 3
	} else {
		return 0
	}
}

func (c *Converter) isOperand(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
