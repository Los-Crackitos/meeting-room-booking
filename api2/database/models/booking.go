package models

import "time"

type Booking struct {
	ID        uint
	StartDate time.Time
	EndDate   time.Time
	RoomID    uint
	UserID    uint
}
