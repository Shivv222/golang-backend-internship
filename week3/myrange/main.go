package main

import "fmt"

func main() {
	nums := []int{5, 6, 7, 8, 9}

	for i, num := range nums {
		fmt.Println(num, i)
	}

}
