package models

type Package struct {
	ID            int     `json:"id"`
	DestinationID int     `json:"destination_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
}
