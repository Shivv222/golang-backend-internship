package main

import (
	"fmt"
)

func main() {

	//Normal switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("Four")
	// default:
	// 	fmt.Println("Others")
	// }

	//Multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	//      fmt.Println("It's working day")
	// }

	// type switch
	Whoami := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an intezer")
		case bool:
			fmt.Println("it's an bpplean")
		case string:
			fmt.Println("it's an string")
		default:
			fmt.Println("Others")

		}

	}

	Whoami("Shiv")

}
