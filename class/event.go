package class

import "time"

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
	return nil, nil
}

// GetEventByName is
func GetEventByName(eventName string) (*Event, error) {
	return nil, nil
}

// SetEvent is
func SetEvent(event Event) error {
	return nil
}
