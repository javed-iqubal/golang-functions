package main

import (
	"fmt"
)

//

func login(resolve func(), reject func()) {

	loggedIn := true
	if loggedIn {
		resolve()
	} else {
		reject()
	}

}

func main() {

	// passed two functions as parameter

	login(func() {
		fmt.Println("Login Success")
	}, func() {
		fmt.Println("Login failed")
	})
}
