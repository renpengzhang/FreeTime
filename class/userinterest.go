package class

import "FreeTime/database"

// UserInterest is
type UserInterest struct {
	Interest string
	UserID   string
}

// GetInterestsByUserID is
func GetInterestsByUserID(userID string) []UserInterest {
	db := database.GetAzureMysqlDB()

	var userInterestList []UserInterest

	dbUserInterestList, _ := db.GetInterestsByUserID(userID)
	for _, dbUserInterest := range dbUserInterestList {
		userInterestList = append(userInterestList, UserInterest{dbUserInterestList.UserID, dbUserInterestList.Interest})
	}
	return userInterestList
}

// AddUserInterest is
func AddUserInterest(userInterest UserInterest) error {
	db := database.GetAzureMysqlDB()

	dbUserInterest := dabase.DBUserInterest{userInterest.UserID, userInterest.Interest}

	return db.AddUserInterest(dbUserInterest)
}
