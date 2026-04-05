package main

import "fmt"

// variable with inference - recomended

var hey = func() {
	fmt.Println("Say hello functional programming")
}

func main() {
	hey()
}
