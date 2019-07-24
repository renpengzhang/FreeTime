package class

import "FreeTime/database"

// EventInterest is
type EventInterest struct {
	EventID  string
	Interest string
}

// GetInterestsByEventID is
func GetInterestsByEventID(eventID string) ([]EventInterest, error) {
	db := database.GetAzureMysqlDB()

	var eventInterestList []EventInterest

	dbEventInterestList, err := db.GetInterestsByEventID(eventID)
	if err != nil {
		return nil, err
	}
	for _, dbEventInterest := range dbEventInterestList {
		eventInterestList = append(eventInterestList, EventInterest{dbEventInterest.EventID, dbEventInterest.Interest})
	}
	return eventInterestList, nil
}

// AddEventInterest is
func AddEventInterest(eventInterest EventInterest) error {
	db := database.GetAzureMysqlDB()

	dbEventInterest := database.DBEventInterest{eventInterest.EventID, eventInterest.Interest}

	return db.AddEventInterest(dbEventInterest)
}
