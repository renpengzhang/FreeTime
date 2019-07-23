package class

import "FreeTime/database"

// EventInterest is
type EventInterest struct {
	EventID  string
	Interest string
}

// GetInterestsByEventID is
func GetInterestsByEventID(eventID string) []EventInterest {
	db := database.GetAzureMysqlDB()

	var eventInterestList []EventInterest

	dbEventInterestList, _ := db.GetInterestsByEventID(eventID)
	for _, dbEventInterest := range dbEventInterestList {
		eventInterestList = append(eventInterestList, EventInterest{dbEventInterest.EventID, dbEventInterest.Interest})
	}
	return eventInterestList
}

// AddEventInterest is
func AddEventInterest(eventInterest EventInterest) error {
	db := database.GetAzureMysqlDB()

	dbEventInterest := dabase.DBEventInterest{eventInterest.EventID, eventInterest.Interest}

	return db.AddEventInterest(dbEventInterest)
}
