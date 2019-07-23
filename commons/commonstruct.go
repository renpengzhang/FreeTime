package commons

type SignupBody struct {
	Username  string
	Interests string
}

type CreatEventBody struct {
	Username  	string
	Name      	string
	StartTime 	string
	Location  	string
	Interests 	string
	Description string
}

type JoinEventBody struct {
	Username string
	Eventid  string
}