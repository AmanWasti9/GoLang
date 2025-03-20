package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")
	// No inheritance in golang; No super or parent

	Aman := User{"Amanullah", "amanwasti@gmail.com", true, 21}
	fmt.Println(Aman)
	fmt.Printf("Aman details are : %v \n", Aman)
	fmt.Printf("Name is %v and email is %v \n", Aman.Name, Aman.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
