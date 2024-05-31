package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tourism/database"
	"tourism/destination-service/models"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	db = database.Connect()
}

func GetDestinations(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, location FROM destinations")
	fmt.Fprintln(w, "~~~ Welome To Destinations API.....!!! ~~~")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	destinations := []models.Destination{}
	for rows.Next() {
		var destination models.Destination
		if err := rows.Scan(&destination.ID, &destination.Name, &destination.Description, &destination.Location); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		destinations = append(destinations, destination)
	}
	json.NewEncoder(w).Encode(destinations)
}

func CreateDestination(w http.ResponseWriter, r *http.Request) {

	var destination models.Destination
	if err := json.NewDecoder(r.Body).Decode(&destination); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO destinations (name, description, location) VALUES (?, ?, ?)",
		destination.Name, destination.Description, destination.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	destination.ID = int(id)
	json.NewEncoder(w).Encode(destination)
}

func GetDestination(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var destination models.Destination
	if err := db.QueryRow("SELECT id, name, description, location FROM destinations WHERE id = ?", id).
		Scan(&destination.ID, &destination.Name, &destination.Description, &destination.Location); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(destination)
}

func UpdateDestination(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var destination models.Destination
	if err := json.NewDecoder(r.Body).Decode(&destination); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE destinations SET name = ?, description = ?, location = ? WHERE id = ?",
		destination.Name, destination.Description, destination.Location, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	destination.ID = id
	json.NewEncoder(w).Encode(destination)
}

func DeleteDestination(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if _, err := db.Exec("DELETE FROM destinations WHERE id = ?", id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
