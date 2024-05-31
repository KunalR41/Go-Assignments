package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"tourism/database"
	"tourism/package-service/models"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	db = database.Connect()
}

func GetPackages(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, destination_id, name, price FROM packages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	packages := []models.Package{}
	for rows.Next() {
		var pkg models.Package
		if err := rows.Scan(&pkg.ID, &pkg.DestinationID, &pkg.Name, &pkg.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		packages = append(packages, pkg)
	}
	json.NewEncoder(w).Encode(packages)
}

func CreatePackage(w http.ResponseWriter, r *http.Request) {
	var pkg models.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO packages (destination_id, name, price) VALUES (?, ?, ?)",
		pkg.DestinationID, pkg.Name, pkg.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	pkg.ID = int(id)
	json.NewEncoder(w).Encode(pkg)
}

func GetPackage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var pkg models.Package
	if err := db.QueryRow("SELECT id, destination_id, name, price FROM packages WHERE id = ?", id).
		Scan(&pkg.ID, &pkg.DestinationID, &pkg.Name, &pkg.Price); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pkg)
}

func UpdatePackage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var pkg models.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE packages SET destination_id = ?, name = ?, price = ? WHERE id = ?",
		pkg.DestinationID, pkg.Name, pkg.Price, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pkg.ID = id
	json.NewEncoder(w).Encode(pkg)
}

func DeletePackage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if _, err := db.Exec("DELETE FROM packages WHERE id = ?", id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
