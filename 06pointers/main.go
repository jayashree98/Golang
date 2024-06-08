package main

import "fmt"

func main() {
	// fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("pvalue of ptr is ", ptr)

	// myNumber := 23
	// var myptr = &myNumber
	// fmt.Println(myptr)
	// fmt.Println(*myptr)
	// fmt.Println(&myNumber)
	// *myptr = *myptr + 2
	// fmt.Println(myNumber)

	var x int = 123
	fmt.Println("value of x ", x)
	var a *int
	a = &x
	fmt.Println("address of x ", &x)
	fmt.Println("value of a ", a)
	fmt.Println("address of a ", &a)

}
