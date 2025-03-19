package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input !!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// Comma ok  or error ok syntax use instead of try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T", input)

	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
	}

	value, ok := myMap["apple"]
	if ok {
		fmt.Println("Value found ", value)
	} else {
		fmt.Println("Value not found ")
	}

}
