package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculatePostfix(expression string) (int, error) {
	stack := make([]int, 0)

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough operands %s", token)
			}

			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, fmt.Errorf("divizion by zero")
				}
				result = operand1 / operand2
			}

			stack = append(stack, result)

		default:
			number, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("invalid number or operator format: %s", token)
			}
			stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("incorrect number of operands and operators in the expression")
	}

	return stack[0], nil
}
