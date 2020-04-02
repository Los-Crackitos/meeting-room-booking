package dbmodels

import "time"

type Booking struct {
	ID        int `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time	`json:"end_date"`
	RoomID    int `json:"room_id"`
	UserID    int `json:"user_id"`
}
