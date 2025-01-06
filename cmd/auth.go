// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/akolybelnikov/aoc-cli/internal/auth"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage your Advent of Code session.",
	Long: `
Advent of Code puzzles and inputs are protected behind a login.
AoC uses a session cookie (session token) to authenticate users.
The CLI checks for a .aoc-session file in the user's home directory.
If the file doesn't exist, the CLI prompts the user to paste their session token (copied from their browser).
This token is saved in .aoc-session for future use.`,
	Run: func(cmd *cobra.Command, args []string) {
		year := time.Now().Year()
		fmt.Printf("Managing session for year %d...\n", year)

		session, err := auth.GetSession()
		if err != nil {
			fmt.Println("No session found. Please paste your AoC session cookie:")
			reader := bufio.NewReader(os.Stdin)
			session, _ = reader.ReadString('\n')
			session = strings.TrimSpace(session)
			if err := auth.SaveSession(session); err != nil {
				fmt.Println("Failed to save session.", err)
				os.Exit(1)
			}
		}

		// Validate the session cookie
		err = auth.ValidateSession(session, year)
		if err != nil {
			fmt.Println("Invalid session. Please paste a fresh AoC session cookie:")
			reader := bufio.NewReader(os.Stdin)
			session, _ = reader.ReadString('\n')
			session = strings.TrimSpace(session)
			if err := auth.ValidateSession(session, year); err != nil {
				fmt.Println("Session validation failed. Exiting.")
				os.Exit(1)
			}

			if err := auth.SaveSession(session); err != nil {
				fmt.Println("Failed to save session.", err)
				os.Exit(1)
			}
		}
		fmt.Println("Session validated and saved successfully.")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
