package main

import "fmt"

func main() {
	user := User{"Jyothis", 24, true, "jyothis@gmail.com"}
	fmt.Println(user)
	fmt.Printf("user details are %+v \n", user)
	fmt.Printf("user name:%v, email:%v \n", user.Name, user.Email)

	user.getStatus()
	user.increaseAge()
	fmt.Println("updated age outside is", user.Age)
	user.increaseAgeRef()
	fmt.Println("updated age outside is", user.Age)

}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (u User) getStatus() {
	fmt.Println("Name of the user is:", u.Name)
}

func (u User) increaseAge() {
	u.Age = u.Age + 1
	fmt.Println("Increased age is:", u.Age)
}
func (u *User) increaseAgeRef() {
	u.Age = u.Age + 1
	fmt.Println("Increased age is:", u.Age)
}
