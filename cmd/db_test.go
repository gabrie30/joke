package cmd

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gabrie30/joke/configs"
)

func TestSetup(t *testing.T) {
	file, err := ioutil.TempFile("", "test_joke_db")

	if err != nil {
		log.Fatal(err)
	}

	if err := os.Chmod(file.Name(), 0644); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	configs.DBPath = file.Name()
	setup()

	t.Run("setup should setup jokes table", func(tt *testing.T) {

		db, err := sql.Open("sqlite3", configs.DBPath)
		defer db.Close()
		if err != nil {
			tt.Fatalf("Could not open jokes database, err: %v", err)
		}

		_, err = db.Query("SELECT * FROM jokes")
		if err != nil {
			tt.Fatalf("Jokes table should exist but does not, err: %v", err)
		}

		score := 10
		_, err = db.Exec("INSERT INTO jokes (setup, punchline, score) VALUES (?, ?, ?)", "Why did the chicken cross the road, again?", "To get to the other side.", score)

		if err != nil {
			tt.Fatal(err)
		}

		id := 1
		row := db.QueryRow("SELECT punchline,favorite FROM jokes WHERE id=?", id)

		var punchline string
		var favorite bool

		row.Scan(&punchline, &favorite)
		want := "To get to the other side."
		if punchline != want {
			tt.Fatalf("Could not insert joke properly, should have punchline set corretly, got: %v, wanted: %v", punchline, want)
		}

		if favorite != false {
			tt.Fatalf("Could not insert joke properly, should have favorite column set to false by default, got: %v, wanted: %v", favorite, false)
		}

	})

	t.Run("setup should setup dates table", func(tt *testing.T) {

		db, err := sql.Open("sqlite3", configs.DBPath)
		defer db.Close()
		if err != nil {
			tt.Fatalf("Could not open jokes database, err: %v", err)
		}

		_, err = db.Query("SELECT * FROM dates")
		if err != nil {
			tt.Fatalf("Dates table should exist but does not, err: %v", err)
		}
	})

}
