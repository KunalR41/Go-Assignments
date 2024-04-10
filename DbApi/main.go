package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

var db *sql.DB

func main() {
	// Connecting to the database
	var err error
	db, err = sql.Open("mysql", "root:root@123@tcp(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	//http://localhost:9000
	r.HandleFunc("/", getMessage).Methods("GET")
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Println("Server is Running on port 9000")
	// Starting server
	log.Fatal(http.ListenAndServe(":9000", r))
}

// /-welcome message
func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME....Please Enter Valid Endpoint....!!!")

}
func getStudents(w http.ResponseWriter, r *http.Request) {
	// "/students"-Fetching all students Info
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Marks); err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	// Marshal and send response
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	// "/students/{id}"- Fetching single student Info  by ID
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECT * FROM students WHERE id = ?", id)

	var student Student
	if err := row.Scan(&student.ID, &student.Name, &student.Marks); err != nil {
		log.Fatal(err)
	}

	// Marshal and send response
	json.NewEncoder(w).Encode(student)
}

func createStudent(w http.ResponseWriter, r *http.Request) {

	var student Student
	// Unmarshal getting json data from request
	json.NewDecoder(r.Body).Decode(&student)

	// "/students"-creating new student
	_, err := db.Exec("INSERT INTO students(name, marks) VALUES(?, ?)", student.Name, student.Marks)
	if err != nil {
		log.Fatal(err)
	}

}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var student Student
	// Unmarshal getting json data from request
	json.NewDecoder(r.Body).Decode(&student)

	// "/students/{id}"-Update student by ID
	_, err = db.Exec("UPDATE students SET name = ?, marks = ? WHERE id = ?", student.Name, student.Marks, id)
	if err != nil {
		log.Fatal(err)
	}

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	// "/students/{id}"-Delete student by ID
	_, err = db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

}
