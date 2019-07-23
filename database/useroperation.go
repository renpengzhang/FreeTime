package database

import (
	"errors"
	"fmt"
)

type DBUser struct {
	ID       string
	Username string
}

// DBUserInterest is
type DBUserInterest struct {
	Interest string
	UserID   string
}

type DBUserJoinedEvent struct {
	UserID  string
	EventID string
}

// GetUserByName is
func (azureMysqlDB AzureMysqlDB) GetUserByName(userName string) (*DBUser, error) {
	queryString := fmt.Sprintf("SELECT * from user where username = '%s';", userName)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := DBUser{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == "" {
		return nil, errors.New("User not exist")
	}

	return &user, nil
}

// GetUserByID is
func (azureMysqlDB AzureMysqlDB) GetUserByID(userID string) (*DBUser, error) {
	queryString := fmt.Sprintf("SELECT * from user where id = '%s';", userID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := DBUser{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == "" {
		return nil, errors.New("User not exist")
	}

	return &user, nil
}

// SetUser is
func (azureMysqlDB AzureMysqlDB) SetUser(userName, userID string) error {
	queryString := fmt.Sprintf("INSERT into user (id, username) values ('%s', '%s');", userID, userName)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

// JoinEvent is
func (azureMysqlDB AzureMysqlDB) JoinEvent(userID string, eventID string) error {
	queryString := fmt.Sprintf("INSERT into userjoinedevent (userid, eventid) values ('%s', '%s');", userID, eventID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (azureMysqlDB AzureMysqlDB) GetEventsByUserID(userID string) ([]*DBEvent, error) {
	var eventList []*DBEvent
	queryString := fmt.Sprintf("SELECT eventid from event where ownerid = '%s';", userID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		eventID := ""
		err := rows.Scan(&eventID)
		if err != nil {
			return nil, err
		}

		event, err := azureMysqlDB.GetEventByID(eventID)
		if err != nil {
			return nil, err
		}

		eventList = append(eventList, event)
	}

	return eventList, nil
}

// GetInterestsByUserID is
func (azureMysqlDB AzureMysqlDB) GetInterestsByUserID(userID string) ([]DBUserInterest, error) {
	var interestList []DBUserInterest
	queryString := fmt.Sprintf("SELECT interest from userinterest where userid = '%s';", userID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		interest := ""
		err := rows.Scan(&interest)
		if err != nil {
			return nil, err
		}

		interestList = append(interestList, DBUserInterest{userID, interest})
	}

	return interestList, nil
}

// AddUserInterest is
func (azureMysqlDB AzureMysqlDB) AddUserInterest(userInterest DBUserInterest) error {
	queryString := fmt.Sprintf("INSERT into userinterest (userid, interest) values ('%s', '%s');", userInterest.UserID, userInterest.Interest)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

// GetUserJoinedEvent is
func (azureMysqlDB AzureMysqlDB) GetUserJoinedEvents(userID string) ([]DBUserJoinedEvent, error) {
	var eventIDList []DBUserJoinedEvent
	queryString := fmt.Sprintf("SELECT eventid from userjoinedevent where userid = '%s';", userID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		eventid := ""
		err := rows.Scan(&eventid)
		if err != nil {
			return nil, err
		}

		eventIDList = append(eventIDList, DBUserJoinedEvent{userID, eventid})
	}

	return eventIDList, nil
}

// AddUserJoinedEvent is
func (azureMysqlDB AzureMysqlDB) AddUserJoinedEvent(userJoinedEvent DBUserJoinedEvent) error {
	queryString := fmt.Sprintf("INSERT into userjoinedevent (userid, eventid) values ('%s', '%s');", userJoinedEvent.UserID, userJoinedEvent.EventID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}
