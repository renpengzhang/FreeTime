package main

import (
	"fmt"
	"log"
	"net/http"

	"FreeTime/clients"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/profileimages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/eventimages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
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
