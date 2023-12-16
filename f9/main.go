package main

func main() {

	var users = make(map[string]string)
	users["username"] = "user1"

	delete(users, "username")
}
