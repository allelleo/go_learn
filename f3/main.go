package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := devise(4, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func devise(x int, y int) (float64, error) {
	if y == 0 {
		return float64(0), errors.New("На 0 делить нельзя")
	}
	return float64(x) / float64(y), nil
}
