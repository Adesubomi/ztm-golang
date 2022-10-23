//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	add Operation = iota
	sub
	mul
	div
)

func (o Operation) calculate(operator float64, operand float64) float64 {
	if o == add {
		return operator + operand
	} else if o == sub {
		return operator - operand
	} else if o == mul {
		return operator * operand
	} else if o == div {
		return operator / operand
	}

	return 0
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
