package main

import "fmt"

/**
Higher Order Function or HOF - Passing function as a parameter to another function

-> beacause function is a value..
*/

func welcome(callback func()) {
	// invoke the function by invoking variable called callback
	callback()
}

// The function to be passed

func sayHello() {
	fmt.Println("Hello!!")
}

var p func() = func() {
	fmt.Println("Function with variable and type")
}

func main() {

	welcome(sayHello)

	welcome(p)

	welcome(func() {
		fmt.Println("Annonymous function")
	})
}
