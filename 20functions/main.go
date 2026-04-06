package main

import "fmt"

func main() {

	// how clousure can read and modify the outer scope variable

	counter := 0

	// clousure function
	increment := func() {
		counter++
		fmt.Printf("Counter: %d \n", counter)
	}
	increment()
	increment()
	increment()
	increment()
	increment()
}
