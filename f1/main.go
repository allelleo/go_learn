package main

import "fmt"

const login string = "allelleo"
const password string = "secret"

func main() {
	userLogin := "allelleo"
	userPasssword := "secret"

	if login == userLogin && password == userPasssword {
		fmt.Println("login success")
	} else {
		fmt.Println("login fail")
	}
}
