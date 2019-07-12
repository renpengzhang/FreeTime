package operations

import (
	"FreeTime/class"
	// "fmt"
	"github.com/google/uuid"
	"time"
)

// SignUp is
func SignUp(userName string, interests string) {
	userID := uuid.New()

	class.SetUser(userName, userID.String())

	// Need to iterate throught interests and set UserInterest by each
	class.AddUserInterest(class.UserInterest{userName, interests})
}

// SignIn is
func SignIn(userName string) {
	class.GetUserByName(userName)
}

// CreateEvent is
func CreateEvent(userName string, eventName string, startTime string, location string) {
	eventID := uuid.New()
	owner, error := class.GetUserByName(userName)

	if error != nil {
		ownerID := owner.ID
		class.SetEvent(class.Event{eventID.String(), eventName, ownerID, time.Now(), location, 0})
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
