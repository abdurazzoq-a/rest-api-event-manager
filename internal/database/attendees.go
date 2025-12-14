package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeesModel struct {
	DB *sql.DB
}


type Attendee struct {
	Id      int `json:"id"`
	UserId    int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttendeesModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "INSERT INTO attendees(event_id, user_id) VALUES ($1, $2) RETURNING id"

	err := m.DB.QueryRowContext(ctx, query).Scan(&attendee.Id)

	if err != nil {
		return nil, err
	}

	return attendee, nil
}

func (m *AttendeesModel) GetByEventAndAttendee(eventId int, userId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()


	query := "SELECT * FROM attendees WHERE event_id = $1 AND user_id = $2"

	var attendee Attendee
	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(&attendee.Id, &attendee.EventId, &attendee.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &attendee, nil
}



func (m * AttendeesModel) GetAttendeesByEvent(eventId int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT u.id, u.name, u.email 
		FROM users as u
		JOIN attendees as a 
		ON u.id = a.user_id
		WHERE a.event_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, eventId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User


	for rows.Next() {
		var user User 

		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		
		if err == nil {
			return nil, err
		}

		users = append(users, &user)
	}


	return users, nil

}
