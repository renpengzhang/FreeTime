package class

// UserInterest is
type UserInterest struct {
	Interest string
	UserID   string
}

// GetInterestsByUserID is
func GetInterestsByUserID(userID string) []UserInterest {
	return []UserInterest{UserInterest{"TestInterest", userID}}
}

// AddUserInterest is
func AddUserInterest(userInterest UserInterest) error {
	return nil
}
