package converter

import (
	"CalcInfixNotation/internal/stack"
	"log/slog"
)

type Converter struct {
	stack  *stack.Stack
	logger *slog.Logger
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
			c.stack.Push(string(ch))
		} else if ch == ')' {
			for top, exist := c.stack.Top(); exist && top != "("; top, exist = c.stack.Top() {
				op, _ := c.stack.Pop()
				postfix += op + " "
			}
			c.stack.Pop() // remove '(' from stack
		} else {
			if lastOperator, exist := c.stack.Top(); !exist {
				c.stack.Push(string(ch))
			} else if c.getOperatorPriority(lastOperator) < c.getOperatorPriority(string(ch)) {
				c.stack.Push(string(ch))
			} else {
				for top, exist := c.stack.Top(); exist && c.getOperatorPriority(top) >= c.getOperatorPriority(string(ch)); top, exist = c.stack.Top() {
					op, _ := c.stack.Pop()
					postfix += string(op) + " "
				}
				c.stack.Push(string(ch))
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
