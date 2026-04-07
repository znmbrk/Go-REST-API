package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Lessons int    `json:"lessons"`
}

func courseGet(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from courses;")

	if err != nil {
		http.Error(w, "Failed to query courses from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var course Course
	var courseSlice = []Course{}

	for rows.Next() {
		err := rows.Scan(&course.ID, &course.Title, &course.Lessons)
		if err != nil {
			http.Error(w, "Failed to scan course row", http.StatusInternalServerError)
			return
		}
		courseSlice = append(courseSlice, course)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courseSlice)
}

func coursePost(w http.ResponseWriter, r *http.Request) {
	var course = Course{}
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if course.ID == "" {
		http.Error(w, "ID cannot be empty", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("insert into courses (id, title, lessons) values ($1, $2, $3)", course.ID, course.Title, course.Lessons)
	if err != nil {
		http.Error(w, "POST method failed to insert course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func specificCourseGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var course = Course{}

	err := db.QueryRow("select * from courses where id = $1", id).Scan(&course.ID, &course.Title, &course.Lessons)
	if err != nil {
		http.Error(w, "Failed to query course from database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func specificCoursePut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var course = Course{}
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	_, err = db.Exec("update courses set lessons = $1, title = $2 where id = $3", course.Lessons, course.Title, id)

	if err != nil {
		http.Error(w, "PUT method failed to update course", http.StatusInternalServerError)
		return
	}
	course.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func specificCourseDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	_, err := db.Exec("delete from courses where id = $1", id)
	if err != nil {
		http.Error(w, "DELETE method failed to delete course", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"deleted": id})
}
