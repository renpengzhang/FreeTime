package clients

import (
	"FreeTime/operations"
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

	var signupbody SignupBody
	json.Unmarshal(bodyBytes, &signupbody)

	userName := signupbody.username
	interests := signupbody.interests

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
		//fmt.Fprintf(w, "%s failed to sign in", userName)
		errMsg := fmt.Sprintf("%s failed to sign in, please sign up first", userName)
		http.Error(w, errMsg, http.StatusBadRequest)

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

	var creatEventBody CreatEventBody

	json.Unmarshal(bodyBytes, &creatEventBody)

	userName := creatEventBody.username
	eventName := creatEventBody.name
	startTime := creatEventBody.startTime
	location := creatEventBody.location
	interests := creatEventBody.interests
	// description

	operations.CreateEvent(userName, eventName, startTime, location, interests)
	//fmt.Fprintf(w, "%s created event %s with interest tag %s at %s in %s", userName, eventName, interests, startTime, location)

	w.Write([]byte("Succeed to Creat Event"))
	w.WriteHeader(http.StatusOK)
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

	var joinEventBody JoinEventBody

	json.Unmarshal(bodyBytes, &joinEventBody)

	userName := joinEventBody.username
	eventID := joinEventBody.eventid

	operations.JoinEvent(userName, eventID)
	//fmt.Fprintf(w, "%s joined event %s", userName, eventID)

	w.Write([]byte("Succeed to Join Event"))
	w.WriteHeader(http.StatusOK)
}

// GetEvents is
func GetEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	userName := r.URL.Query().Get("username")

	eventsList := operations.GetEvents(userName)
	fmt.Fprintf(w, "Got events for %s:\n", userName)

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(eventsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetUserProfile is
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	userName := r.URL.Query().Get("username")

	interests, eventsList := operations.GetUserProfile(userName)
	fmt.Fprintf(w, "Got profile for %s:\n", userName)

	w.Header().Set("Content-Type", "application/json")

	js, err := operations.WrapProfileJson(interests, eventsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
