package main

import (
	"fmt"

	"github.com/golangroma/raffle/pkg/cli"
)

func main() {
	rootCmd := cli.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
