package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("Welcome to time Study of Golang")
	presentTime := time.Now()

	fmt.Println(presentTime)

	//chnaging or format the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:06 Monday"))

	// Note: we can create time and date i.e is the past date  all will be any thing but the month will be time.month that means the month which we are required
	createDate:= time.Date(2023, time.April,10,23,23,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println("The created Date and time is:",createDate)
}