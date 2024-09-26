package main

import "fmt"

func main() {
	fmt.Println("hello")

	//LIFO: last in one first out
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
}
