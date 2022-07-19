package main

import (
	"fmt"
	"net/http"
)

func main() {

	connectDatabse()
	fmt.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Start)
	http.HandleFunc("/showauthors", getUsersFromDB)
	http.HandleFunc("/showbooks", getBooksFromDB)
	http.HandleFunc("/search", search)
	http.HandleFunc("/insertNew", InsertNew)
	//	http.HandleFunc("/update", Update)
	//	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
