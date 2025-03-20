package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices datatype in Go Lang")

	var fruitList = []string{"Apple", "Peach", "Tomato"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)
	// Output
	// Type of fruitList is []string

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	// Output
	// [Apple Peach Tomato Mango Banana]

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 245
	highScore[3] = 259
	// highScore[4] = 785

	highScore = append(highScore, 555, 478, 562)

	fmt.Println("HighScores", highScore)

	// For sorting
	sort.Ints(highScore)
	fmt.Println(highScore)

	// It will return false value
	fmt.Println(sort.IntsAreSorted(highScore))

	// 	Remove the value from slice based on index
	var courses = []string{"reactjs", "java", "python", "ruby", "golang"}
	fmt.Println("courser", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
