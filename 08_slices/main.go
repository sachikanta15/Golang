package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Slices in Golang")
	var fruitList = []string{"Apple", "Mango", "Pineapple"} //same as array but here we don't declare the length in it
	fmt.Println("fruitlist before append", fruitList)
	fruitList = append(fruitList, "Banana", "Jackfruit", "peach")
	fmt.Println("fruitlist after append", fruitList)

	// how we slice from array is same here but the syntax is different
	fruitList = append(fruitList[1:]) //this means the 1st index will be slices
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) //this menas the last index is exculsive  it will omit 0th index, print from 1,2  and 3rd index will be exculsive
	fmt.Println(fruitList)

	//how to remove the values from the slices in index based, to get we will use smae append menthod but in different apporach

	var courses = []string{"reactjs", "nextjs", "Golang", "Javascript"}
	fmt.Println("Printing the courses", courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

// slices are same like Arraylist but through append we will add new value into it
