package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["J"] = "Java"
	languages["GL"] = "GOLang"
	languages["RB"] = "Ruby"

	fmt.Println("Languages", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Languages", languages)

	// Loops are interesting in go lang
	for _, value := range languages {
		fmt.Printf("For key v, value is %v \n", value)
	}

}
