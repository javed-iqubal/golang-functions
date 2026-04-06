package main

import "fmt"

/**
clousure - returning function from another function is called clousure
*/

func counter() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	next := counter()

	fmt.Printf("Next value is: %d \n", next())
	fmt.Printf("Next value is: %d \n", next())
	fmt.Printf("Next value is: %d \n", next())
	fmt.Printf("Next value is: %d \n", next())
	fmt.Printf("Next value is: %d \n", next())

	fmt.Println()
	nextState := counter()

	fmt.Printf("NextState value is: %d \n", nextState())
	fmt.Printf("NextState value is: %d \n", nextState())
	fmt.Printf("NextState value is: %d \n", nextState())
	fmt.Printf("NextState value is: %d \n", nextState())
	fmt.Printf("NextState value is: %d \n", nextState())

}
