package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Maps in Golang")

	language := make(map[string]string) // declariation of map in key value pair --- here [string] i.e key
	// by this we can add value int teh map as much we want
	language["JS"] = "Javascript"
	language["GO"] = "Golang"
	language["TS"] = "Typescript"
	language["PY"] = "Python"
	fmt.Println("printing all the values in Languages", language)
	// printing one value from the map. it can be accessed through key
	fmt.Println("Printing one value from the map", language["TS"])

	// Deleting the values from the map
	deleted_value := language["PY"]
	delete(language, "PY")
	fmt.Println("Printig the languages after deleting", language)
	fmt.Println("Printig the deleted value after deleting from the language map", deleted_value)

	// looping through the map
	for key, value := range language { //key value from the range of length as language
		fmt.Printf("for key %v, is value is %v \n", key, value) //key value with respective value to it
	}
}
