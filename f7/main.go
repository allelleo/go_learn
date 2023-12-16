package main

import "fmt"

func main() {
	// Create a new map and add two key-value pairs to it
	var dict = make(map[string]int)

	dict["users"] = 5
	dict["username"] = 1122

	// Print the map
	fmt.Println(dict)

	// Access the value associated with the "username" key and print it
	fmt.Println(dict["username"])
}
