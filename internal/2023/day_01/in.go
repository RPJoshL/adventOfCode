package day_01

import (
	"fmt"
	"regexp"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

// findFirstAndLastNumber finds the first and last occurence of a number by the given regex
// expression for the number (within a match group).
// If the found string value is present in the map, the mapped value is used
func (d *Day) findFirstAndLastNumber(in string, regex string, stringMap map[string]int) int {
	// Build regex to get the first and last number in the string (because of greediness we get the last number)
	regexFirst := regexp.MustCompile(`^.*?` + regex + `.*`)
	regexLast := regexp.MustCompile(`^.*` + regex)
	count := 0

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		matchesFirst := regexFirst.FindStringSubmatch(val)
		matchesLast := regexLast.FindStringSubmatch(val)

		// We expect exactly at least one match
		if len(matchesFirst) < 1 {
			logger.Fatal("Failed to get at least one number (got %d): %q", len(matchesFirst), val)
		}

		str := d.replaceDigitByInt(matchesFirst[1], stringMap)
		if len(matchesLast) > 1 {
			str += d.replaceDigitByInt(matchesLast[1], stringMap)
		}

		count += utils.ToInt(str)
	}

	return count
}

func (d *Day) replaceDigitByInt(val string, stringMap map[string]int) string {
	if v, exists := stringMap[val]; exists {
		return fmt.Sprintf("%d", v)
	}

	return val
}

func (d *Day) Part1(in string) string {
	return fmt.Sprintf("%d", d.findFirstAndLastNumber(in, `(\d)`, make(map[string]int)))
}

func (d *Day) Part2(in string) string {
	digist := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// Build regex string by looping through all digists
	regex := `(\d`
	for k, _ := range digist {
		regex += "|" + k
	}
	regex += ")"

	return fmt.Sprintf("%d", d.findFirstAndLastNumber(in, regex, digist))
}
