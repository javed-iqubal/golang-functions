package main

import "fmt"

// Name return

/*

	func functionName(args) (variableName type){

		return
	}
*/

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {

	result := multiply(2, 5)
	fmt.Println("Result = ", result)
}
