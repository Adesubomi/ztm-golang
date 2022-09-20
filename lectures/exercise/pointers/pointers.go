//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name  string
	state SecurityTag
}

func activateTag(state *SecurityTag) {
	*state = Active
}

func deactivateTag(state *SecurityTag) {
	*state = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivateTag(&items[i].state)
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	antelope := Item{name: "Antelope", state: Active}
	dove := Item{name: "Dove", state: Active}
	lion := Item{name: "Lion", state: Active}
	platypus := Item{name: "Platypus", state: Active}

	//  - Store them in a slice or array
	items := []Item{antelope, dove, lion, platypus}
	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	deactivateTag(&items[0].state)
	deactivateTag(&items[2].state)
	fmt.Println(items)

	// oversabi
	activateTag(&items[2].state)
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println(items)

	//  - Print out the array/slice after each change
}
