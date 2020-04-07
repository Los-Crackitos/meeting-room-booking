package dbhandlers

import (
	"api-go/database"
	"api-go/database/dbmodels"
)

func CreateUser(user *dbmodels.User) error {

	if err := database.Db.Insert(user); err != nil {
		return err
	}

	return database.Db.Select(user)
}

func GetUser(user *dbmodels.User) error {
	return database.Db.Select(user)
}

func GetAllUsers() ([]dbmodels.User, error) {
	var users []dbmodels.User

	if err := database.Db.Model(&users).Order("id ASC").Select(); err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(user *dbmodels.User) error {
	return database.Db.Delete(user)
}

func UpdateUser(user *dbmodels.User) error {
	return database.Db.Update(user)
}
