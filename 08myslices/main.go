package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to my slices")
	firstSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(firstSlice)

	firstSlice = append(firstSlice, 7, 8, 9)
	fmt.Println(firstSlice)

	firstSlice = append(firstSlice[1:5])
	fmt.Println(firstSlice)

	firstSlice = firstSlice[1:2]
	fmt.Println(firstSlice)
	fmt.Println(cap(firstSlice), len(firstSlice))

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	highScores[3] = 400
	fmt.Println(firstSlice)

	highScores = append(highScores, highScores...)
	fmt.Println(highScores)
	//highScores[8] = 900 only append words, because it creates a new pointer, existing pointer have length with got initialized
	fmt.Println(cap(highScores), len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove index 2

	highScores = append(highScores[:2], highScores[3:]...)
	fmt.Println(highScores)
}
