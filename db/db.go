package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.sql")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("could not create users table." + err.Error())
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id 				INTEGER PRIMARY KEY AUTOINCREMENT,
			name 			TEXT NOT NULL,
			description 	TEXT NOT NULL,
			location 		TEXT NOT NULL,
			dateTime 		DATETIME NOT NULL,
			userID 			INTEGER,
			FOREIGN KEY (userID) REFERENCES Users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table." + err.Error())
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			eventId INTEGER,
			userId INTEGER,
			FOREIGN KEY (eventId) REFERENCES events (Id),
			FOREIGN KEY (userId) REFERENCES users (Id)
		)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("could not create registrations table." + err.Error())
	}
}
