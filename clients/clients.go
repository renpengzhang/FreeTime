package clients

import (
	"FreeTime/commons"
	"FreeTime/operations"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// SignUp is
func SignUp(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println(err.Error())
	}

	userName := r.FormValue("username")
	interests := r.FormValue("interests")

	userID := uuid.New()
	userIDString := userID.String()

	if err := operations.SignUp(userName, interests, userIDString); err != nil {
		errMsg := fmt.Sprintf("%s failed to sign up", userName)
		if err.Error() == commons.DuplicatedUser {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
		return
	}

	file, _, err := r.FormFile("profileimage")
	if err != nil {
		fmt.Printf("No profile image detected: %v\n", err)
		file, err = os.Open("./profileimages/default.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer file.Close()

	f, err := os.Create("./profileimages/" + userIDString + ".jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	w.Write([]byte("Succeed to Sign Up"))
	w.WriteHeader(http.StatusOK)
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
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println(err.Error())
	}

	userName := r.FormValue("username")
	eventName := r.FormValue("name")
	startTime := r.FormValue("starttime")
	location := r.FormValue("location")
	interests := r.FormValue("interests")
	description := r.FormValue("description")

	eventID := uuid.New()
	eventIDString := eventID.String()

	if err := operations.CreateEvent(userName, eventName, startTime, location, interests, description, eventIDString); err != nil {
		errMsg := fmt.Sprintf("%s failed to create event", userName)
		if err.Error() == commons.UserNotExist {
			http.Error(w, errMsg, http.StatusBadRequest)
		} else {
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
		return
	}

	file, _, err := r.FormFile("eventimage")
	if err != nil {
		fmt.Printf("No event image detected: %v\n", err)
		file, err = os.Open("./eventimages/default.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer file.Close()

	f, err := os.Create("./eventimages/" + eventIDString + ".jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	w.Write([]byte("Succeed to Create Event"))
	w.WriteHeader(http.StatusOK)
}

// JoinEvent is
func JoinEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" {
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
