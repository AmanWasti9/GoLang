package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, mymsg := proAdder(3, 5, 2, 7)
	fmt.Println("Result is: ", proResult)
	fmt.Println("Message is: ", mymsg)

}

// (valOne int, valTwo int) this specifies the function to what value pass on
// int this specifies the function to what value to be return
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// func proAdder(values ...int) int {
// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}
// 	return total
// }
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro adder result function"
}

func greeter() {
	fmt.Println("Hello from golang")
}
