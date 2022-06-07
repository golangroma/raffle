package cli

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/golangroma/meetup-20220614/pkg/raffle"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "raffle",
		Short:         "raffle issue owners!",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	rootCmd.AddCommand(NewRunCommand())
	rootCmd.AddCommand(NewListCommand())

	return rootCmd
}

func NewListCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:           "list",
		Short:         "list partecipants",
		Args:          UserRepoArg(),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			userRepo := strings.Split(args[0], "/")
			user, repo := userRepo[0], userRepo[1]

			client := github.NewClient(nil)

			labels, _ := cmd.Flags().GetStringSlice("label")
			issues, err := raffle.GetIssues(client, user, repo, labels)
			if err != nil {
				return err
			}

			users := raffle.GetUsersFromIssues(issues)
			sort.Strings(users)

			fmt.Printf("There are %d partecipants:\n", len(users))
			for _, user := range users {
				time.Sleep(time.Millisecond * 500)
				fmt.Printf(" - %s\n", user)
			}

			return nil
		},
	}

	listCmd.Flags().StringSlice("label", []string{}, "labels")

	return listCmd
}

func NewRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:           "run",
		Short:         "run raffle",
		Args:          UserRepoArg(),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			userRepo := strings.Split(args[0], "/")
			user, repo := userRepo[0], userRepo[1]

			client := github.NewClient(nil)

			labels, _ := cmd.Flags().GetStringSlice("label")
			issues, err := raffle.GetIssues(client, user, repo, labels)
			if err != nil {
				return err
			}

			users := raffle.GetUsersFromIssues(issues)

			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(users), func(i, j int) {
				users[i], users[j] = users[j], users[i]
			})

			return nil
		},
	}

	runCmd.Flags().StringSlice("label", []string{}, "labels")

	return runCmd
}

func UserRepoArg() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts only one argument OWNER/REPO")
		}

		userRepo := strings.Split(args[0], "/")
		if len(userRepo) != 2 {
			return fmt.Errorf("accepts only one argument OWNER/REPO")
		}

		return nil
	}
}
