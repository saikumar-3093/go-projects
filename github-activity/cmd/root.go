package cmd

import (
	"errors"

	"github-activity/github"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Github User Activity is a CLI tool for fetching user activity",
	Long: `Github User Activity is a CLI tool for fetching user activity. It allows you to fetch user activity by providing the username.

Example:
> github-activity saikumar-3093

Complete code available at "https://github.com/saikumar-3093/github-user-activity-cli-app"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("\033[31mplease provide username\033[0m")

		}
		user := args[0]
		return github.Event(user)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
