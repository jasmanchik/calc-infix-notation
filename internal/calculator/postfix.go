package calculator

import (
	"CalcInfixNotation/internal/converter"
	"CalcInfixNotation/internal/stack"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func calculatePostfix(expression string) (float64, error) {
	s := make([]float64, 0)

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(s) < 2 {
				return 0, fmt.Errorf("not enough operands %s", token)
			}

			operand2 := s[len(s)-1]
			operand1 := s[len(s)-2]
			s = s[:len(s)-2]

			var result float64
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

			s = append(s, result)

		default:
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number or operator format: %s", token)
			}
			s = append(s, number)
		}
	}

	if len(s) != 1 {
		return 0, fmt.Errorf("incorrect number of operands and operators in the expression")
	}

	return s[0], nil
}

func Calculate(logger *slog.Logger, stack *stack.Stack, expressionChan chan string, outChan chan float64) {
	c := converter.NewConverter(stack)
	for expression := range expressionChan {
		postfix, err := c.ConvertInfixToPostfix(expression)
		if err != nil {
			logger.Error("can't convert infix to postfix", slog.String("expression", expression))
			continue
		}

		result, err := calculatePostfix(postfix)
		if err != nil {
			logger.Error("can't calculate postfix expression", slog.String("postfix", postfix))
			continue
		}
		outChan <- result
	}
}
