package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Golang Web request")
	myurl := "http://localhost:8000/get"
	PerformGetRequest(myurl)
}

func PerformGetRequest(url string) {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// to read from thebody ioutil

	content, _ := io.ReadAll(response.Body)

	fmt.Println("The content is : ", string(content))

	//the above code can be done through a method which is strings.Builder ---

	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)

	fmt.Println("by this the data which contains in string will not print ", bytecount)
	fmt.Println("But by this the data which are under it will print\n", responseString.String())
}
