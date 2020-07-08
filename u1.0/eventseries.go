package main

type EventSeries struct {
	Id   int
	Name string
}

var (
	DefaultEventSeries EventSeries
)

func (evntsrs *EventSeries) Create() error {
	query := "INSERT INTO eventseries (name) VALUES ($2);"

	_, err := Db.Exec(query, evntsrs.Name)

	return err
}

func (evntsrs *EventSeries) All() ([]EventSeries, error) {
	query := "SELECT id, name FROM eventseries;"

	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}

	all := []EventSeries{}
	for rows.Next() {
		evntsrs := EventSeries{}

		err := rows.Scan(&evntsrs.Id, &evntsrs.Name)
		if err != nil {
			return nil, err
		}

		all = append(all, evntsrs)
	}

	return all, nil
}

func GetEventSeries(id string) (EventSeries, error) {
	query := "SELECT id, name FROM eventseries WHERE id=?;"

	evntsrs := EventSeries{}
	err := Db.QueryRow(query, id).Scan(&evntsrs.Id, &evntsrs.Name)

	return evntsrs, err
}
