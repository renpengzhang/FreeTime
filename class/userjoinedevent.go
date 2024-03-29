package class

import "FreeTime/database"

// UserJoinedEvent is
type UserJoinedEvent struct {
	UserID  string
	EventID string
}

// GetUserJoinedEvent is
func GetUserJoinedEvents(userID string) ([]UserJoinedEvent, error) {
	db := database.GetAzureMysqlDB()

	var userJoinedEventList []UserJoinedEvent

	dbUserEventsList, err := db.GetUserJoinedEvents(userID)
	if err != nil {
		return nil, err
	}
	for _, dbUserEvent := range dbUserEventsList {
		userJoinedEventList = append(userJoinedEventList, UserJoinedEvent{dbUserEvent.UserID, dbUserEvent.EventID})
	}
	return userJoinedEventList, nil
}

// AddUserJoinedEvent is
func AddUserJoinedEvent(userJoinedEvent UserJoinedEvent) error {
	db := database.GetAzureMysqlDB()

	dbUserEvent := database.DBUserJoinedEvent{userJoinedEvent.UserID, userJoinedEvent.EventID}

	return db.AddUserJoinedEvent(dbUserEvent)
}
