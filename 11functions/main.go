package main

import "fmt"

/**
IIFE - Immediately Invoked Function Expression
*/

func main() {

	result := func(a, b int) int {
		return a + b
	}(10, 20) // invoking function

	fmt.Println("Addition: ", result)

	// without vairiable
	func() {
		fmt.Println("Hello in annonymous function")
	}()

}
