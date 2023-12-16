package main

import "fmt"

func main() {
	res := add(1, 5)
	fmt.Println(res)
}

func add(x int, y int) int {
	return x + y
}