package model

import "time"

type Book struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Edition int    `json:"edition,omitempty"`
}

type PickupSchedule struct {
	BookTitle  string    `json:"book_title"`
	UserName   string    `json:"user_name"`
	PickupTime time.Time `json:"pickup_time"`
}
