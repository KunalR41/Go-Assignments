package repository

import (
	"MyHotel/Hotels/models"
	"database/sql"
)

type HotelRepository struct {
	DB *sql.DB
}

func (repo *HotelRepository) GetAll() ([]models.Hotel, error) {
	rows, err := repo.DB.Query("SELECT id, name, address, rating FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Rating); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (repo *HotelRepository) GetByID(id int) (*models.Hotel, error) {
	var hotel models.Hotel
	err := repo.DB.QueryRow("SELECT id, name, address, rating FROM hotels WHERE id = ?", id).Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Rating)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (repo *HotelRepository) Create(hotel *models.Hotel) error {
	result, err := repo.DB.Exec("INSERT INTO hotels (name, address, rating) VALUES (?, ?, ?)", hotel.Name, hotel.Address, hotel.Rating)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	hotel.ID = int(id)
	return nil
}

func (repo *HotelRepository) Update(hotel *models.Hotel) error {
	_, err := repo.DB.Exec("UPDATE hotels SET name = ?, address = ?, rating = ? WHERE id = ?", hotel.Name, hotel.Address, hotel.Rating, hotel.ID)
	return err
}

func (repo *HotelRepository) Delete(id int) error {
	_, err := repo.DB.Exec("DELETE FROM hotels WHERE id = ?", id)
	return err
}
