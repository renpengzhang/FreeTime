package operations

import (
	"FreeTime/class"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// SignUp is
func SignUp(userName string, interests string) {
	userID := uuid.New()
	userIDString := userID.String()
	if class.SetUser(userName, userIDString) == nil {
		// Print successful msg to console
		fmt.Printf("Set UserID: %s to %s successfully\n", userIDString, userName)
	} else {
		fmt.Printf("Set User failed\n")
	}

	// Need to iterate throught interests and set UserInterest by each
	if class.AddUserInterest(class.UserInterest{userName, interests}) == nil {
		// Print successful msg to console
		fmt.Printf("Add %s - %s successfully\n", userName, interests)
	} else {
		fmt.Printf("Add Interests failed\n")
	}
}

// SignIn is
func SignIn(userName string) {
	user, error := class.GetUserByName(userName)

	if error == nil {
		// Print successful msg to console
		fmt.Printf("%s userID is %s\n", user.Username, user.ID)
	} else {
		fmt.Printf("%s sign in failed\n", userName)
	}
}

// CreateEvent is
func CreateEvent(userName string, eventName string, startTime string, location string) {
	eventID := uuid.New()
	eventIDString := eventID.String()
	owner, userError := class.GetUserByName(userName)

	if userError == nil {
		ownerID := owner.ID
		eventTime, timeError := time.Parse("2006-01-02 15:04:05", startTime)

		if timeError == nil && class.SetEvent(class.Event{eventIDString, eventName, ownerID, eventTime, location, 1}) == nil {
			// Print successful msg to console
			fmt.Printf("Event: %s - %s, Time: %s, \n", eventName, eventIDString, eventTime)
			fmt.Printf("UserID: %s, Location: %s, PeopleCount: 1\n", ownerID, location)
		} else if timeError == nil {
			fmt.Printf("Time parse error\n")
		} else {
			fmt.Printf("Set event failed\n")
		}
	} else {
		fmt.Printf("There is no %s in db\n", userName)
	}
}

// JoinEvent is
func JoinEvent() {

}

// GetEvents is
func GetEvents() {

}

// GetUserProfile is
func GetUserProfile() {

}
