package main

import "fmt"

func main() {
	fmt.Println("Structs")
	//there is no inheritance in go; so super or parent
	user := User{"Jyothis", 20, true, "jyothis@gmail.com"}
	fmt.Println(user)
	fmt.Printf("user details are %+v \n", user)
	fmt.Printf("user name:%v, email:%v \n", user.Name, user.Email)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}
