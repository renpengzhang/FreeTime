package class

// UserJoinedEvent is
type UserJoinedEvent struct {
	UserID  string
	EventID string
}

// GetUserJoinedEvent is
func GetUserJoinedEvents(userID string) []UserJoinedEvent {
	return []UserJoinedEvent{UserJoinedEvent{userID, "TestEventID"}}
}

// AddUserJoinedEvent is
func AddUserJoinedEvent(userJoinedEvent UserJoinedEvent) error {
	return nil
}