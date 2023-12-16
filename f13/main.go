package main

type User struct {
	username   string
	first_name string
	last_name  string
	email      string
	password   string
}

func (u *User) set_password(password string) {
	u.password = password
}

func main() {
	user := User{
		username:   "admin",
		first_name: "admin",
		last_name:  "admin",
		email:      "admin@example.com",
	}
	user.set_password("admin")
}
