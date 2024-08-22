package model

import (
	"time"
)

type Order struct {
	ID         string    `json:"id"`
	UserID     uint      `json:"userID"`
	TicketID   uint      `json:"ticketID"`
	Num        int       `json:"num"`
	TotalPrice float64   `json:"totalPrice"`
	OrderDate  time.Time `json:"orderDate"`
}
