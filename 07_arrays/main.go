package main

import "fmt"

func main() {
	fmt.Println("Welcome to thr Arrays")
var fruitList [4] string // declartion of arrays
fruitList[0]="Apple"
fruitList[1]="Banana"
fmt.Println("Fruit list is:",fruitList)
fmt.Println("Length of teh fruitlist is:",len(fruitList))   //finding the length of the arrays



//declartion of arrays and assigning values in it directly

var vegList =[3] string {"Potato","Cucmber","Sweet Potato"}  //initlize and put the values

fmt.Println("VegList is:",vegList)
fmt.Println("Length of the vegList:",len(vegList))
}
