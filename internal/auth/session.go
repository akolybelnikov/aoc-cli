package auth

import (
	"fmt"
	"os"
	"path/filepath"
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

// getSessionPath constructs the path to the ".aoc-session" file in the user's home directory and returns it or an error.
func getSessionPath() (string, error) {
	const sessionFile = ".aoc-session"
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, sessionFile), nil
}
