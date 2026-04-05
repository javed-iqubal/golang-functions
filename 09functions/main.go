package main

import "fmt"

func main() {

	// with type
	// add => is vairiable
	// func(int, int) int => is type
	// func(a, b int) int {
	//	return a + b
	// }	=>	is value

	var add func(int, int) int = func(a, b int) int {
		return a + b
	}

	a := 10
	b := 20

	fmt.Printf("Sum of %d & %d is: %d \n", a, b, add(a, b))

	var substract = func(a, b int) int {
		return a - b
	}

	a = 60
	b = 20

	fmt.Printf("Substraction of %d & %d is: %d \n", a, b, substract(a, b))
}
