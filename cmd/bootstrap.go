// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/akolybelnikov/aoc-cli/internal"
	"github.com/akolybelnikov/aoc-cli/internal/auth"
	"github.com/akolybelnikov/aoc-cli/internal/download"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var downloadPath string

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a solution for a specific day",
	Long: `Bootstrap a solution for a specific day. Requires a valid session and a path to the project root.
Downloaded input will be stored in the /inputs directory.`,
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

		if downloadPath == "" {
			fmt.Printf("Please provide a path to the project root.\n")
			return
		}
		cmdPath := filepath.Join(downloadPath, "cmd")

		err := os.MkdirAll(cmdPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory at %s: %v\n", downloadPath, err)
			return
		}

		if err != nil {
			fmt.Printf("Failed to get working directory: %v\n", err)
			return
		}

		dayFolder := filepath.Join(downloadPath, "cmd", fmt.Sprintf("day%02d", day))
		err = copyTemplate(dayFolder)
		if err != nil {
			fmt.Printf("Failed to copy template: %v\n", err)
			return
		}

		fmt.Printf("Package for Day %02d created successfully!\n", day)

		session, err := auth.GetSession()
		if err != nil {
			fmt.Println("Invalid or expired session. Please run auth to update your session.")
			return
		}

		err = auth.ValidateSession(session, downloadYear)
		if err != nil {
			fmt.Println("Invalid session. Please run auth to update your session.")
		}

		err = download.Input(downloadYear, day, session, downloadPath)
		if err != nil {
			fmt.Printf("Failed to download the input: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Input for Day %02d downloaded successfully!\n", day)
	},
}

func init() {
	bootstrapCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of Advent of Code (1-25)")
	bootstrapCmd.Flags().IntVarP(&downloadYear, "year", "y", 0, "Advent of Code year (default: current year)")
	bootstrapCmd.Flags().StringVarP(&downloadPath, "path", "p", "", "Custom path for downloading files")
	rootCmd.AddCommand(bootstrapCmd)
}

// copyTemplate copies files from templatePath to targetPath
func copyTemplate(targetPath string) error {
	dayStr := fmt.Sprintf("%02d", day)

	return fs.WalkDir(internal.Templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root "templates" directory itself
		if path == "templates" {
			return nil
		}

		if d.IsDir() {
			// Mirror the directory structure
			destDir := filepath.Join(targetPath, path)
			return os.MkdirAll(destDir, os.ModePerm)
		}

		// Read file content from embedded template
		content, err := internal.Templates.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace placeholders
		updatedContent := strings.ReplaceAll(string(content), "{{DAY}}", dayStr)
		updatedContent = strings.ReplaceAll(updatedContent, "package templates", fmt.Sprintf("package main"))

		// Handle file renaming logic
		fileName := filepath.Base(path)
		if strings.HasSuffix(fileName, "solution.go") {
			fileName = fmt.Sprintf("day%s.go", dayStr)
		} else if strings.HasSuffix(fileName, "test.go") {
			fileName = fmt.Sprintf("day%s_test.go", dayStr)
		}

		// Determine destination path
		destPath := filepath.Join(targetPath, fileName)

		// Write the updated content to the destination file
		return os.WriteFile(destPath, []byte(updatedContent), os.ModePerm)
	})
}
