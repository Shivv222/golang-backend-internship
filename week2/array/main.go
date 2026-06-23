package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitlist [5]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	fruitlist[4] = "peach"

	fmt.Println("Fruitlist is: ", fruitlist)
	fmt.Println("Fruitlist is of length: ", len(fruitlist))

	var vegList = [5]string{"ladyfinger", "potato", "onion"}
	fmt.Println("Vegtable List is: ", vegList)

}
