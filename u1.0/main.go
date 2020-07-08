/*
	this app is supposed to be a fun project
	that I use to build my program outline for the semeter for the CCR shepherding ministry
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func server() {

	//
	server := http.Server{Addr: ":8080"}

	// add routes
	routing()

	// start server
	server.ListenAndServe()
}

func stdio() {

	evnt := Event{
		Subject:     "First general meeting",
		Time:        time.Now(),
		EventSeries: EventSeries{Id: 1},
	}

	err := evnt.Create()
	if err != nil {
		fmt.Println("Error creating new event :", err.Error())
		return
	}

	fmt.Println(evnt.All())
}

func main() {
	server()
}

/*
	one page to produce the general outline
		there would be a place to put start date
		there would be place for time interval
		when the above is provided, a form will be generated
		with inputs for the Subjects of the events, and a date on the side

	one page like the first but will allow changes ?not exactly clear
	one page to add details
	one page to see every thing
*/
