package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "freetimedbserver.mysql.database.azure.com"
	database = "freetimedb"
	user     = "foobaradmin@freetimedbserver"
	password = "FreeTimeAdmin123!"
)

// GetAzureMysqlDB is
func GetAzureMysqlDB() DatabaseInterface {
	return AzureMysqlDB{}
}

// AzureMysqlDB is
type AzureMysqlDB struct {
}

func (azureMysqlDB *AzureMysqlDB) getDBConnection() (*sql.DB, error) {
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database.")

	return db, nil
}

func (azureMysqlDB *AzureMysqlDB) execQuery(queryString string) (*sql.Rows, error) {
	db, err := azureMysqlDB.getDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db.Query(queryString)
}
