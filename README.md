# REST API in Go

## What is this? 
- A REST API built in Go that allows the user to manage a set of courses

## What can it do? 
- CRUD operations for managing courses
- GET, POST, PUT, DELETE

## How do you run it? 
- Install go @ https://go.dev/dl/ 
- "go run *.go" - this starts the server on 8080
- Use curl to hit the endpoints /courses, /courses/, /health

## Project structure
- main.go - main function with home handler
- courses.go - Course struct and course handler functions
- health.go - Health struct and health handler functions
- Files grouped by domain 

## What you learned building it?
- How APIs and JSON work. Reading request bodies and understanding how communication between the client and server works
- How to extract IDs from the URL to handle specific requests