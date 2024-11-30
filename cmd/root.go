package cmd

import (
	githubuserdata "github-user-activity/github"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "to fetch last activity in github repo of a user",
}

func Execute() {
	rootCmd.AddCommand(userCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "to fet user last activity",
	RunE: func(cmd *cobra.Command, args []string) error {
		githubuserdata.Event()
		return nil
	},
}
