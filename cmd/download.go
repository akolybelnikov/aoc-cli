// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"fmt"
	"github/akolybelnikov/aoc-cli/internal/auth"
	"github/akolybelnikov/aoc-cli/internal/download"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var day int
var downloadYear int

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download Advent of Code input for a specific day",
	Long: `
Download Advent of Code input for a specific day. The year defaults to the current year, and the day defaults to the current day.
Pass the --year flag to specify a different year, and the --day flag to specify a specific day.
Example: aoc download --day 1 --year 2025
`,
	Run: func(cmd *cobra.Command, args []string) {
		if downloadYear == 0 {
			downloadYear = time.Now().Year()
		}
		if day == 0 {
			day = time.Now().Day()
		}
		if day < 1 || day > 25 {
			fmt.Println("Invalid day. Please choose a day between 1 and 25.")
			return
		}

		fmt.Printf("Downloading the input for Day %02d, %d...\n", day, downloadYear)

		session, err := auth.GetSession()
		if err != nil {
			fmt.Println("Invalid or expired session. Please run auth to update your session.")
			return
		}

		err = auth.ValidateSession(session, downloadYear)
		if err != nil {
			fmt.Println("Invalid session. Please run auth to update your session.")
		}

		err = download.Input(downloadYear, day, session)
		if err != nil {
			fmt.Printf("Failed to download the input: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Input for Day %02d downloaded successfully!\n", day)
	},
}

func init() {
	downloadCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of Advent of Code (1-25)")
	downloadCmd.Flags().IntVarP(&downloadYear, "year", "y", 0, "Advent of Code year (default: current year)")
	rootCmd.AddCommand(downloadCmd)
}
