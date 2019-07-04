package main

import (
	"fmt"
	"log"
	"net/http"

	"FreeTime/operations"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		fmt.Fprintf(w, "Hello, %s", username)
	})

	http.HandleFunc("/signup", operations.SignUp)

	http.HandleFunc("/signin", operations.SignIn)

	http.HandleFunc("/createevent", operations.CreateEvent)

	http.HandleFunc("/joinevent", operations.JoinEvent)

	http.HandleFunc("/getevents", operations.GetEvents)

	http.HandleFunc("/getuserprofile", operations.GetUserProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
