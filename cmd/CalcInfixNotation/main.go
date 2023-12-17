package main

import (
	"CalcInfixNotation/internal/calculator"
	"CalcInfixNotation/internal/converter"
	"CalcInfixNotation/internal/stack"
	"fmt"
	"log"
)

func main() {

	s := stack.New(50)
	c := converter.NewConverter(s)

	infix := "(944+397)*(639-221)+751/83-236*74" //543083.048192771
	postfix, err := c.ConvertInfixToPostfix(infix)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("postfix:", postfix)

	result, err := calculator.CalculatePostfix(postfix)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("result:", result)
}
