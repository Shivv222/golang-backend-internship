package main

import "fmt"

func main() {
	fmt.Println("Wewlcome to the class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointers is ", ptr)

	myNumber := 95
	var ptr = &myNumber

	fmt.Println("address of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr + 5
	fmt.Println("New value is ", myNumber)

}
