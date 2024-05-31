package models

type Booking struct {
	ID          int    `json:"id"`
	PackageID   int    `json:"package_id"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	BookingDate string `json:"booking_date"`
}
