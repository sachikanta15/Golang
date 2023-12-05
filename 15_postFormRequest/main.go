package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the Golang Post form")

	PerformPostFormRequest()
}
func PerformPostFormRequest() {
	myurl := "http://localhost:8000/postform"
	//formdata

	data := url.Values{}
	data.Add("firstname", "sachiknata")
	data.Add("lastname", "Raul")
	data.Add("email", "sachikanta@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("The content is : ", string(content))

}
