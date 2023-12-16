package main

import "fmt"

type User struct {
	username string
	email    string
	password string
}

func (user User) check_password(password string) bool {
	return user.password == password
}

func main() {
	user1 := User{
		username: "admin",
		email:    "admin@example.com",
		password: "admin",
	}

	input_password := "no_admin"
	input_password2 := "admin"

	fmt.Println(user1.check_password(input_password))
	fmt.Println(user1.check_password(input_password2))
}
