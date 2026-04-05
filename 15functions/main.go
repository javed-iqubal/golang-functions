package main

import (
	"fmt"
)

func login(resolve func(status string) int, reject func(err string) int) {
	loggedIn := false
	var statuscode int
	if loggedIn {
		statuscode = resolve("Login Success")
	} else {
		statuscode = reject("Login Failed")
	}
	fmt.Println(statuscode)
}
func main() {

	login(func(status string) int {
		fmt.Println(status)
		return 200
	}, func(err string) int {
		fmt.Println(err)
		return 500
	})
}
