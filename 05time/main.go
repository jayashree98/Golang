package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.April, 10, 12, 23, 4, 5, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006"))
}
