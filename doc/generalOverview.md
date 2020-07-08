ProgramOutline 


The popular norm is that: at the start of a semester, the various different group, organisation and institutions will produce a document detailing the dates and time for all their events and activities. This is usually known as the time-table, program-outline, or calendar. Some of the our at a regular interval. Example: CCR Shepherding Ministry will meet every 2weeks Saturday behind the main auditorium. Others don't follow any particular pattern. 

This is a web application.
It can be used to create a program outline that can be printed out as a pdf document or sent as link. Users of the link can follow it to set a reminder in their Google Calendar. The Calendar notification set will have a link to a web page. Before the Event on the program outline, it will have information on the event like MC, and many more. After the program the link will, hold information on the event plus some or information on what happend durring the events. 

The purpose of this applications is:
+ Provide an easy flow for creating these program outlines. 
+ Generate a pdf version of the program outline that can be easily distributed.
+ Generate a link for event attendee to add programs to their calendar. This means that when the time draws near, their own cell phones and computers will notify them.
+ A web page to store information on the events before and after the event occured.

The system maintains data on two things:
+ Events
+ EventSeries

Events
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

EventSeries
	type EventSeries struct {
		Id   int
		Name string
	}

What are something that can be done on the application.
+ CRUD event
+ CRUD eventseries
+ Adding of eventseries should be flexible
	+ Which means we create the eventSeries shell and then we add the events 
	+ As part of flexibility, events that occur periodically can be easily added. eg. Every 14days, or Every two weeks
	+ Single occuring events can be added
	+ Special periodically occuring events. eg. Every first Monday of the month
	+ An existing event can be added to an event series
+ There should a single pushButton to print the information about an event into a pdf
+ There should be a sharable link generated that can be used to mark a google calendar
	+ There should a single pushbutton that would allow user to share this link to the various different socail media plaforms
+ A page should be automatically synthesized that would contain all the neccessary information on an event. 
+ The synthesized page should be update-able
+ The synthesized page should be reachable from the reminders set on the calendars
+ + after the program, the synthesized page can be update to hold information about the program

Database
	Tables
		Event
		EventSeries
	Relationships
		An EventSeries is made up of events
		An Event can exist outside an event series 

MS
	programOutline
	

