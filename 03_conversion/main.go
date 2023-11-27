package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please Rate our App from 1 to 5")


	reader:=bufio.NewReader(os.Stdin)   
	//this reader don't do anything this is the only takes refrence from this and use the ReadString(it can be anything , 
	// there anre many data types to read) method to read

	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for Rating",input)

	// numrating,err:= strconv.ParseFloat(input,64) 
	// above will not work bcz while entering the value , we enter 5 with space i.e enter t, so that we need to trim the space for this we need to use package strings.TrimSpace()
	numrating,err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating", numrating+1)
	}
}