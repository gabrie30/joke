package date

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gabrie30/joke/configs"
)

// Date recoreds data from fetching jokes
type Date struct {
	ID           int
	DateFetched  string
	JokesFetched int
}

// Save saves a date to the db
func (d *Date) Save() error {

	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO dates (date_fetched, jokes_fetched) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = statement.Exec(d.DateFetched, d.JokesFetched)

	if err != nil {
		return err
	}

	return nil
}

// NewDate creates a new joke
func NewDate() *Date {
	return &Date{}
}

// NewFetch creates a new date entry
func NewFetch(c int) error {
	d := NewDate()
	todaysDate := time.Now().Format("01-02-2006")
	d.DateFetched = todaysDate
	d.JokesFetched = c

	err := d.Save()
	if err != nil {
		fmt.Println("Could not save new fetch data, removing jokes")
		return err
	}

	return nil
}

// LastFetchDate returns the last fetch date
func LastFetchDate() string {
	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `SELECT date_fetched FROM dates ORDER BY id DESC LIMIT 1`
	var dateFetched string
	err = db.QueryRow(sqlStmt).Scan(&dateFetched)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatalf("Error: %v, try running 'joke db setup' then try again", err)
		}

		return ""
	}
	return dateFetched
}

// HasEntry determines if there are any dates in the database
func HasEntry() bool {
	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `SELECT date_fetched FROM dates WHERE id = ?`
	var date string
	err = db.QueryRow(sqlStmt, 1).Scan(&date)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}

		return false
	}

	return true
}
