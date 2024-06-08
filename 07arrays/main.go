package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitlist [4]string
	fruitlist[0] = "orange"
	fruitlist[1] = "apple"
	fruitlist[3] = "kiwi"
	fmt.Println(fruitlist)
	fmt.Println("length of fruitlist ", len(fruitlist))

	var vegList = [4]string{"potato", "beans", "mushroom"}
	fmt.Println("vegies are ", vegList)
	fmt.Println(len(vegList))
	fmt.Println(fruitlist == vegList)

	alphabets := [5]string{"a", "b", "c", "d"}
	fmt.Println("alphabelts ", alphabets)
	fmt.Println("length of alphabets ", len(alphabets))
}
