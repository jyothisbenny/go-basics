package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	nullCheck(err)

	io.WriteString(file, "Hello World")

	read, err := os.ReadFile("./test.txt")
	nullCheck(err)

	fmt.Println(string(read))

	defer file.Close()
}

func nullCheck(err error) {
	if err != nil {
		panic(err)
	}

}
