package main

import "fmt"

type user struct {
	username string
	email    string
	password string
}

func main() {
	user1 := user{
		username: "allelleo",
		email:    "dev.allelleo@internet.ru",
		password: "secret",
	}
	user2 := user{
		username: "user2",
		email:    "dev.user2@internet.ru",
		password: "secret",
	}
	fmt.Println(user1)
	fmt.Println(user2)
}
