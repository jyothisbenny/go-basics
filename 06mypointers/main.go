package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var ptr *int
	fmt.Println(ptr)

	myNumber := 10
	myNumberPtr := &myNumber
	fmt.Println(*myNumberPtr)
	fmt.Println(myNumberPtr)

	*myNumberPtr = *myNumberPtr + 2

	fmt.Println(myNumber)
}
