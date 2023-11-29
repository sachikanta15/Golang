package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Golang Web request")
	myurl := "http://localhost:8000/post"
	PerformPostRequest(myurl)
}

func PerformPostRequest(url string) {
	//fake json payload

	requestBody := strings.NewReader(`
	{
		"CourseName":"Golang",
		"Version":"1.0",
		"Date":"29-11-2023",
		"Year":"2023"
	}
	
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error encounter")
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("The content is : ", string(content))
}
