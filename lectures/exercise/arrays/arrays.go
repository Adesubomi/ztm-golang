//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

type ShoppingList [4]Product

func (sL ShoppingList) printStat() {
	var totalCost float64
	var totalItems int

	for i := 0; i < len(sL); i++ {
		if sL[i].name == "" {
			continue
		}
		totalItems++
		totalCost += sL[i].price
	}

	fmt.Println("Last Item", sL[totalItems-1])
	fmt.Println("Total number of items", totalItems)
	fmt.Println("Total cost of items", totalCost)
}

func main() {
	var shoppingList ShoppingList
	shoppingList[0] = Product{name: "Cassava", price: 13.99}
	shoppingList[1] = Product{name: "Garri", price: 14.99}
	shoppingList[2] = Product{name: "Ewa", price: 22.99}
	shoppingList.printStat()

	fmt.Println(" - - - Adding one more item - - ")
	shoppingList[3] = Product{name: "Agbado", price: 2.99}
	shoppingList.printStat()
}
