package main

import "fmt"

func main() {

	jaya := User{"jay", "jaya@gmail.com", true, 25}
	fmt.Printf("User jaya %v \n ", jaya)
	fmt.Printf("user details of jaya is %+v ", jaya)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
