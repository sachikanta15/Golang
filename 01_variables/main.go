package main

import "fmt"

// Println is a simple and concise way to print messages with a new line at the
// end, while Printf allows you to specify the format of the printed values.

// trying using := operator outside the func, it will provide the syntax error
// for this we need to declare in natural manner ho we declare
// var number int =5000;

// fmt.Println(number)


func main() {
	var username string = "sachikanta"
	fmt.Printf("Variable is of type: %T \n", username)
	fmt.Println("Variables")

	var isBoolean bool = true
	fmt.Printf("Variable is of type: %T \n", isBoolean)
	fmt.Println("isBoolean")

	// implicit type here we don't declare any data type go automatically decides as per input

	var website = "sachikantaraul.com"
	fmt.Println(website)

	// no var style
	numberStudents:=3000
	fmt.Println(numberStudents)

}
