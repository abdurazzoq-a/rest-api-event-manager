package database

import "database/sql"



type AttendeesModel struct {
	DB *sql.DB
}


type Attendee struct {
	Id int `json:"id"`
	User int `json:"userId"`
	EventId int `json:"eventId"`
}