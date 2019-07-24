package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"FreeTime/clients"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//username := r.URL.Query().Get("username")
		//fmt.Fprintf(w, "Hello, %s", username)
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

		body := string(bodyBytes)
		fmt.Fprintf(w, "The body is: %s", body)
	})

	http.HandleFunc("/signup", clients.SignUp)

	http.HandleFunc("/signin", clients.SignIn)

	http.HandleFunc("/createevent", clients.CreateEvent)

	http.HandleFunc("/joinevent", clients.JoinEvent)

	http.HandleFunc("/getjoinedevents", clients.GetJoinedEvents)

	http.HandleFunc("/getallevents", clients.GetAllEvents)

	http.HandleFunc("/getuserprofile", clients.GetUserProfile)

	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
