package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the Golang Handling Web Request")

	url := "https://sachikantaonemoney.blogspot.com/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("printing the response\n", response)
	fmt.Printf("Printig the type of response %T", response)
	defer response.Body.Close() //Note: IMP-- caller's responsibility is to close the connection

	// to read the the response.body use ioutil
}
