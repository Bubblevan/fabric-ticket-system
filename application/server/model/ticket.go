package model

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	EventName string
	Venue     string
	EventDate time.Time
	Price     float64
	Num       uint
}
