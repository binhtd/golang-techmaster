package main

import (
	"fmt"
)

type Operator func(a int, b int) int

type Poperator *Operator

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Devide(a int, b int) int {
	return a / b
}

func (op Operator) compute(a int, b int) int {
	return op(a, b)
}

func (op *Operator) compute2(a int, b int) int {
	return (*op)(a, b)
}

func DemoOperator() {
	var op Operator
	op = func(a int, b int) int {
		return a + b
	}
	fmt.Println(op.compute(1, 2))
	fmt.Println(op.compute2(3, 4))

	op = Subtract
	fmt.Println(op.compute(6, 5))
	fmt.Println(op.compute2(6, 5))

	var op2 Poperator
	op2 = &op
	fmt.Println((*op2).compute(9, 7))
}
