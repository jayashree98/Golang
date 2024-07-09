package main

import (
	"fmt"
	"net/http"
)

const url = "https://www.google.com"

func main() {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the url", err)
	}

	fmt.Printf("Response is of type %T\n", response)

	fmt.Println("Response code", response.StatusCode)

	defer response.Body.Close() // callers responsibility to close the response body

	// databytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body", err)
	// }
	//fmt.Println("Data from the response", string(databytes))

}
