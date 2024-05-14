package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Menu struct {
	ID    int     `json:"id"`
	Dish  string  `json:"dish"`
	Price float64 `json:"price"`
}

func GetMenus(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM menu")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		if err := rows.Scan(&menu.ID, &menu.Dish, &menu.Price); err != nil {
			log.Fatal(err)
		}
		menus = append(menus, menu)
	}

	// Marshal and send response
	json.NewEncoder(w).Encode(menus)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECT * FROM menu WHERE id=?", id)

	var menu Menu
	if err := row.Scan(&menu.ID, &menu.Dish, &menu.Price); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(menu)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu Menu
	json.NewDecoder(r.Body).Decode(&menu)

	_, err := db.Exec("INSERT INTO menu(Dish,Price) VALUES(?,?)", menu.Dish, menu.Price)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var menu Menu
	json.NewDecoder(r.Body).Decode(&menu)

	_, err = db.Exec("UPDATE menu SET dish=?,price=? WHERE id=?", menu.Dish, menu.Price, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM menu WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}
