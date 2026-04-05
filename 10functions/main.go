package main

import "fmt"

// Anonymous function - is a function without name

func main() {

	add := func(a, b int) int {
		return a + b
	}

	fmt.Printf("Addition: %d \n", add(10, 20))
}
