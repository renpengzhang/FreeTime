package clients

import (
	"FreeTime/operations"
	"FreeTime/commons"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var signupbody commons.SignupBody
	json.Unmarshal(bodyBytes, &signupbody)

	userName := signupbody.Username
	interests := signupbody.Interests

	if err := operations.SignUp(userName, interests); err != nil {
		errMsg := fmt.Sprintf("%s failed to sign up", userName)
		http.Error(w, errMsg, http.StatusInternalServerError)
	} else {
		w.Write([]byte("Succeed to Sign Up"))
		w.WriteHeader(http.StatusOK)
	}
}

// SignIn is
func SignIn(w http.ResponseWriter, r *http.Request) {
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
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var creatEventBody commons.CreatEventBody

	json.Unmarshal(bodyBytes, &creatEventBody)

	userName := creatEventBody.Username
	eventName := creatEventBody.Name
	startTime := creatEventBody.StartTime
	location := creatEventBody.Location
	interests := creatEventBody.Interests
	description := creatEventBody.Description

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
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var joinEventBody commons.JoinEventBody

	json.Unmarshal(bodyBytes, &joinEventBody)

	userName := joinEventBody.Username
	eventID := joinEventBody.Eventid

	if err := operations.JoinEvent(userName, eventID); err != nil {
		errMsg := fmt.Sprintf("%s failed to join event", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	} else {
		w.Write([]byte("Succeed to Join Event"))
		w.WriteHeader(http.StatusOK)
	}
}

// GetEvents is
func GetEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	userName := r.URL.Query().Get("username")

	eventsList, err := operations.GetEvents(userName)
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

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
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
