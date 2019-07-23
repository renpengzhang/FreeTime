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
}

// GetEventByID is
func GetEventByID(eventID string) (*Event, error) {
	db := database.GetAzureMysqlDB()
	dbevent, err := db.GetEventByID(userName)
	if err != nil {
		return nil, err
	}
	event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount}
	return &event, nil
}

// GetEventByName is
func GetEventByName(eventName string) (*Event, error) {
	db := database.GetAzureMysqlDB()
	dbevent, err := db.GetEventByName(userName)
	if err != nil {
		return nil, err
	}
	event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount}
	return &event, nil
}

// SetEvent is
func SetEvent(event Event) error {
	db := database.GetAzureMysqlDB()
	dbevent := database.DBEvent{event.EventID, event.Name, event.OwnerID, event.StartTime, event.Location, event.ParticipantCount}
	return db.SetEvent(dbevent)
}
