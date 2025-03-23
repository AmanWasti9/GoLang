package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=ghjs453hdj"

func main() {
	fmt.Println("URL Handlings in Go lang")

	fmt.Println("URL", myurl)

	// parsing the URL
	result, err := url.Parse(myurl)
	checkNilErr(err)

	fmt.Println("result", result.Scheme)
	fmt.Println("result", result.Host)
	fmt.Println("result", result.Path)
	fmt.Println("result", result.Port())
	fmt.Println("result", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T \n", qparams)

	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	// MAKING URL FROM CHUNKS
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=aman",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("Url", anotherUrl)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
