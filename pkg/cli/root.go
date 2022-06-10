package cli

import (
	"fmt"
	"strings"

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
