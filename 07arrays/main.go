package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Carrot"
	fruitList[4] = "Peach"

	fmt.Println("Fruite list is : ", fruitList)
	fmt.Println("Fruite list is : ", len(fruitList))

	var vegList = [3]string{"potato", "bean", "mushroom"}
	fmt.Println("Veg list is : ", vegList)
	fmt.Println("Length", len(vegList))
}
