package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitlist = []string{"Apple", "Mango", "graphs"}
	fmt.Printf("Type of fruit list is%T\n", fruitlist)

	//adding extra fruits
	fruitlist= append(fruitlist, "cheery","watermelon","peach")
	fmt.Println("New list ", fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)

	highscore[0] =123
	highscore[1] =975
	highscore[2] =432
	highscore[3] =752

	highscore = append(highscore, 333,444,555,666,777,888)

	
	sort.Ints(highscore)

	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

}
