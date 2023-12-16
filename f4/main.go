package main

import "fmt"

// main is the entry point of the program
func main() {
	// create a slice of strings with three elements
	var todo = [3]string{
		"todo1",
		"todo2",
		"todo3",
	}

	// create a slice of integers with six elements
	fmt.Println([...]int{1, 2, 3, 4, 5, 6})

	// loop through each element in the slice and print the index and value
	for index, item := range todo {
		fmt.Println("Todo: ", index, " do: ", item)
	}
}
