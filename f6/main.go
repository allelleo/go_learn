package main

import "fmt"

func main() {
	var users = [...]string{
		"user1",
		"user2",
		"user3",
		"user4",
	}

	var logins = [len(users)]string{}

	for index, value := range users {
		logins[index] = value
	}

	fmt.Println(logins)
}
