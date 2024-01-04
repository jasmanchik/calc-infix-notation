package calculator

import (
	"CalcInfixNotation/internal/converter"
	"CalcInfixNotation/internal/stack"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"
)

type Calculator struct {
	wg     *sync.WaitGroup
	stack  *stack.Stack
	logger *slog.Logger
}

func (c *Calculator) calculatePostfix(expression string) (float64, error) {
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

func (c *Calculator) Calculate(inputChan chan string, outChan chan float64) {
	defer c.wg.Done()
	const op = "postfix.Calculate"
	conv := converter.NewConverter(c.stack)

	for {
		select {
		case expression, ok := <-inputChan:
			if !ok {
				return
			}
			postfix, err := conv.ConvertInfixToPostfix(expression)
			if err != nil {
				c.logger.With(slog.String("op", op)).Error("can't convert infix to postfix", slog.String("expression", expression))
				break
			}
			result, err := c.calculatePostfix(postfix)
			if err != nil {
				c.logger.With(slog.String("op", op)).Error("can't calculate postfix expression", slog.String("postfix", postfix))
				break
			}
			outChan <- result
		default:
		}
	}
}

func NewCalculator(wg *sync.WaitGroup, logger *slog.Logger) *Calculator {
	s := stack.New()
	return &Calculator{wg: wg, stack: s, logger: logger}
}
