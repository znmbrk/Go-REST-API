package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the incident Academy API!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/courses", courseHandler)
	http.HandleFunc("/courses/", specificCourseHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
