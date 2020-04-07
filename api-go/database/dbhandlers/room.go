package dbhandlers

import (
	"api-go/database"
	"api-go/database/dbmodels"
)

func CreateRoom(room *dbmodels.Room) error {
	if err := database.Db.Insert(room); err != nil {
		return err
	}

	return database.Db.Select(room)
}

func GetRoom(room *dbmodels.Room) error {
	return database.Db.Select(room)
}

func GetAllRooms() ([]dbmodels.Room, error) {
	var rooms []dbmodels.Room

	if err := database.Db.Model(&rooms).Order("id ASC").Select(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func DeleteRoom(room *dbmodels.Room) error {
	return database.Db.Delete(room)
}

func UpdateRoom(room *dbmodels.Room) error {
	return database.Db.Update(room)
}
