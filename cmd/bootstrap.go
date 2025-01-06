// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/akolybelnikov/aoc-cli/internal/auth"
	"github.com/akolybelnikov/aoc-cli/internal/download"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a solution for a specific day",
	Long: `Bootstrap a solution for a specific day. Downloaded input will be stored in the /inputs directory.
It will also create a new directory for the day in the /cmd directory and add a test file for the solution.`,
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

		// Get the current working directory
		cwd, err := os.Getwd()
		fmt.Println(cwd)
		if err != nil {
			fmt.Printf("Failed to get working directory: %v\n", err)
			return
		}

		dayFolder := fmt.Sprintf("%s/cmd/day%02d", cwd, day)
		fmt.Println(dayFolder)
		err = os.MkdirAll(dayFolder, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory: %v\n", err)
			return
		}

		templatePath := "internal/templates/"
		err = copyTemplate(templatePath, dayFolder)
		if err != nil {
			fmt.Printf("Failed to copy template: %v\n", err)
			return
		}
		fmt.Printf("Solution for Day %02d created successfully!\n", day)

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
	bootstrapCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of Advent of Code (1-25)")
	bootstrapCmd.Flags().IntVarP(&downloadYear, "year", "y", 0, "Advent of Code year (default: current year)")
	rootCmd.AddCommand(bootstrapCmd)
}

// copyTemplate copies files from templatePath to targetPath
func copyTemplate(templatePath, targetPath string) error {
	dayStr := fmt.Sprintf("%02d", day)

	return filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(templatePath, path)
		destPath := filepath.Join(targetPath, relPath)

		// Handle file renaming
		if strings.HasSuffix(info.Name(), ".go") {
			if info.Name() == "solution.go" {
				destPath = filepath.Join(targetPath, fmt.Sprintf("day%s.go", dayStr))
			} else if info.Name() == "test.go" {
				destPath = filepath.Join(targetPath, fmt.Sprintf("day%s_test.go", dayStr))
			}
		}

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// Read source file
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func(srcFile *os.File) {
			_ = srcFile.Close()
		}(srcFile)

		content, err := io.ReadAll(srcFile)
		if err != nil {
			return err
		}

		// Replace placeholders
		updatedContent := strings.ReplaceAll(string(content), "{{DAY}}", dayStr)
		updatedContent = strings.ReplaceAll(updatedContent, "package templates", fmt.Sprintf("package main"))

		// If it's the test file, replace TestSolution with TestDayXX
		if strings.Contains(info.Name(), "test.go") {
			updatedContent = strings.ReplaceAll(updatedContent, "TestSolution", fmt.Sprintf("TestDay%s", dayStr))
		}

		// Write to the new destination
		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer func(destFile *os.File) {
			_ = destFile.Close()
		}(destFile)

		_, err = destFile.Write([]byte(updatedContent))
		return err
	})
}
