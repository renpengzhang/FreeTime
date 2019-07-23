package clients

import (
	"FreeTime/operations"
	"fmt"
	"net/http"
)

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	userName := parameters.Get("username")
	interests := parameters.Get("interests")

	if err := operations.SignUp(userName, interests); err != nil {
		fmt.Fprintf(w, "%s failed to sign up with interests %s", userName, interests)
	} else {
		fmt.Fprintf(w, "%s succeeded to sign up with interests %s", userName, interests)
	}
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")

	if err := operations.SignIn(userName); err != nil {
		fmt.Fprintf(w, "%s failed to sign in", userName)
	} else {
		fmt.Fprintf(w, "%s succeeded to sign in", userName)
	}
}

// CreateEvent is
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	userName := parameters.Get("username")
	eventName := parameters.Get("name")
	startTime := parameters.Get("startTime")
	location := parameters.Get("location")

	operations.CreateEvent(userName, eventName, startTime, location)
	fmt.Fprintf(w, "%s created event %s at %s in %s", userName, eventName, startTime, location)
}

// JoinEvent is
func JoinEvent(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	userName := parameters.Get("username")
	eventID := parameters.Get("eventId")

	operations.JoinEvent(userName, eventID)
	fmt.Fprintf(w, "%s joined event %s", userName, eventID)
}

// GetEvents is
func GetEvents(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")

	eventsList := operations.GetEvents(userName)
	fmt.Fprintf(w, "Got %v for %s", eventsList, userName)
}

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")

	interests, eventsList := operations.GetUserProfile(userName)
	fmt.Fprintf(w, "Got interests: %v\nand events: %v\nfor %s successfully", interests, eventsList, userName)
}
