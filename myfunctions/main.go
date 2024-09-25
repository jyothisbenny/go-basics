package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	greet()
	fmt.Println(greetName("josan"))
	fmt.Println(add(1, 3))

	res, err := divide(1, 0)
	fmt.Println(res, err)

	total := variadic(1, 3, 5, 6)
	fmt.Println(total)

}

func greet() {
	fmt.Println("Hello Jyothis")
}

func greetName(name string) string {
	return "Hello " + name
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("worng divider")
	}
	return a / b, nil
}

func variadic(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
