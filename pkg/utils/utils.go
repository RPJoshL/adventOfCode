package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Copies the given string to the clipboard.
// The tool 'xclip' needs to be installed
func CopyToClipboard(text string) {
	command := fmt.Sprintf("echo '%s' | xclip -selection clipboard", strings.ReplaceAll(text, "'", "''"))
	cmd := exec.Command("bash", "-c", command)

	if err := cmd.Start(); err != nil {
		PrintError("Error while executing command: %s", err)
	}

	fmt.Print(text + "\n")
}

// Gets the content of the clipboard.
// The tool 'xclip' needs to be installed
func GetFromClipboard() string {
	out, err := exec.Command("bash", "-c", "xclip -o -selection clipboard").Output()
	if err != nil {
		PrintError("Error while getting clipboard content: %s", err)
	}

	return string(out)
}

// Get's the input data for the task
func GetInputData(year int, day int) string {
	// Check if the ressource was already requested
	fileLocation := fmt.Sprintf("./inputData/%d/day_%02d", year, day)
	cachedInput, err := os.ReadFile(fileLocation)
	if err == nil {
		return string(cachedInput)
	}

	// Make a request to get the input from advent of code and save it in a file
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	// Read cookie session vaue
	session, err := os.ReadFile("./session.txt")
	if err != nil {
		PrintError("Failed to read cookie session file: %s", err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		PrintError("%s", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: string(session)})

	// Execute the query
	res, err := client.Do(req)
	if err != nil {
		PrintError("Failed to execute request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		PrintError("Incorrect status code: %d", res.StatusCode)
	}
	inputData, _ := io.ReadAll(res.Body)
	err = os.WriteFile(fileLocation, inputData, 0644)
	if err != nil {
		PrintError("Failed to write input data to file: %s", err)
	}

	return string(inputData)
}

func PrintError(format string, a ...any) {
	fmt.Printf(format, a...)
	os.Exit(-1)
}

func ToInt(val string) int {
	value, err := strconv.Atoi(val)
	if err != nil {
		PrintError("%s", err)
	}

	return value
}
