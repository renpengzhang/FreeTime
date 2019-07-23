package class

import (
	"FreeTime/database"
)

// User is
type User struct {
	ID       string
	Username string
}

// GetUserByName is
func GetUserByName(userName string) (*User, error) {
	db := database.GetAzureMysqlDB()
	dbUser, err := db.GetUserByName(userName)
	if err != nil {
		return nil, err
	}
	user := User{dbUser.ID, dbUser.Username}
	return &user, nil
}

// GetUserByID is
func GetUserByID(userID string) (*User, error) {
	db := database.GetAzureMysqlDB()
	dbUser, err := db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	user := User{dbUser.ID, dbUser.Username}
	return &user, nil
}

// SetUser is
func SetUser(userName, userID string) error {
	db := database.GetAzureMysqlDB()
	return db.SetUser(userName, userID)
}

// JoinEvent is
func JoinEvent(userID string, eventID string) error {
	db := database.GetAzureMysqlDB()
	return db.JoinEvent(userID, eventID)
}

// GetEventsByUserID is
func GetCreatedEventsByUserID(userID string) ([]*Event, error) {
	db := database.GetAzureMysqlDB()
	dbEventList, err := db.GetEventsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var eventList []*Event
	for _, dbevent := range dbEventList {
		event := Event{dbevent.EventID, dbevent.Name, dbevent.OwnerID, dbevent.StartTime, dbevent.Location, dbevent.ParticipantCount, dbevent.Description}
		eventList = append(eventList, &event)
	}
	return eventList, nil
}
