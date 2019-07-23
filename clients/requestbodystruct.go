package clients

type SignupBody struct {
	username  string
	interests string
}

type CreatEventBody struct {
	username  	string
	name      	string
	startTime 	string
	location  	string
	interests 	string
	description string
}

type JoinEventBody struct {
	username string
	eventid  string
}
