//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printStats(parts []Part) {
	fmt.Printf("\nAssembly line with %v parts\n", len(parts))

	for i := 0; i < len(parts); i++ {
		fmt.Println("Part", i+1, ":", parts[i])
	}
}

func main() {
	//  - Create an assembly line having any three parts
	assemplyLine := []Part{"Late", "Press", "Paint Shop"}
	printStats(assemplyLine)

	//  - Add two new parts to the line
	assemplyLine = append(assemplyLine, "Logs", "Packaging")
	printStats(assemplyLine)

	//  - Slice the assembly line so it contains only the two new parts
	slice := assemplyLine[len(assemplyLine)-2:]
	printStats(slice)

	//  - Print out the contents of the assembly line at each step
	fmt.Printf("\n")
}
