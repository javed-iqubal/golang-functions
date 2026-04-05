package main

import "fmt"

/*
*
functions are first class citizen
- functions are values like ints, floats, strings, booleans, objects
eg:

	var a int = 10
	var is keyword
	a is variable
	int is datatype
	10 is value/literal

	like 10 we can assign functions into a variable

	var myfunc = function definition(){}

	we can invoke function using myfunc variable
*/
func main() {

	// with type
	var greet func() = func() {
		fmt.Println("Greeting..")
	}
	greet()

	// without type - type inference
	var hello = func() {
		fmt.Println("Hello world")
	}

	hello()

	// without type and var - shortcut
	hey := func() {
		fmt.Println("Say hi")
	}
	hey()

}
