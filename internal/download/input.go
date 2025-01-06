package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Input(year, day int, session string) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	destPath := filepath.Join("inputs", fmt.Sprintf("day%02d.txt", day))

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "session="+session)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: %s", url, resp.Status)
	}

	err = os.MkdirAll(filepath.Dir(destPath), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, resp.Body)
	return err
}
