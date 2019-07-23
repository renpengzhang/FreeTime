package operations

import (
	"FreeTime/class"
	"fmt"
	"time"

	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

// SignUp is
func SignUp(userName string, interests string) error {
	userID := uuid.New()
	userIDString := userID.String()
	if err := class.SetUser(userName, userIDString); err != nil {
		fmt.Printf("Set User failed\n")
		return err
	}

	// Print successful msg to console
	fmt.Printf("Set UserID: %s to %s successfully\n", userIDString, userName)

	// Need to iterate throught interests and set UserInterest by each
	if err := class.AddUserInterest(class.UserInterest{userName, interests}); err != nil {
		fmt.Printf("Add Interests failed\n")
		return err
	}

	// Print successful msg to console
	fmt.Printf("Add %s - %s successfully\n", userName, interests)
	return nil
}

// SignIn is
func SignIn(userName string) error {
	user, error := class.GetUserByName(userName)

	if error == nil {
		// Print successful msg to console
		fmt.Printf("%s userID is %s\n", user.Username, user.ID)
		return nil
	} else {
		fmt.Printf("%s sign in failed\n", userName)
		return error
	}
}

// CreateEvent is
func CreateEvent(userName string, eventName string, startTime string, location string, interests string) {
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

		for _, interest := range strings.Split(interests, ",") {
			eventInterest := class.EventInterest{eventIDString, interest}
			if class.AddEventInterest(eventInterest) == nil {
				fmt.Printf("Add with Interest: %s\n", interest)
			} else {
				fmt.Printf("Add event with interest failed\n")
			}
		}

	} else {
		fmt.Printf("There is no %s in db\n", userName)
	}
}

// JoinEvent is
func JoinEvent(userName string, eventID string) {
	user, userError := class.GetUserByName(userName)
	if userError == nil {
		joinError := class.AddUserJoinedEvent(class.UserJoinedEvent{user.ID, eventID})
		if joinError == nil {
			fmt.Printf("%s - %s join %s\n", userName, user.ID, eventID)
		} else {
			fmt.Printf("Join event failed\n")
		}
	} else {
		fmt.Printf("There is no %s in db\n", userName)
	}
}

// GetEvents is
func GetEvents(userName string) []*class.Event {
	user, userError := class.GetUserByName(userName)
	var eventsList []*class.Event
	var eventErr error

	if userError == nil {
		userIDString := user.ID
		eventsList, eventErr = class.GetEventsByUserID(userIDString)
		if eventsList != nil && eventErr == nil {
			// Print successful msg to console
			fmt.Printf("Get events for %s - %s: %v\n", user.Username, userIDString, eventsList)
		}
	} else {
		fmt.Printf("There is no %s in db\n", userName)
	}

	return eventsList
}

// GetUserProfile is
func GetUserProfile(userName string) ([]string, []*class.Event) {
	user, userError := class.GetUserByName(userName)
	var interests []string
	var eventsList []*class.Event
	var eventErr error

	if userError == nil {
		userIDString := user.ID
		userInterests := class.GetInterestsByUserID(userIDString)

		for _, userInterest := range userInterests {
			interests = append(interests, userInterest.Interest)
		}

		if interests != nil {
			// Print successful msg to console
			fmt.Printf("Get interests for %s - %s: %v\n", user.Username, userIDString, interests)
		}

		eventsList, eventErr = class.GetEventsByUserID(userIDString)
		if eventsList != nil && eventErr == nil {
			// Print successful msg to console
			fmt.Printf("Get events for %s - %s: %v\n", user.Username, userIDString, eventsList)
		}
	} else {
		fmt.Printf("There is no %s in db\n", userName)
	}

	return interests, eventsList
}

func WrapProfileJson(interests []string, events []*class.Event) ([]byte, error) {
	type Profile struct {
		Interests []string
		Events    []*class.Event
	}

	profile := Profile{interests, events}
	js, err := json.Marshal(profile)

	return js, err
}
