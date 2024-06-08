package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user inputs"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for food ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)
	fmt.Printf("type of the rating is %T ", input)
}
