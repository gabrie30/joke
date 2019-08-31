package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// favoritesCmd represents the favorites command
var favoritesCmd = &cobra.Command{
	Use:   "favorites",
	Short: "Tell one of your favorite jokes",
	Long:  `Tell one of your favorite jokes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("favorites called")
	},
}

func init() {
	rootCmd.AddCommand(favoritesCmd)
	// TODO: add a --count to tell many of your favorites
}
