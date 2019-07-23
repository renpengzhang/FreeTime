package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DBEvent struct {
	EventID          string
	Name             string
	OwnerID          string
	StartTime        time.Time
	Location         string
	ParticipantCount int
}

type DBEventInterest struct {
	EventID  string
	Interest string
}

func (azureMysqlDB AzureMysqlDB) GetEventByID(eventID string) (*DBEvent, error) {
	queryString := fmt.Sprintf("SELECT * from event where eventid = '%s';", eventID)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	event := DBEvent{}
	var nt mysql.NullTime

	if rows.Next() {
		err := rows.Scan(&event.EventID, &event.Name, &event.OwnerID, &nt, &event.Location, &event.ParticipantCount)
		if err != nil {
			return nil, err
		}
	}
	if event.EventID == "" {
		return nil, errors.New("Event not exist")
	}

	event.StartTime = nt.Time

	return &event, nil
}

func (azureMysqlDB AzureMysqlDB) GetEventByName(eventName string) (*DBEvent, error) {
	queryString := fmt.Sprintf("SELECT * from event where name = '%s';", eventName)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	event := DBEvent{}
	var nt mysql.NullTime

	if rows.Next() {
		err := rows.Scan(&event.EventID, &event.Name, &event.OwnerID, &nt, &event.Location, &event.ParticipantCount)
		if err != nil {
			return nil, err
		}
	}
	if event.EventID == "" {
		return nil, errors.New("Event not exist")
	}

	event.StartTime = nt.Time

	return &event, nil
}

func (azureMysqlDB AzureMysqlDB) SetEvent(event DBEvent) error {
	// Need to Update
	queryString := fmt.Sprintf("INSERT into event (eventid, name, ownerid, starttime, location, participantCount) values ('%s', '%s', '%s', NOW(), '%s', %d);", event.EventID, event.Name, event.OwnerID, event.Location, event.ParticipantCount)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

// GetInterestsByEventID is
func (azureMysqlDB AzureMysqlDB) GetInterestsByEventID(eventID string) ([]DBEventInterest, error) {
	var interestList []DBEventInterest
	queryString := fmt.Sprintf("SELECT interest from eventinterest where eventid = '%s';", eventID)

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

		interestList = append(interestList, DBEventInterest{eventID, interest})
	}

	return interestList, nil
}

// AddEventInterest is
func (azureMysqlDB AzureMysqlDB) AddEventInterest(eventInterest DBEventInterest) error {
	queryString := fmt.Sprintf("INSERT into eventinterest (eventid, interest) values ('%s', '%s');", eventInterest.EventID, eventInterest.Interest)

	rows, err := azureMysqlDB.execQuery(queryString)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}
