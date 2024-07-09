package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.geeksforgeeks.org/how-to-trim-a-string-in-golang/?ref=lbp"

func main() {

	result, _ := url.Parse(myurl)
	fmt.Println("Scheme ", result.Scheme)
	fmt.Println("Host ", result.Host)
	fmt.Println("Port ", result.Port())
	fmt.Println("Path ", result.Path)
	fmt.Println("Raw Query ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Query params %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("Param value", val)
	}

	pathOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.geeksforgeeks.org",
		Path:    "/how-to-trim-a-string-in-golang/?",
		RawPath: "ref=lbp",
	}

	anotherURL := pathOfUrl.String()
	fmt.Println("Another URL ", anotherURL)

}
