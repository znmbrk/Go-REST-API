package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/courses", courseHandler)
	http.HandleFunc("/courses/", specificCourseHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
