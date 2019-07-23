package main

import (
	"fmt"
	"log"
	"net/http"

	"FreeTime/clients"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		fmt.Fprintf(w, "Hello, %s", username)
	})

	http.HandleFunc("/signup", clients.SignUp)

	http.HandleFunc("/signin", clients.SignIn)

	http.HandleFunc("/createevent", clients.CreateEvent)

	http.HandleFunc("/joinevent", clients.JoinEvent)

	http.HandleFunc("/getevents", clients.GetEvents)

	http.HandleFunc("/getuserprofile", clients.GetUserProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
