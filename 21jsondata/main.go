package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // Do not display the password
	Tags     []string `json:"tags,omitempty"` // omitempty will display not display the field if it is empty
}

func main() {
	fmt.Println("Json data")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"Python Programming", 499, "Udemy", "python123", []string{"numpy", "python"}},
		{"React Js BootCamp", 1005, "Coursera", "react123", []string{"react", "redux"}},
		{"Go Lang Programming", 5001, "Coursera", "go123", nil},
	}

	// Package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkNilErr(err)

	fmt.Printf("Result %s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Python Programming",
		"Price": 499,
		"website": "Udemy",
		"tags": ["numpy","python"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSON was not Valid")
	}

	// Some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T \n", k, v, v)
	}

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
