package main

import "fmt"

func login(resolve func(status string), reject func(err string)) {
	loggedIn := false

	if loggedIn {
		resolve("LoggedIn success")
	} else {
		reject("loggedIn failed")
	}
}
func main() {

	login(func(status string) {
		fmt.Println(status)
	}, func(err string) {
		fmt.Println(err)

	})

}
