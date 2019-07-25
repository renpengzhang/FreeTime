package operations

import (
	"FreeTime/class"
	"fmt"
	"time"

	"encoding/json"
	"strings"
)

// Used for wrap json
type Profile struct {
	Interests []string
	Events    []*class.Event
	UserID    string
}

type EventInterests struct {
	Event     *class.Event
	Interests []string
}

// SignUp is
func SignUp(userName string, interests string, userIDString string) error {
	if err := class.SetUser(userName, userIDString); err != nil {
		fmt.Printf("Set User failed\n")
		fmt.Println(err)
		return err
	}

	// Print successful msg to console
	fmt.Printf("Set UserID: %s to %s successfully\n", userIDString, userName)

	// Need to iterate throught interests and set UserInterest by each
	for _, interest := range strings.Split(interests, ",") {
		if err := class.AddUserInterest(class.UserInterest{userIDString, interest}); err != nil {
			fmt.Printf("Add Interests failed\n")
			fmt.Println(err)
			return err
		}
	}

	// Print successful msg to console
	fmt.Printf("Add %s - %s successfully\n", userName, interests)
	return nil
}

// SignIn is
func SignIn(userName string) error {
	user, userError := class.GetUserByName(userName)

	if userError == nil {
		// Print successful msg to console
		fmt.Printf("%s userID is %s\n", user.Username, user.ID)
		return nil
	} else {
		fmt.Printf("%s sign in failed\n", userName)
		fmt.Println(userError)
		return userError
	}
}

// CreateEvent is
func CreateEvent(userName string, eventName string, startTime string, location string, interests string, description string, eventIDString string) error {
	owner, userError := class.GetUserByName(userName)

	if userError == nil {
		ownerID := owner.ID
		eventTime, timeError := time.Parse("2006-1-2 15:04:05", startTime)
		eventErr := class.SetEvent(class.Event{eventIDString, eventName, ownerID, eventTime, location, 1, description})
		if timeError == nil && eventErr == nil {
			// Print successful msg to console
			fmt.Printf("Event: %s - %s, Time: %s, \n", eventName, eventIDString, eventTime)
			fmt.Printf("UserID: %s, Location: %s, PeopleCount: 1\n", ownerID, location)
			joinedErr := class.AddUserJoinedEvent(class.UserJoinedEvent{ownerID, eventIDString})
			if joinedErr != nil {
				fmt.Printf("Join event failed\n")
				fmt.Println(joinedErr)
				return joinedErr
			}
		} else if timeError == nil {
			fmt.Printf("Time parse error\n")
			fmt.Println(timeError)
			return timeError
		} else {
			fmt.Printf("Set event failed\n")
			fmt.Println(eventErr)
			return eventErr
		}

		for _, interest := range strings.Split(interests, ",") {
			eventInterest := class.EventInterest{eventIDString, interest}
			eventErr = class.AddEventInterest(eventInterest)
			if eventErr == nil {
				fmt.Printf("Add with Interest: %s\n", interest)
			} else {
				fmt.Printf("Add event with interest failed\n")
				fmt.Println(eventErr)
				return eventErr
			}
		}

	} else {
		fmt.Printf("Get user %s failed\n", userName)
		fmt.Println(userError)
		return userError
	}

	return nil
}

// JoinEvent is
func JoinEvent(userName string, eventID string) error {
	user, userError := class.GetUserByName(userName)
	if userError == nil {
		joinError := class.AddUserJoinedEvent(class.UserJoinedEvent{user.ID, eventID})
		if joinError == nil {
			fmt.Printf("%s - %s join %s\n", userName, user.ID, eventID)
			event, eventErr := class.GetEventByID(eventID)
			if eventErr == nil {
				event.ParticipantCount = event.ParticipantCount + 1
				class.SetEvent(*event)
			} else {
				fmt.Printf("Get event failed\n")
				fmt.Println(eventErr)
				return eventErr
			}
		} else {
			fmt.Printf("Join event failed\n")
			fmt.Println(joinError)
			return joinError
		}
	} else {
		fmt.Printf("Get user %s failed\n", userName)
		fmt.Println(userError)
		return userError
	}

	return nil
}

// GetJoinedEvents is
func GetJoinedEvents(userName string) ([]*class.Event, error) {
	user, userError := class.GetUserByName(userName)
	var eventsList []*class.Event

	if userError == nil {
		userIDString := user.ID
		joinedEvents, err := class.GetUserJoinedEvents(userIDString)
		if err != nil {
			return nil, err
		}

		if joinedEvents != nil {
			// Print successful msg to console
			fmt.Printf("Get events for %s - %s: %v\n", user.Username, userIDString, joinedEvents)
			for _, joinedEvent := range joinedEvents {
				event, eventErr := class.GetEventByID(joinedEvent.EventID)
				if eventErr == nil {
					eventsList = append(eventsList, event)
				} else {
					fmt.Printf("Get event by id failed\n")
					fmt.Println(eventErr)
					return nil, eventErr
				}
			}
		} else {
			fmt.Printf("There is no events for user %s\n", userName)
		}
	} else {
		fmt.Printf("Get user %s failed\n", userName)
		fmt.Println(userError)
		return nil, userError
	}

	return eventsList, nil
}

func GetAllEvents() ([]*EventInterests, error) {
	events, eventsErr := class.GetAllEvents()
	if eventsErr != nil {
		fmt.Println("Get all events failed")
		fmt.Println(eventsErr)
		return nil, eventsErr
	}

	var eventsInterests []*EventInterests
	for _, event := range events {
		interestsEvent, interestsErr := class.GetInterestsByEventID(event.EventID)
		if interestsErr != nil {
			return nil, interestsErr
		}
		var interests []string
		for _, interest := range interestsEvent {
			interests = append(interests, interest.Interest)
		}
		eventInterest := EventInterests{event, interests}
		eventsInterests = append(eventsInterests, &eventInterest)
	}
	return eventsInterests, nil
}

// GetCreatedEvents is
func GetCreatedEvents(userName string) ([]*class.Event, error) {
	user, userError := class.GetUserByName(userName)
	var eventsList []*class.Event
	var eventErr error

	if userError == nil {
		userIDString := user.ID
		eventsList, eventErr = class.GetCreatedEventsByUserID(userIDString)
		if eventErr == nil {
			// Print successful msg to console
			if eventsList == nil {
				fmt.Printf("There is no events for %s\n", user.Username)
			} else {
				fmt.Printf("Get events for %s - %s: %v\n", user.Username, userIDString, eventsList)
			}
		} else {
			fmt.Printf("Get events for %s failed\n", user.Username)
			fmt.Println(eventErr)
			return nil, eventErr
		}
	} else {
		fmt.Printf("Get user %s failed\n", userName)
		fmt.Println(userError)
		return nil, userError
	}

	return eventsList, nil
}

// GetUserProfile is
func GetUserProfile(userName string) ([]string, []*class.Event, string, error) {
	user, userError := class.GetUserByName(userName)
	var interests []string
	var createdEvents []*class.Event

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

		var err error
		createdEvents, err = class.GetCreatedEventsByUserID(userIDString)
		if err != nil {
			return nil, nil, "", err
		}

		if createdEvents != nil {
			// Print successful msg to console
			fmt.Printf("Get created events for %s - %s: %v\n", user.Username, userIDString, createdEvents)
		} else {
			fmt.Printf("There is no events for user %s\n", userName)
		}
	} else {
		fmt.Printf("Get user %s failed\n", userName)
		fmt.Println(userError)
		return nil, nil, "", userError
	}

	return interests, createdEvents, user.ID, nil
}

func WrapProfileJson(interests []string, events []*class.Event, userID string) ([]byte, error) {
	profile := Profile{interests, events, userID}
	js, err := json.Marshal(profile)

	return js, err
}
