// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Bootstrapping and session management for Advent of Code",
	Long: `
Aoc-cli is a CLI tool for bootstrapping daily challenges for Advent of Code.
It optionally manages your session for you, which allows you to download the input data.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
