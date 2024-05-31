package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"tourism/booking-service/models"
	"tourism/database"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	db = database.Connect()
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, package_id, user_name, user_email, booking_date FROM bookings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bookings := []models.Booking{}
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.PackageID, &booking.UserName, &booking.UserEmail, &booking.BookingDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, booking)
	}
	json.NewEncoder(w).Encode(bookings)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO bookings (package_id, user_name, user_email, booking_date) VALUES (?, ?, ?, ?)",
		booking.PackageID, booking.UserName, booking.UserEmail, booking.BookingDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	booking.ID = int(id)
	json.NewEncoder(w).Encode(booking)
}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var booking models.Booking
	if err := db.QueryRow("SELECT id, package_id, user_name, user_email, booking_date FROM bookings WHERE id = ?", id).
		Scan(&booking.ID, &booking.PackageID, &booking.UserName, &booking.UserEmail, &booking.BookingDate); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(booking)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE bookings SET package_id = ?, user_name = ?, user_email = ?, booking_date = ? WHERE id = ?",
		booking.PackageID, booking.UserName, booking.UserEmail, booking.BookingDate, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	booking.ID = id
	json.NewEncoder(w).Encode(booking)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if _, err := db.Exec("DELETE FROM bookings WHERE id = ?", id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
