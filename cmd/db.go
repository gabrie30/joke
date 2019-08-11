/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gabrie30/joke/configs"
	// used for sqlite
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// dbinitCmd represents the dbinit command
var db = &cobra.Command{
	Use:   "db",
	Short: "Database related commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("See db subcommands")
	},
}

func init() {
	rootCmd.AddCommand(db)
	db.AddCommand(setupCmd)
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up your joke sqlite database at $HOME/.joke.db",
	Long:  `Sets up your joke sqlite database at $HOME/.joke.db`,
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func setup() {
	if _, err := os.Stat(configs.DBPath); os.IsNotExist(err) {
		fmt.Printf("Setting up jokes database at %s", configs.DBPath)
		os.Create(configs.DBPath)
	} else {
		fmt.Printf("Jokes database detected at %s -- Manually remove this file to start fresh.\n", configs.DBPath)
	}

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
