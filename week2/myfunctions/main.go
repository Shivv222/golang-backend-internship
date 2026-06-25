package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greet()

	result := adder(4, 5)

	fmt.Println("Result is: ", result)

	proRes, mymess := proAdder(5,5,5,10,4,1,100)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Mymessgae is: ", mymess)

}

func adder(valone int, valtwo int) int  {
	return valone + valtwo
	
}

func greet()  {
	fmt.Println("Namstey from server")
}

func proAdder(values...int) (int, string) {
	total :=0

	for _, val := range values {
		total += val
	}
	return   total, "hi from shiv"

}