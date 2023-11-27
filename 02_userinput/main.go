package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the userInput"
	fmt.Println(welcome)

	//Here we willlearn how to take input from teh user
	// bufio is the package which helps to take the input from the user --> bbufio contains an NewReadrer property and from this we will get help

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of the tuts:")

	// What the reader reads i want to store in some variable

	// comma ok || err ok ----> in go we don't had try catch , to deal with it . As go expect to user deal with true/false or in variable

	input, _ := reader.ReadString('\n') //what does this line do ? input,err(or it can be _, "underscore") it is something like try catch , here if there is no input as err and

	// (/n) means till the line break it will read

	fmt.Println("Thanks for rating,", input)  
	//but through this taking input it will store as string , we cna only do string operation in this

}
