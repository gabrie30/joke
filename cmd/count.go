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
	"fmt"
	"log"

	"github.com/gabrie30/joke/pkg/joke"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Returns the number of jokes saved in your database ($HOME/.jokes.db)",
	Long:  `Returns the number of jokes saved in your database ($HOME/.jokes.db)`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := joke.Count()
		if err != nil {
			log.Fatalf("Could not count jokes, err: %v", err)
		}

		fmt.Printf("Total Jokes: %v\n", c)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
