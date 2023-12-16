package main

import "fmt"

func main() {
	var dictionary = make(map[string]string)

	dictionary["username"] = "allelleo"
	dictionary["password"] = "secret"
	dictionary["email"] = "dev.allelleo@internet.ru"

	value, exists := dictionary["first_name"]
	if !exists {
		fmt.Println("Такого значения нет в словаре!")
	} else {
		fmt.Println(value)
	}
}
