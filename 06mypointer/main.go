package main

import "fmt"

func main() {
	fmt.Println("Pointer concepts")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is: ", ptr)
	fmt.Println("Value of actual pointer is: ", *ptr)

	// For Reference: & to use
	// For value: * to use

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber)

}
