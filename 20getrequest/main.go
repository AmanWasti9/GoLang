package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET Request")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)
	checkNilErr(err)

	defer response.Body.Close()
	// check response status code
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	// FIRST WAY TO READ AND DISPLAY THE RESPONSE

	// ioutil use to read the response body
	// content, err := ioutil.ReadAll(response.Body)
	// Content is in a byte format so we have to convert it into string

	// checkNilErr(err)

	// fmt.Println(string(content))

	// SECOND WAY TO READ AND DISPLAY THE RESPONSE WITH LIBRARY WHICH IS MORE POWERFUL
	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is : ", byteCount)

	fmt.Println("Response is : ", responseString.String())

}

func PerformPostRequest() {
	const myURL = "http://localhost:8000/post"

	// fake json payload or data
	requestBody := strings.NewReader(`
		{
			"name": "Amanullah",
			"coursename": "Go Lang",
			"price": 10
		}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("response", string(content))

}

func PerformPostJsonRequest() {
	const myURL = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Amanullah")
	data.Add("lastname", "Wasti")
	data.Add("email", "amanwasti5@gmail.com")

	response, err := http.PostForm(myURL, data)
	checkNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response", string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
