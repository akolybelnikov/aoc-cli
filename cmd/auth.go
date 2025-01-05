// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github/akolybelnikov/aoc-cli/internal/auth"
	"os"
	"strings"

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
		session, err := auth.GetSession()
		if err != nil {
			fmt.Println("No session found. Please paste your AoC session cookie:")
			reader := bufio.NewReader(os.Stdin)
			session, _ = reader.ReadString('\n')
			session = strings.TrimSpace(session)
			err = auth.SaveSession(session)
			if err != nil {
				fmt.Println("Failed to save session.", err)
				os.Exit(1)
			}
			fmt.Println("Session saved successfully.")
		} else {
			fmt.Println("Session found. You can download puzzles and inputs with the 'aoc download --day <day>' command.")
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
