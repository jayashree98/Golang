package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println("Welcome to Slices")
	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type of fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "Banana", "mango")
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highScores := make([]int, 3)
	highScores[0] = 21
	highScores[1] = 34
	highScores[2] = 13
	//highScores[4] = 56 // will give error as index out of range [4] with length
	highScores = append(highScores, 15, 22, 43, 54)
	fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a vaue from the slices

	var courses = []string{"angular", "go", "python", "java", "c"}
	//	var courses1 = []string{"angular", "go", "python", "java", "c"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	fmt.Println(courses == nil)
	// fmt.Println(courses == courses1) // will give error as slices can only be compared to nil

	//trimming the slices

	var trimmingslice = []byte{'@', '3', '4', '5', '6', '7', '8', '9', '0', '@'}
	fmt.Printf("trim slices %s\n", trimmingslice)

	trimmingsliceResult := bytes.Trim(trimmingslice, "@3")
	fmt.Printf("after trim %s ", trimmingsliceResult)
}
