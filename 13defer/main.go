package main

import "fmt"

func main() {

	// defer will execute last in first out
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello") // hello two one world
	myDefer()  // hello 4 3 2 1 0 two one world

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
