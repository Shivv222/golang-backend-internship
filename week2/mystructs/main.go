package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 20}
	shiv := User{"Shiv", "shiv@go.dev", true, 21}
	gautam := User{"Gautam", "gautam@go.dev", true, 21}

	fmt.Printf("User details are: %+v\n", gautam)
	fmt.Printf("User details are: %+v\n", shiv)
	fmt.Printf("User details are: %+v\n", hitesh)

	fmt.Printf("User is: %v && Email is: %v", shiv.Name, shiv.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
