package main

import "fmt"

func main() {
	shoppingLists := make(map[string]int)
	shoppingLists["eggs"] = 11
	shoppingLists["milk"] = 1
	shoppingLists["bread"] += 1

	shoppingLists["eggs"] += 1
	fmt.Println(shoppingLists)

	delete(shoppingLists, "milk")
	fmt.Println("Milk deleted, new list:", shoppingLists)

	fmt.Println("Need", shoppingLists["eggs"], "eggs")

	cereal, found := shoppingLists["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("    Nope, cereal not on list")
	} else {
		fmt.Println("    Yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingLists {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "items on the shopping list")

}
