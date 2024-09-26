package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitsList [3]string
	fruitsList[0] = "Apple"
	fruitsList[1] = "Banana"
	fruitsList[2] = "Orange"
	fmt.Println(fruitsList, len(fruitsList))

	var vegList = [3]string{"Potato", "Banana", "Orange"}
	fmt.Println(vegList, len(vegList))

}
