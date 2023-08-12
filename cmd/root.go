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
	"os"

	"github.com/gabrie30/joke/configs"
	"github.com/gabrie30/joke/pkg/joke"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	count   int
	last    int
	noFetch bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "joke",
	Short: "Tells a joke",
	Long:  `Tells a joke`,
	Run: func(cmd *cobra.Command, args []string) {
		var jokesToTell int
		if cmd.Flags().Changed("count") {
			jokesToTell = count
		} else if cmd.Flags().Changed("last") {
			jokesToTell = last
			configs.LastNJokesToTell = jokesToTell
		} else {
			jokesToTell = 1
		}

		if cmd.Flags().Changed("no-fetch") {
			noFetch = true
		}

		if !noFetch {
			joke.FetchIfNeeded()
		}
		joke.Tell(jokesToTell)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&count, "count", "c", 1, "count of jokes to tell")
	rootCmd.Flags().IntVarP(&last, "last", "l", 0, "count of last n jokes fetched to tell")
	rootCmd.Flags().Bool("no-fetch", noFetch, "skip fetching of new jokes")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
