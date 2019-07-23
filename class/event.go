package class

import (
	"FreeTime/database"
	"time"
)

// Event is
type Event struct {
	EventID          string
	Name             string
	OwnerID          string
	StartTime        time.Time
	Location         string
	ParticipantCount int
	Description      string
}

func GetAllEvents() ([]*Event, error) {
	db := database.GetAzureMysqlDB()
	dbeventlist, err := db.GetAllEvent()
	if err != nil {
		return nil, err
	}

	var eventlist []*Event
	for _, dbevent := range dbeventlist {
		event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount, dbevent.Description}
		eventlist = append(eventlist, &event)
	}

	return eventlist, nil
}

// GetEventByID is
func GetEventByID(eventID string) (*Event, error) {
	db := database.GetAzureMysqlDB()
	dbevent, err := db.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}
	event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount, dbevent.Description}
	return &event, nil
}

// GetEventByName is
func GetEventByName(eventName string) (*Event, error) {
	db := database.GetAzureMysqlDB()
	dbevent, err := db.GetEventByName(eventName)
	if err != nil {
		return nil, err
	}
	event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount, dbevent.Description}
	return &event, nil
}

// SetEvent is
func SetEvent(event Event) error {
	db := database.GetAzureMysqlDB()
	dbevent := database.DBEvent{event.EventID, event.Name, event.OwnerID, event.StartTime, event.Location, event.ParticipantCount, event.Description}
	return db.SetEvent(dbevent)
}
