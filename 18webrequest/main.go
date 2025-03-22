package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response : %T \n", response)
	// fmt.Println("Response : \n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	content := string(databytes)
	fmt.Println("Content", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
