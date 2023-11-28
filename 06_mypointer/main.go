package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Pointers")
}



// what are pointer?
// sometime when we pass the value into any func or nay object than mostly the copy of that memory is pas on but if we pass through pointers than it will be suerity that 
// the actual address will be pass. As we are pssing the actual memory address than it is 100 % garantee that memory adress will be pass

// How we will create pointers--------> var sum int =2 (this is the way of declareing the varibale ) but in pointer var ptr *int ------(here we didn't assign any value into it)

// NOTE: whne we are craetig the pointer than we will us the * symbol just like the above but if we are creating pointer where we are assigning the refernce

// for example :
//   ptr= &myNumber