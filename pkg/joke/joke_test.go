package joke_test

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gabrie30/joke/configs"
	"github.com/gabrie30/joke/pkg/joke"
	_ "github.com/mattn/go-sqlite3"
)

func TestSaveJoke(t *testing.T) {
	dbSetup()
	defer os.Remove(configs.DBPath)

	j := joke.New()

	j.Setup = "Why did the chicken cross the road?"
	j.Punchline = "To get to the other side."
	j.Score = 10

	err := j.Save()

	if err != nil {
		log.Fatalf("Could not save joke; err: %v", err)
	}

	db, err := sql.Open("sqlite3", configs.DBPath)
	defer db.Close()
	var pline string
	got := db.QueryRow("SELECT punchline FROM jokes WHERE punchline='To get to the other side.'")
	got.Scan(&pline)
	want := "To get to the other side."

	if pline != want {
		log.Fatalf("Could not save joke; got: %s, wanted: %s", pline, want)
	}
}

func TestCount(t *testing.T) {
	dbSetup()
	defer os.Remove(configs.DBPath)

	j := joke.New()

	j.Setup = "Why did the chicken cross the road?"
	j.Punchline = "To get to the other side."
	j.Score = 10

	err := j.Save()

	j2 := joke.New()

	j2.Setup = "Why did the chicken cross the road? Twice."
	j2.Punchline = "To get to the other side, again."
	j2.Score = 11

	err = j2.Save()

	if err != nil {
		log.Fatalf("Could not save joke; err: %v", err)
	}

	c, err := joke.Count()
	if err != nil {
		log.Fatalf("Could not count jokes err: %v", err)
	}

	if c != 2 {
		log.Fatalf("Count should return 2 when there are two jokes in the database, got: %v, wanted: %v", c, 2)
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
