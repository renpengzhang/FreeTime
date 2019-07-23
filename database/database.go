package database

// DatabaseInterface is
type DatabaseInterface interface {
	GetUserByName(userName string) (*DBUser, error)

	GetUserByID(userID string) (*DBUser, error)

	SetUser(userName, userID string) error

	JoinEvent(userID string, eventID string) error

	GetEventsByUserID(userID string) ([]*DBEvent, error)

	GetEventByID(eventID string) (*DBEvent, error)

	GetEventByName(eventName string) (*DBEvent, error)

	SetEvent(event DBEvent) error

	GetInterestsByEventID(eventID string) ([]DBEventInterest, error)

	AddEventInterest(eventInterest DBEventInterest) error

	GetInterestsByUserID(userID string) ([]DBUserInterest, error)

	AddUserInterest(userInterest DBUserInterest) error

	GetUserJoinedEvents(userID string) ([]DBUserJoinedEvent, error)

	AddUserJoinedEvent(userJoinedEvent DBUserJoinedEvent) error
}
