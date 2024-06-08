package main

import "fmt"

func main() {

	Hitesh := User{"Jaya", "jaya@gmail.com", true, 25}
	fmt.Println("hitesh details are ", Hitesh)
	fmt.Printf("Name %v and Emial %v ", Hitesh.Name, Hitesh.Email)
	Hitesh.GetValues()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetValues() {
	fmt.Println(u.Status)
	u.Email = "divi@gmail.com"
	fmt.Println(u.Email)
}
