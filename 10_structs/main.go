package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Golang Structs")
	sachiknata := User{"Sachikanta Raul", "sachikanta@google.com", true, 23}
	fmt.Println("Printing user details:", sachiknata)
	fmt.Printf("%+v\n", sachiknata)
	fmt.Printf("Name is %v and email is %v", sachiknata.Name, sachiknata.Email)

}

//declartion of structs ,-----> these are basically like objects

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
