package main

import "fmt"

func main() {
	fmt.Println("Welcome to teh golang loops")

	// days := []string{"sunday", "wednesday", "friday", "saturday"}
	// // fmt.Println("printings the days", days)

	// //printing through for looop
	// fmt.Println("Printing through loops")
	// for i := 0; i < len(days); i++ {

	// 	fmt.Println(days[i])
	// }

	// //printing through range

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// //for each loop in golang
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and the days is %v\n", index, day)
	// }
	// //if we want only days not want inde than we can go with ok comma
	// for _, day := range days {
	// 	fmt.Printf("The days is %v\n", day)
	// }

	// Break and continue in Golang

	rougevalue := 1
	for rougevalue < 10 {

		if rougevalue==3{
			goto lco
		}
		if rougevalue == 5 {
			rougevalue++
			continue
		}
		// if rougevalue == 6 {
		// 	break
		// }
		fmt.Println(rougevalue)
		rougevalue++
	}


	//another thing is goto statement here we had to provide label and when the condition satisfy than we can directly jump into go to and when goto trigers than remain will
	// not be run
	lco:
	fmt.Println("Jumping into sachikanta.com as rougevalue satisfy at 2")
}
