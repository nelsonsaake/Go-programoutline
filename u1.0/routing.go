package main

import (
	"net/http"

	"github.com/rs/cors"
)

func servePublicFiles() {
	bsjq := http.FileServer(http.Dir("./public/bsjq/"))
	http.Handle("/bsjq/", http.StripPrefix("/bsjq/", bsjq))

	img := http.FileServer(http.Dir("./public/img/"))
	http.Handle("/img/", http.StripPrefix("/img/", img))

	js := http.FileServer(http.Dir("./public/js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	css := http.FileServer(http.Dir("./public/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
}

func routing() {
	// serve public files and resources like css, js...
	servePublicFiles()

	// handle cors
	c := cors.AllowAll()

	// mostly templating
	http.Handle("/", c.Handler(http.HandlerFunc(index)))
	http.Handle("/outlineReq", c.Handler(http.HandlerFunc(outlineReq)))
	http.Handle("/outline", c.Handler(http.HandlerFunc(outline)))
	http.Handle("/defaults", c.Handler(http.HandlerFunc(defaults)))
	http.Handle("/eventseries", c.Handler(http.HandlerFunc(eventseries)))
	http.Handle("/events/create", c.Handler(http.HandlerFunc(createOutline)))
	http.Handle("/events/all", c.Handler(http.HandlerFunc(allEvents)))
	http.Handle("/events/edit/", c.Handler(http.HandlerFunc(editEvent)))
	http.Handle("/events/update/", c.Handler(http.HandlerFunc(updateEvent)))
	http.Handle("/eventseries/all/", c.Handler(http.HandlerFunc(allEventseries)))
	http.Handle("/eventseries/events/", c.Handler(http.HandlerFunc(eventSeriesEvents)))
}
