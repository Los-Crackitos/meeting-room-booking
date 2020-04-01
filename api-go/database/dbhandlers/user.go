package dbhandlers

import (
	"api-go/database"
	"api-go/database/dbmodels"
)

func CreateUser(user *dbmodels.User) error {
	err := database.Db.Insert(user)
	if err != nil {
		return err
	}

	return database.Db.Select(user)
}

func GetUser(user *dbmodels.User) error {
	return database.Db.Select(user)
}
