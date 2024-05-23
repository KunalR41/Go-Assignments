package handlers

import (
	"MyHotel/Hotels/models"
	"MyHotel/Hotels/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	Repo *repository.HotelRepository
}

func (h *HotelHandler) GetAllHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := h.Repo.GetAll()
	fmt.Fprintln(w, "~~~ Welome To MyHotel API.....!!! ~~~")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hotels)
}

func (h *HotelHandler) GetHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid hotel ID", http.StatusBadRequest)
		return
	}

	hotel, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(hotel)
}

func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.Create(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hotel)
}

func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid hotel ID", http.StatusBadRequest)
		return
	}

	var hotel models.Hotel
	err = json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hotel.ID = id

	err = h.Repo.Update(&hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hotel)
}

func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid hotel ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Hotel deleted successfully"})
}
