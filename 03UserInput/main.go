package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the playground!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("name is %S", name)
}
