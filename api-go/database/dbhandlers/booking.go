package dbhandlers

import (
	"api-go/database"
	"api-go/database/dbmodels"
)

func CreateBooking(booking *dbmodels.Booking) error {
	err := database.Db.Insert(booking)

	if err != nil {
		return err
	}
	return database.Db.Select(booking)
}

func GetAllBookings() ([]dbmodels.Booking, error) {
	var bookings []dbmodels.Booking

	if err := database.Db.Model(&bookings).Order("start_date ASC").Select(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func GetBooking(booking *dbmodels.Booking) error {
	return database.Db.Select(booking)
}

func DeleteBooking(booking *dbmodels.Booking) error {
	return database.Db.Delete(booking)
}

func UpdateBooking(booking *dbmodels.Booking) error {
	return database.Db.Update(booking)
}
