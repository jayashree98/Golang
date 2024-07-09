package main

import "fmt"

func main() {
	fmt.Println(add(1, 4))
	fmt.Println(addVarArgs(1, 4, 3, 6, 2, 1, 4))
	ananymaous()

}

func add(num1, num2 int) int {
	return num1 + num2
}

// func addVarArgs(num ...int) int {
// 	var total int
// 	for count := 0; count < len(num); count++ {
// 		total = total + num[count]
// 	}

// 	return total
// }

func addVarArgs(num ...int) int {
	total := 0
	for _, val := range num {
		// fmt.Println(index)
		total += val
	}
	return total
}

func ananymaous() {

	func() {
		fmt.Println("ananymous function")
	}()
}
