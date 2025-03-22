package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files concepts")

	content := "This needs to go in file - Amanullah Wasti"

	file, err := os.Create("./myfile1.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is", length)

	defer file.Close()

	// Call the function
	readFile("./myfile1.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
