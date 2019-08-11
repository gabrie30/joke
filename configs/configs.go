package configs

import (
	"log"
	"reflect"

	"github.com/mitchellh/go-homedir"
)

var (
	// DBPath is the path to users database
	DBPath string
	// CreateDatesDB string to create dates database
	CreateDatesDB = "CREATE TABLE IF NOT EXISTS dates (id INTEGER PRIMARY KEY, date_fetched TEXT UNIQUE, jokes_fetched INTEGER)"
	// CreateJokesDB string to create jokes database
	CreateJokesDB = "CREATE TABLE IF NOT EXISTS jokes (id INTEGER PRIMARY KEY, setup TEXT UNIQUE, punchline TEXT, score INTEGER)"
)

func init() {
	DBPath = defaultDBPath()
}

// HomeDir returns users homedir
func HomeDir() string {
	path, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Could not determine users home dir, try again -- err: %v", err)
	}

	return path
}

func isZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

// defaultDBPath provides the default database path at $HOME/.jokes.db
func defaultDBPath() string {

	if isZero(DBPath) == false {
		return DBPath
	}

	path := HomeDir()

	dbPath := path + "/jokes.db"

	return dbPath
}
