package main

import "fmt"

// recursive function

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {

	n := 5
	fmt.Printf("Factorial of %d is %d \n", n, factorial(n))

}
