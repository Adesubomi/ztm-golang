//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(r Rectangle) int {
	return int(math.Abs(float64(r.b.x - r.a.x)))
}

func height(r Rectangle) int {
	return int(math.Abs(float64(r.b.y - r.a.y)))
}

func area(r Rectangle) int {
	return width(r) * height(r)
}

func perimeter(r Rectangle) int {
	return (width(r) + height(r)) * 2
}

func printInfo(r Rectangle) {
	fmt.Println("Rectangle,", r)
	fmt.Println("Area is", area(r))
	fmt.Println("Perimeter is", perimeter(r))
}

func main() {
	rect := Rectangle{
		a: Coordinate{0, 8},
		b: Coordinate{7, 0},
	}
	printInfo(rect)

	fmt.Println("- - - - - - Doubling Rectangle Size - - - - - - - -")
	rect.b.x *= 2
	rect.a.y *= 2
	printInfo(rect)
}
