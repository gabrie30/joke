package configs_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gabrie30/joke/configs"
	"github.com/mitchellh/go-homedir"
)

func TestConfigsDefault(t *testing.T) {

	got := configs.DBPath

	want, err := homedir.Dir()
	if err != nil {
		t.Fatalf("Could not determine users home dir, try again -- err: %v", err)
	}
	want = want + "/.jokes.db"

	if got != want {
		t.Fatalf("Could not correctly setup db path, got: %v, wanted: %v", got, want)
	}

}

func TestConfigsManuallySet(t *testing.T) {
	file, err := ioutil.TempFile("", "test_joke_db")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	configs.DBPath = file.Name()

	want := file.Name()
	got := configs.DBPath

	if got != want {
		t.Fatalf("Could not correctly setup db path, got: %v, wanted: %v", got, want)
	}

}
