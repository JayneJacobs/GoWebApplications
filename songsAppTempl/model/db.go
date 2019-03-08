package model

import "database/sql"

var db *sql.DB

// SetDatabase points to the connection in Main from the Model Layer
func SetDatabase(database *sql.DB) {
	db = database
}
