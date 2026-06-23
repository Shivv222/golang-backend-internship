package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	language := make(map[string]string)

	language["JS"] = "Javascrit"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages:", language)
	fmt.Println("JS shorts for: ", language["JS"]) //single

	delete(language, "RB") //delet
	fmt.Println("List of all languages:", language)

	//using loops

	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
