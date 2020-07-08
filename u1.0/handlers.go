package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"
)

var (
	dates    []string
	subjects []string
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/index.html",
	}
	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}
	t.ExecuteTemplate(w, "layout", "")
}

func outlineReq(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/outlineReq.html",
	}
	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}
	t.ExecuteTemplate(w, "layout", "")
}

func extractOutlineReq(r *http.Request) (startDate, endDate time.Time, interval int, err error) {

	dateFormat := "2006-01-02"

	startDate, err = time.Parse(dateFormat, r.FormValue("startDate"))
	if err != nil {
		err = fmt.Errorf("Error parsing start date : %s", err.Error())
		return
	}

	endDate, err = time.Parse(dateFormat, r.FormValue("endDate"))
	if err != nil {
		err = fmt.Errorf("Error parsing end dat : %s", err.Error())
		return
	}

	interval, err = strconv.Atoi(r.FormValue("dysBtwnEvents"))
	if err != nil {
		err = fmt.Errorf("Error parsing interval : %s", err.Error())
		return
	}

	return
}

func outline(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/createOutline.html",
	}

	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}

	startDate, endDate, interval, err := extractOutlineReq(r)
	if err != nil {
		fmt.Println("Error extracting outline requirement: ", err.Error())
		return
	}

	outlineDates := []time.Time{}
	for tempDate := startDate; !tempDate.After(endDate); tempDate = tempDate.AddDate(0, 0, interval) {
		outlineDates = append(outlineDates, tempDate)
	}

	t.ExecuteTemplate(w, "layout", outlineDates)
	return
}

func saveOutline(r *http.Request) {
	r.ParseForm()
	dates = r.Form["date"]
	subjects = r.Form["subject"]
	return
}

func defaults(w http.ResponseWriter, r *http.Request) {
	saveOutline(r)

	tmplFiles := []string{
		"./public/html/layout.html",
		"./public/html/defaults.html",
		"./public/html/eventseries.html",
	}

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		fmt.Println("Error parsing defaults template files: ", err.Error())
		return
	}

	evntsrs := EventSeries{}
	all, err := evntsrs.All()
	if err != nil {
		fmt.Println("Error getting all event series from db: ", err.Error())
		return
	}

	t.ExecuteTemplate(w, "layout", all)
}

func eventseries(w http.ResponseWriter, r *http.Request) {
	DefaultEventSeries.Name = r.FormValue("name")
	err := DefaultEventSeries.Create()
	if err != nil {
		fmt.Println("Error creating new event series:", err.Error())
		return
	}

	tmpFiles := []string{
		"./public/html/eventseries.html",
	}
	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}

	es, err := DefaultEventSeries.All()
	if err != nil {
		fmt.Println("Error event series from db:", err.Error())
		return
	}

	t.ExecuteTemplate(w, "eventseries", es)
	return
}

func extractEventData(r *http.Request) (event Event, err error) {
	event = Event{
		Subject:     r.FormValue("subject"),
		Venue:       r.FormValue("venue"),
		Agenda:      r.FormValue("agenda"),
		Salutation:  r.FormValue("salutation"),
		Signed:      r.FormValue("signed"),
		SocialGroup: r.FormValue("socialgroup"),
	}

	event.Item13, err = strconv.ParseBool(r.FormValue("item13"))
	if err != nil {
		err = fmt.Errorf("Error parsing item 13 : %s", err.Error())
		return
	}

	event.EventSeries.Id, err = strconv.Atoi(r.FormValue("eventseries"))
	if err != nil {
		err = fmt.Errorf("Error parsing eventseries id : %s", err.Error())
		return
	}

	return
}

func extractEventTime(dateStr, timeStr string) (dateTime time.Time, err error) {
	dateTimeStr := dateStr + " " + timeStr
	layout := "2006-01-02 15:04"
	dateTime, err = time.Parse(layout, dateTimeStr)
	return
}

func createOutline(w http.ResponseWriter, r *http.Request) {

	defaultEvent, err := extractEventData(r)
	if err != nil {
		fmt.Println("Error extracting default values: ", err.Error())
		return
	}

	for i, val := range subjects {
		tempEvent := defaultEvent
		tempEvent.Subject = val

		tempEvent.Time, err = extractEventTime(dates[i], r.FormValue("time"))
		if err != nil {
			fmt.Println("Error parsing time :", err.Error())
			return
		}

		tempEvent.Create()
	}

	allEvents(w, r)
	return
}

func allEvents(w http.ResponseWriter, r *http.Request) {

	event := Event{}
	allEvents, err := event.All()
	if err != nil {
		fmt.Println("Error getting all events from db : ", err.Error())
		return
	}

	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/allevents.html",
	}

	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}

	t.ExecuteTemplate(w, "layout", allEvents)
	return
}

func getEditEventData(r *http.Request) (event Event, allEvntsrs []EventSeries, err error) {
	idstr := path.Base(r.URL.Path)
	_, err = strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error, bad id provided : ", err.Error())
		return
	}

	err = event.Fill(idstr)
	if err != nil {
		err = fmt.Errorf("Error reading event from db : %s", err.Error())
		return
	}

	allEvntsrs, err = DefaultEventSeries.All()
	if err != nil {
		err = fmt.Errorf("Error reading all event series from db : %s", err.Error())
		return
	}

	return
}

func editEvent(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/eventseries.html",
		"./public/html/editEvent.html",
	}

	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files : ", err.Error())
		return
	}

	evnt, alles, err := getEditEventData(r)
	if err != nil {
		fmt.Println("Error parsing getting data from db :", err.Error())
		return
	}

	data := struct {
		Event          Event
		AllEventseries []EventSeries
	}{evnt, alles}

	t.ExecuteTemplate(w, "layout", data)
	return
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

	event, err := extractEventData(r)
	if err != nil {
		fmt.Println("Error extracting event data : ", err.Error())
		return
	}

	event.Time, err = extractEventTime(r.FormValue("date"), r.FormValue("time"))
	if err != nil {
		fmt.Println("Error extracting event time : ", err.Error())
		return
	}

	err = event.Update()
	if err != nil {
		fmt.Println("Error Updating event : ", err.Error())
		return
	}

	allEvents(w, r)
}

func allEventseries(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/eventseries.html",
		"./public/html/allEventSeries.html",
	}

	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}

	all, err := DefaultEventSeries.All()
	if err != nil {
		err = fmt.Errorf("Error reading all event series from db : %s", err.Error())
		return
	}

	t.ExecuteTemplate(w, "layout", all)
}

func eventSeriesEvents(w http.ResponseWriter, r *http.Request) {
	idstr := path.Base(r.URL.Path)
	_, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error, bad id provided : ", err.Error())
		return
	}

	esrs, err := GetEventSeries(idstr)
	if err != nil {
		fmt.Println("Error, getting event series from db : ", err.Error())
		return
	}

	all, err := eventsBelongingToSeries(idstr)
	if err != nil {
		err = fmt.Errorf("Error reading all events belonging to event series from db : %s", err.Error())
		return
	}

	tmpFiles := []string{
		"./public/html/layout.html",
		"./public/html/eventseries.html",
		"./public/html/showOutline.html",
	}

	t, err := template.ParseFiles(tmpFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err.Error())
		return
	}

	data := struct {
		Eventseries string
		Events      []Event
	}{esrs.Name, all}

	t.ExecuteTemplate(w, "layout", data)
}
