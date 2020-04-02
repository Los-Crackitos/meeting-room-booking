package models

import "api-go/database/dbmodels"

type User struct {
	Params dbmodels.User `json:"params"`
}

type Room struct {
	Params dbmodels.Room `json:"params"`
}

type Booking struct {
	Params dbmodels.Booking `json:"params"`
}
