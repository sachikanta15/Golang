package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func main() {
	fmt.Println("Handling Routes in Golang by using Gorilla Mux")
	// how to initialize routes
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{FullName: "Sachiknata", Website: "sachiknata.dev"}})

	// routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/welcome",welcome).Methods("GET")
	// r.HandleFunc("/welcome{id}",welcome).Methods("GET")
	

	// listening to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to routing</h1>"))
}
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome </h1>"))
}

// func welcome(w http.ResponseWriter, r *http.Request) {
	// params:= mux.var(r)
	// something = params["id"]  ---- this id synatx will be same thing as of route --// r.HandleFunc("/welcome{id}",welcome).Methods("GET") i.e is in welcome{id}
// 	w.Write([]byte("<h1>Welcome </h1>"))
// }
