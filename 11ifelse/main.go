package main

import "fmt"

func main() {
	fmt.Println("If else concept")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regulate user"
	} else if loginCount > 10 {
		result = "Allow user"
	} else {
		result = "Invalid login"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Numbe is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
