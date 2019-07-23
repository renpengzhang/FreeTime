package clients

import (
	"FreeTime/operations"
	"fmt"
	"net/http"
	"encoding/json"
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
	interests := parameters.Get("interests")
	// description

	operations.CreateEvent(userName, eventName, startTime, location, interests)
	fmt.Fprintf(w, "%s created event %s with interest tag %s at %s in %s", userName, eventName, interests, startTime, location)
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
	fmt.Fprintf(w, "Got events for %s:\n", userName)

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(eventsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")

	interests, eventsList := operations.GetUserProfile(userName)
	fmt.Fprintf(w, "Got profile for %s:\n", userName)

	w.Header().Set("Content-Type", "application/json")

	js, err := operations.WrapProfileJson(interests, eventsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
