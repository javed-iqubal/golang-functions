package main

import "fmt"

// function with parameter
/*
	func functName(variable dataType){
		// function body
	}
*/

func greet(name string) {
	fmt.Printf("Hello, %s! \n", name)
}

func main() {

	greet("Javed")
}
