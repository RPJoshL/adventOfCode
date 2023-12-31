package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
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
	url := fmt.Sprintf("https://adventOfCode.com/%d/day/%d/input", year, day)

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
		panic(fmt.Sprintf("Failed to  convert %q to a number: %s", val, err))
	}

	return value
}

func IsInt(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

// GetMinValue returns the smallest number
// within the given array
func GetMinValue(values []int) int {
	if len(values) == 0 {
		return 0
	}

	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}

	return min
}

// ConvertToInt converts each number within
// []string to a number and returns the resulting array.
// Empty string values are ignored
func ConvertArrayToInt(values []string) []int {
	rtc := make([]int, 0)

	// Convert each value
	for _, nmbString := range values {
		if nmbString == "" {
			continue
		}

		nmb := ToInt(nmbString)
		rtc = append(rtc, nmb)
	}

	return rtc
}

// SortRunes sorts a string based on the rune value
// of every character
func SortRunes(value string) string {
	runeSlice := []rune(value)
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}

// ReplaceCharacterInString replaces the character at the given index with
// the provided value
func ReplaceCharacterInString(str, replace string, index int) string {
	return str[:index] + replace + str[index+1:]
}

// CalculateGCD (greates common divisor) calculates the biggest
// positive integer that is divisible by a and b
func CalculateGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// CalculateLCM (least common multiple) calculates the smallest
// positive integer that is divisible by all numbers
func CalculateLCM(ints ...int) int {
	if len(ints) == 1 {
		return ints[0]
	} else if len(ints) == 2 {
		return ints[0] * ints[1] / CalculateGCD(ints[0], ints[1])
	}

	return CalculateLCM(ints[0], CalculateLCM(ints[1:]...))
}

func AreAllElementsEqual[T any](vals []T) bool {
	if len(vals) == 0 {
		return true
	}

	prev := vals[0]
	for i := 0; i < len(vals); i++ {
		if any(prev) != any(vals[i]) {
			return false
		}

		prev = vals[i]
	}

	return true
}
