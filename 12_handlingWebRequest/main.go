package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the Golang Handling Web Request")

	// url := "https://sachikantaonemoney.blogspot.com/"

	// response, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("printing the response\n", response)
	// fmt.Printf("Printig the type of response %T", response)
	// defer response.Body.Close() //Note: IMP-- caller's responsibility is to close the connection

	// to read the the response.body use ioutil

	//destrucing the url and how to get the data form the url in golang   that is parsing from the url

	myUrl := "https://sachikantaonemoney.blogspot.com/:3000/learn?coursename=reactjs&paymentid=hvsdcjhvgdcj"

	// parsing

	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("query params are : ", qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}
}
