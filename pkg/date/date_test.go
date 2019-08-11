package date_test

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gabrie30/joke/configs"
	"github.com/gabrie30/joke/pkg/date"
	_ "github.com/mattn/go-sqlite3"
)

func TestSaveDate(t *testing.T) {
	dbSetup()
	defer os.Remove(configs.DBPath)

	d := date.NewDate()
	d.DateFetched = "11-11-1111"
	d.JokesFetched = 10

	err := d.Save()

	if err != nil {
		log.Fatalf("Could not save date; err: %v", err)
	}

	db, err := sql.Open("sqlite3", configs.DBPath)
	defer db.Close()
	var datef string
	got := db.QueryRow("SELECT date_fetched FROM dates WHERE jokes_fetched=10")
	got.Scan(&datef)
	want := "11-11-1111"

	if datef != want {
		log.Fatalf("Could not save joke; got: %s, wanted: %s", datef, want)
	}
}

func TestLastFetchDate(t *testing.T) {
	dbSetup()
	defer os.Remove(configs.DBPath)

	neverFetched := date.LastFetchDate()

	if neverFetched != "" {
		log.Fatalf("if no dates exist last fetch date should return an empty string instead got: %v", neverFetched)
	}

	d := date.NewDate()
	d.DateFetched = "11-11-1111"
	d.JokesFetched = 10

	err := d.Save()

	if err != nil {
		log.Fatalf("Could not save date; err: %s", err.Error())
	}

	dz := date.LastFetchDate()

	if dz != "11-11-1111" {
		log.Fatalf("last fetch date was not successful; got %v, wanted: 11-11-1111", d)
	}
}

func TestHasEntry(t *testing.T) {
	dbSetup()
	defer os.Remove(configs.DBPath)

	if date.HasEntry() != false {
		log.Fatal("Should return false if no dates exist in dates table, got: true, wanted: false")
	}

	d := date.NewDate()

	d.DateFetched = "11-11-1111"
	d.JokesFetched = 10

	err := d.Save()

	if err != nil {
		log.Fatalf("Could not save date; err: %v", err)
	}

	if date.HasEntry() != true {
		log.Fatal("Should return true when dates exist in dates table, got: false, wanted: true")
	}
}

func dbSetup() {
	testDb := configs.HomeDir()
	file, err := ioutil.TempFile(testDb, "test_joke_db")

	if err != nil {
		log.Fatal(err)
	}

	configs.DBPath = file.Name()
	db, err := sql.Open("sqlite3", configs.DBPath)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not open jokes database, err: %v", err)
	}

	statement, err := db.Prepare(configs.CreateJokesDB)

	if err != nil {
		log.Fatalf("Could not prepare jokes database, err: %v", err)
	}

	_, err = statement.Exec()

	if err != nil {
		log.Fatalf("Could not create jokes database, err: %v", err)
	}

	statement, err = db.Prepare(configs.CreateDatesDB)

	if err != nil {
		log.Fatalf("Could not prepare dates database, err: %v", err)
	}

	_, err = statement.Exec()

	if err != nil {
		log.Fatalf("Could not create dates database, err: %v", err)
	}
}
