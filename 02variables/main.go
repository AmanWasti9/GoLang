package main

import "fmt"

// the LoginToken is publically accessible because of Capital "L"
const LoginToken string = "ajkakanajk"

func main() {
	var username string = "Amanullah"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 256.565856485566665
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var num int
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	var name string
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	// implicit type
	var name1 = "Amanullah Wasti"
	fmt.Println(name1)

	// no var style
	// This declaration can be allowed inside the any method but not for global declaration
	numberOfUser := 5000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
