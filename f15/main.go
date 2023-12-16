package main

import "fmt"

func main() {
	var username string = "allelleo"

	switch username {
	case "allelleo":
		fmt.Println("Hello Alberto!")
	case "juanito":
		fmt.Println("Hello Juanito!")
	default:
		fmt.Println("Hello, I don't know you!")
	}
}
