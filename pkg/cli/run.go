package cli

import (
	"strings"

	"github.com/golangroma/raffle/pkg/raffle"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:           "run",
		Short:         "run raffle",
		Args:          UserRepoArg(),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := github.NewClient(nil)

			userRepo := strings.Split(args[0], "/")
			user, repo := userRepo[0], userRepo[1]

			labels, err := cmd.Flags().GetStringSlice("label")
			if err != nil {
				return err
			}

			return raffle.Run(client.Issues, user, repo, labels)
		},
	}

	runCmd.Flags().StringSlice("label", []string{}, "labels")

	return runCmd
}
