package main

import "fmt"

func main() {
	user := struct {
		username string
		email    string
		password string
	}{
		username: "allelleo",
		email:    "dev.allelleo@internet.ru",
		password: "secret",
	}

	fmt.Println(user)
}
