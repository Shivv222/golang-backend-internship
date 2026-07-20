package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Nmae     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Bitmorejson (encoding, decoding)")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnDev.online.in", "abc12345", []string{"Web-dev", "js"}},
		{"MERN Bootcamp", 599, "LearnDevonline.in", "bcd12345", []string{"Fullstack", "js", "mern"}},
		{"Angular Bootcamp", 299, "LearnDevonline.in", "cde12345", nil},
	}

	//Package this data as Json Data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", &finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnDev.online.in",
                "tags": ["Web-dev","js"]
    }
	`)

	var lcoCourse course

	Checkvalid := json.Valid(jsonDataFromWeb)

	if Checkvalid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("JSON WAS INVALID")
	}

}
