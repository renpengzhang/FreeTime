package clients

import (
	"FreeTime/commons"
	"FreeTime/operations"
	"encoding/json"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" && r.Method != "GET"{
		http.Error(w, "Method not allowed", 405)
		return
	}

	parameters := r.URL.Query()
	userName := parameters.Get("username")
	interests := parameters.Get("interests")

	if err := operations.SignUp(userName, interests); err != nil {
		errMsg := fmt.Sprintf("%s failed to sign up", userName)
		if err.Error() == commons.DuplicatedUser {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Succeed to Sign Up"))
		w.WriteHeader(http.StatusOK)
	}
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userName := r.URL.Query().Get("username")

	if err := operations.SignIn(userName); err != nil {
		errMsg := fmt.Sprintf("%s failed to sign in, please sign up first", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Succeed to Sign In"))
		w.WriteHeader(http.StatusOK)
	}
}

// CreateEvent is
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" && r.Method != "GET"{
		http.Error(w, "Method not allowed", 405)
		return
	}

	parameters := r.URL.Query()
	userName := parameters.Get("username")
	eventName := parameters.Get("name")
	startTime := parameters.Get("startTime")
	location := parameters.Get("location")
	interests := parameters.Get("interests")
	description := parameters.Get("description")

	if err := operations.CreateEvent(userName, eventName, startTime, location, interests, description); err != nil {
		errMsg := fmt.Sprintf("%s failed to create event", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Succeed to Create Event"))
		w.WriteHeader(http.StatusOK)
	}
}

// JoinEvent is
func JoinEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" && r.Method != "GET"{
		http.Error(w, "Method not allowed", 405)
		return
	}

	parameters := r.URL.Query()
	userName := parameters.Get("username")
	eventID := parameters.Get("eventId")

	if err := operations.JoinEvent(userName, eventID); err != nil {
		errMsg := fmt.Sprintf("%s failed to join event", userName)
		if err.Error() == commons.UserNotExist || err.Error() == commons.AlreadyJoinedEvent {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Succeed to Join Event"))
		w.WriteHeader(http.StatusOK)
	}
}

// GetJoinedEvents is
func GetJoinedEvents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	userName := r.URL.Query().Get("username")

	eventsList, err := operations.GetJoinedEvents(userName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get events for %s", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		js, err := json.Marshal(eventsList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

// GetAllEvents is
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	eventsList, err := operations.GetAllEvents()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get all events")
		http.Error(w, errMsg, http.StatusInternalServerError)
	} else {
		js, err := json.Marshal(eventsList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	userName := r.URL.Query().Get("username")

	interests, eventsList, err := operations.GetUserProfile(userName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get profile for %s", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
		return
	} else {
		js, err := operations.WrapProfileJson(interests, eventsList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}
