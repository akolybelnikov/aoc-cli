package auth

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// GetSession retrieves the session data from a file named ".aoc-session" located in the user's home directory.
// Returns the session data as a string or an error if the file cannot be read or the home directory is inaccessible.
func GetSession() (string, error) {
	sessionPath, err := getSessionPath()
	if err != nil {
		return "", fmt.Errorf("unable to get home directory: %w", err)
	}
	data, err := os.ReadFile(sessionPath)
	if err != nil {
		return "", fmt.Errorf("unable to read session file: %w", err)
	}
	return string(data), nil
}

// SaveSession saves the provided session string to a file in the user's home directory, creating or overwriting it if needed.
// Returns an error if the session file path cannot be resolved or if writing to the file fails.
func SaveSession(session string) error {
	sessionPath, err := getSessionPath()
	if err != nil {
		return fmt.Errorf("unable to get home directory: %w", err)
	}
	err = os.WriteFile(sessionPath, []byte(session), 0600)
	if err != nil {
		return fmt.Errorf("unable to write session file: %w", err)
	}
	return nil
}

// ValidateSession verifies the validity of the given session token for the specified year's Advent of Code events.
// It sends a GET request to the Advent of Code website using the session token and checks the HTTP response status.
// An error is returned if the session is invalid or if an issue occurs during the request process.
func ValidateSession(session string, year int) error {
	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://adventofcode.com/%d/about", year)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return nil
}

// getSessionPath constructs the path to the ".aoc-session" file in the user's home directory and returns it or an error.
func getSessionPath() (string, error) {
	const sessionFile = ".aoc-session"
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, sessionFile), nil
}
