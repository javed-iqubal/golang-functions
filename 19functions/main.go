package main

import "fmt"

/*
clousre is an annonymous function capture/remembers variables from its surrounding scope
*/
func main() {

	// now inner function accessing outer function variable, this is called clousre
	// clousre
	x := 10

	innerFunc := func() {
		fmt.Printf("x inside inner function %d \n", x)
	}
	innerFunc()
}
