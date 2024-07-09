package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "Welcome to file handling in go"

	file, err := os.Create("./sample.txt")
	if err != nil {
		fmt.Println("error creating file", err)
	}

	len, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("error writing to file", err)
	}
	fmt.Println("length of content is ", len)

	defer file.Close()
	ReadFile("./sample.txt")
}

func ReadFile(filename string) {

	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file", err)
	}
	fmt.Println("dataBytes ", dataBytes)
	data := string(dataBytes)
	fmt.Println("data from file is ", data)
}
