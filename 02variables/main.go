package main

import "fmt"

func main() {
	var inte int = 2
	fmt.Println(inte) //output: 2.3344347
	fmt.Printf("Type of variable is %T \n", inte)

	var smallVal float32 = 2.33443475983475
	fmt.Println(smallVal) //output: 2.3344347
	fmt.Printf("Type of variable is %T \n", smallVal)

	var bigVal float64 = 2.33443475983475
	fmt.Println(bigVal) // output: 2.33443475983475
	fmt.Printf("Type of variable is %T \n", bigVal)

	var isBool bool = false
	fmt.Println(isBool) // false
	fmt.Printf("Type of variable is %T \n", isBool)

	//explicit type
	var name string = "Jyothis"
	fmt.Println(name)

	//implicit type
	var name2 = "Jyothis"
	fmt.Println(name2)

	//walrus op, Note: cant use this method for global declaration
	name3 := "Jyothis"
	fmt.Println(name3)

}
