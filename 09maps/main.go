package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of languages", languages)
	fmt.Println("JS ", languages["JS"])

	delete(languages, "PY")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println("key %v value %v ", key, value)
	}
}
