package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println("Phrase", phrase)
	fmt.Println("hello world :: ", hello, world)

	fmt.Println("")
	fmt.Println(" - - - Pointers - - - ")
	fmt.Println("hello", &hello)
	fmt.Println("world", &world)
	fmt.Println("phrase", &phrase)
	fmt.Println("phrase[0]", &phrase[0])
}
