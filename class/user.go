package class

// User is
type User struct {
	ID       string
	Username string
}

// GetUserByName is
func GetUserByName(userName string) (*User, error) {
	user := User{"TestUserID", "TestUserName"}
	return &user, nil
}

// GetUserByID is
func GetUserByID(userID string) (*User, error) {
	return nil, nil
}

// SetUser is
func SetUser(userName, userID string) error {
	return nil
}

// JoinEvent is
func JoinEvent(userID string, eventID string) error {
	return nil
}

// GetEventsByUserID is
func GetEventsByUserID(userID string) ([]Event, error) {
	return nil, nil
}
