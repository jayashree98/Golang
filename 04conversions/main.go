package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Food App ")
	fmt.Println("Please rate our food between 1 and 10")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	numRating, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("after the change ", numRating+1)
	}

}
