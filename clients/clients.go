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

	operations.SignUp(userName, interests)
	fmt.Fprintf(w, "%s signed up interests %s", userName, interests)
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")

	operations.SignIn(userName)
	fmt.Fprintf(w, "%s signed in", userName)
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
	fmt.Fprintf(w, "Receiving JoinEvent request")
}

// GetEvents is
func GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving GetEvents request")
}

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving GetUserProfile request")
}
