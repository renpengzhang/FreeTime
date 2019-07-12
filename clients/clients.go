package clients

import (
	"FreeTime/operations"
	"fmt"
	"net/http"
)

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving SignUp request")
	parameters := r.URL.Query()
	userName := parameters.Get("username")
	interests := parameters.Get("interests")

	operations.SignUp(userName, interests)
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving SignIn request")
	userName := r.URL.Query().Get("username")

	operations.SignIn(userName)
}

// CreateEvent is
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving CreateEvent request")
	// parameters := r.URL.Query()
	// userName := parameters.Get("username")
	
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
