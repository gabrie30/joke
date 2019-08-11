package joke

import (
	"database/sql"
	"log"

	"github.com/gabrie30/joke/configs"
	// needed for sqlite
	_ "github.com/mattn/go-sqlite3"
)

var charsNotAllowed = "#@%^*()"

// Data holds the data of a joke, not all fields will necessarily be saved to the database
type Data struct {
	ID              int    `json:"-"`
	Setup           string `json:"title"`
	Punchline       string `json:"selftext"`
	WhitelistStatus string `json:"whitelist_status"`
	Score           int    `json:"ups"`
	AdultsOnly      bool   `json:"over_18"`
	IsVideo         bool   `json:"is_video"`
}

// payload is a reddit json payload
type payload struct {
	RedditDataPayload childrenPayload `json:"data"`
}

type childrenPayload struct {
	RedditChild []children `json:"children"`
}

type children struct {
	RedditData Data `json:"data"`
}

// New creates a new joke
func New() *Data {
	return &Data{}
}

// Count returns the number of jokes saved in the database
func Count() (int, error) {
	var count int
	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(DISTINCT setup) FROM jokes")
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Save saves a joke to the db
func (j *Data) Save() error {

	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO jokes (setup, punchline, score) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(j.Setup, j.Punchline, j.Score)

	if err != nil {
		return err
	}

	return nil
}

// GetJokeByID gets a joke by its ID
func GetJokeByID(id int) *Data {
	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `SELECT id,setup,punchline,score FROM jokes WHERE id=?`
	d := New()
	err = db.QueryRow(sqlStmt, id).Scan(&d.ID, &d.Setup, &d.Punchline, &d.Score)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}

		log.Fatal("Could not find joke with that ID, aborting.")
	}

	return d

}

// LastJokeID returns id of the last joke
func LastJokeID() int {
	db, err := sql.Open("sqlite3", configs.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `SELECT id FROM jokes ORDER BY id DESC LIMIT 1`
	var lastID int
	err = db.QueryRow(sqlStmt).Scan(&lastID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}

		log.Fatal("Could not determine the last joke ID, aborting.")
	}

	return lastID
}
