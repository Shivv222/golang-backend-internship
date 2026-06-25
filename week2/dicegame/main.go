package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to dice game (Switch case)")

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value od dice is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("Tou need to check")

	}

}
