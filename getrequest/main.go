package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // This is the line you need to add
	"net/url"  // Move this import statement to the top of the file
	"strings"
)

const geturl = "http://localhost:8000/get"
const posturl = "http://localhost:8000/post"
const formurl = "http://localhost:8000/postform"

func main() {
	//PerformGetRequest(myurl)
	//PerformPostJsonRequest(posturl)
	PerformFormRequest(formurl)

}

func PerformFormRequest(urlstr string) {

	//formdata

	data := url.Values{}
	data.Add("firstname", "jaya")
	data.Add("lastname", "shree")

	response, err := http.PostForm(urlstr, data)
	if err != nil {

		fmt.Println("Error fetching the url", err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content ", string(content))

}

func PerformPostJsonRequest(url string) {

	requestBody := strings.NewReader(`
	{
	"name":"jaya",
	"email":"jaya@gmail.com"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error fetching the url", err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content ", string(content))
}

func PerformGetRequest(url string) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the url", err)
	}

	defer response.Body.Close()

	fmt.Println("Response satatus code ", response.StatusCode)
	fmt.Println("Content length ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body", err)
	}
	// fmt.Println("content ", string(content))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount ", byteCount)
	fmt.Println("responseString ", responseString.String())

}
