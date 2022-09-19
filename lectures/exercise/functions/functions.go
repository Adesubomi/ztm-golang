//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello,", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func getMessage() string {
	return "Hi there"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func randomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)
	return gen.Intn(100)
}

//* Write a function that returns any two numbers
func randomNumberDuo() (int, int) {
	return randomNumber(), randomNumber()
}

func main() {
	greetPerson("Daniel")
	fmt.Println(getMessage())

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	duo1, duo2 := randomNumberDuo()
	addition := addThreeNumbers(duo1, duo2, randomNumber())
	fmt.Println("Result of three numbers:", addition)
}
