package main

import "fmt"

// function return multiple values
/*
	func functionName(args) (type, type,type,...type){
		return var1,var2,var3, ... varN
	}
*/

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func divide_v2(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func main() {

	quotient, remainder := divide(20, 6)
	fmt.Printf("Quotient = %d, Remainder = %d \n", quotient, remainder)

	quotient, remainder = divide_v2(25, 7)
	fmt.Printf("Quotient = %d, Remainder = %d \n", quotient, remainder)
}
