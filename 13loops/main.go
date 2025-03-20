package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Comanly used for loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("The value is %v\n", day)
	}

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			break
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at learn code")

}
