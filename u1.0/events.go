package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Id          int         `json:"id"`
	Subject     string      `json:"subject"`
	Time        time.Time   `json:"time"`
	Venue       string      `json:"venue"`
	Agenda      string      `json:"agenda"`
	Salutation  string      `json:"salutation"`
	Item13      bool        `json:"item13"`
	Signed      string      `json:"signed"`
	SocialGroup string      `json:"socialgroup"`
	EventSeries EventSeries `json:"eventseries"`
}

func (event *Event) Fill(idstr string) (err error) {
	query := "SELECT id, subject, time, venue, agenda, salutation, item13, signed, socialgroup, eventseries_id FROM events where id=$1;"

	err = Db.QueryRow(query, idstr).Scan(&event.Id, &event.Subject, &event.Time, &event.Venue, &event.Agenda, &event.Salutation, &event.Item13, &event.Signed, &event.SocialGroup, &event.EventSeries.Id)

	return
}

func (event *Event) Create() (err error) {
	query := "INSERT INTO events (subject, time, venue, agenda, salutation, item13, signed,  socialgroup, eventseries_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"

	stmt, err := Db.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing statement : ", err.Error())
		return
	}

	_, err = stmt.Exec(event.Subject, event.Time, event.Venue, event.Agenda, event.Salutation, event.Item13, event.Signed, event.SocialGroup, event.EventSeries.Id)

	return
}

func (event *Event) Update() (err error) {
	query := "UPDATE events SET subject=?, time=?, venue=?, agenda=?, salutation=?, item13=?, signed=?,  socialgroup=?, eventseries_id=? WHERE id=?;"

	stmt, err := Db.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing statement : ", err.Error())
		return
	}

	_, err = stmt.Exec(event.Subject, event.Time, event.Venue, event.Agenda, event.Salutation, event.Item13, event.Signed, event.SocialGroup, event.EventSeries.Id, event.Id)

	return
}

func (event *Event) All() ([]Event, error) {
	query := "SELECT id, subject, time, venue, agenda, salutation, item13, signed, socialgroup, eventseries_id FROM events;"

	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}

	events := []Event{}
	for rows.Next() {
		event := Event{}

		err := rows.Scan(&event.Id, &event.Subject, &event.Time, &event.Venue, &event.Agenda, &event.Salutation, &event.Item13, &event.Signed, &event.SocialGroup, &event.EventSeries.Id)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func toJson(events []Event) (string, error) {
	out, err := json.MarshalIndent(events, "", "\n")
	return string(out), err
}

func eventsBelongingToSeries(eventSeriesId string) ([]Event, error) {
	query := "SELECT id, subject, time, venue, agenda, salutation, item13, signed, socialgroup, eventseries_id FROM events WHERE eventseries_id=?;"

	rows, err := Db.Query(query, eventSeriesId)
	if err != nil {
		return nil, err
	}

	events := []Event{}
	for rows.Next() {
		event := Event{}

		err := rows.Scan(&event.Id, &event.Subject, &event.Time, &event.Venue, &event.Agenda, &event.Salutation, &event.Item13, &event.Signed, &event.SocialGroup, &event.EventSeries.Id)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}