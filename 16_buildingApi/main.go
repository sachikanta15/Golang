package main

import (
	"encoding/json"
	"net/http"
)

//creating models for course--------  models same as inertface which we are creating in typescript

type Course struct {
	CourseId    string  `json:"coursename`
	CourseName  string  `json:coursename`
	CoursePrice string  `json:courseprice`
	Author      *Author `json:author`
}

type Author struct {
	FullName string `json:fullname`
	Website  string `json:website`
}

//creating fake db as we are not connected with db

var course []Course //creating through slice as we don't no what will be the data

//creating middlewares , same concept as nodejs ----------- it should be seperate file
//The middlewares are basically used when ever we want to manuplate in the request and resopnse , their we will use middlewares
//use case let we create our website and check if user is sigin than go forward or else redirect to signin----- something like that

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

//what are controlllers --- if any requeest comes than the handling of that situation is controlled by controllers

//creating an api

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Golang Api learning</h1?"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") // it always in key value pair
	json.NewEncoder(w).Encode(course)                  // this is the synatx when ever we want to pass the json when ever the api is called
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	
}
func main() {

}
