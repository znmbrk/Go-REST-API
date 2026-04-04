package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Lessons int    `json:"lessons"`
}

var courses = []Course{
	{ID: "MIM-101", Title: "Modern Incident Management", Lessons: 3},
	{ID: "ONC", Title: "On call basics", Lessons: 5},
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(courses)
	} else if r.Method == "POST" {
		var course = Course{}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if course.ID == "" {
			http.Error(w, "ID cannot be empty", http.StatusBadRequest)
			return
		}
		courses = append(courses, course)
		json.NewEncoder(w).Encode(course)
	} else {
		http.Error(w, "We only support Get and Post methods rightnow", http.StatusMethodNotAllowed)
	}
}

func specificCourseHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[9:]
	matchIndex := courseLoopHelper(id)

	if r.Method == "GET" {
		if matchIndex >= 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(courses[matchIndex])
			return
		}
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	} else if r.Method == "PUT" {
		if matchIndex >= 0 {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewDecoder(r.Body).Decode(&courses[matchIndex])
			if err != nil {
				http.Error(w, "Error", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(courses[matchIndex])
			return
		}
		http.Error(w, "No ID was matched", http.StatusNotFound)
	} else if r.Method == "DELETE" {
		if matchIndex >= 0 {
			toDelete := courses[matchIndex]
			w.Header().Set("Content-Type", "application/json")
			courses = append(courses[:matchIndex], courses[matchIndex+1:]...)
			json.NewEncoder(w).Encode(toDelete)
			return
		}
		http.Error(w, "No ID found", http.StatusNotFound)

	} else {
		http.Error(w, "Only GET method is supported for this", http.StatusMethodNotAllowed)
	}
}

func courseLoopHelper(id string) int {
	for i, course := range courses {
		if course.ID == id {
			return i
		}
	}
	return -1
}
