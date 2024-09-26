package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	days := []string{"sanday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	//3 diff ways to use for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])

	}
	for _, day := range days {
		fmt.Println(day)
	}
	for d := range days {
		fmt.Println(d, days[d])
	}

	// gothrough
	number := 0
	for number < 10 {
		if number == 4 {
			goto loco
		}
		if number == 5 {
			number++
			continue
		}
		if number == 8 {
			number++
			break
		}
		fmt.Println(number)
		number++
	}
loco:
	fmt.Println("hey goto")
}
