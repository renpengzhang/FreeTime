package operations

import (
	"FreeTime/class"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving SignUp request")
	parameters := r.URL.Query()
	userName := parameters.Get("username")
	interests := parameters.Get("interests")
	userID := uuid.New()

	class.SetUser(userName, userID.String())
	class.AddUserInterest(class.UserInterest{userName, interests})
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving SignIn request")
}

// CreateEvent is
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Receiving CreateEvent request")
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
