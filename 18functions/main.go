package main

import "fmt"

/**
 clousre is an annonymous function capture/remembers variables from its surrounding scope

 scope:
	- function scope - any variables/functions/code within function - local scope
	- package scope - any variables/functions/code outside function within package
*/

// package scope
var x int = 10

func getX() {
	fmt.Printf("x inside getX is %d \n", x)
}

func main() {
	// local scope
	var x int = 20

	fmt.Printf("x inside main is %d \n", x)
	getX()
}
