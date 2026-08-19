package model

import (
	"time"

	db "github.com/HirunikaGunathunga/GO_REST/database"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int       `json:"userId"`
}

func (e Event) Save() error {
	//later DB
	//events = append(events, e)
	query := `INSERT INTO EVENT_TAB (name, description, location, dateTime, userId) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM EVENT_TAB"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM EVENT_TAB WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	//defer rows.Close()
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Event) UpdateEventById() error {
	query := `UPDATE EVENT_TAB 
	SET name = ?, description = ?, location = ?, dateTime = ?, userId = ? 
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)

	return err
}

func (e Event) DeleteEvent() error {
	query := "DELETE FROM EVENT_TAB WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
