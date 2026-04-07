package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the incident Academy API!")
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=zainmobarik dbname=course_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/courses", courseGet).Methods("GET")
	r.HandleFunc("/courses", coursePost).Methods("POST")
	r.HandleFunc("/courses/{ID}", specificCourseGet).Methods("GET")
	r.HandleFunc("/courses/{ID}", specificCoursePut).Methods("PUT")
	r.HandleFunc("/courses/{ID}", specificCourseDelete).Methods("DELETE")

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
