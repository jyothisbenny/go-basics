package main

import (
	"errors"
	"fmt"
)

// someFunction returns the square of a number if positive, else returns an error.
func someFunction() (int, error) {
	num := -1
	if num < 0 {
		return 0, errors.New("input cannot be negative")
	}
	return num * num, nil
}

func main() {
	fmt.Println("Hello World")
	number := 9

	if number < 10 {
		fmt.Println("number is less than 10")
	} else if number > 10 {
		fmt.Println("number is greater than 10")
	} else {
		fmt.Println("number is equal to 10")
	}

	if num := 3; num < 4 {
		fmt.Println("num is less than 4")
	}

	if value, err := someFunction(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}
}
