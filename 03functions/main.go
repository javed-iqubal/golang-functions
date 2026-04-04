package main

import "fmt"

// function with return type

/*
	func functionName(variable datatype) datatype{
		// function body

	}
*/

func add(x int, y int) int {

	return x + y
}
func main() {

	result := add(10, 20)

	fmt.Printf("Result : %d \n", result)

}
