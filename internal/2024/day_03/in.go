package day_03

import (
	"fmt"
	"regexp"

	"git.rpjosh.de/RPJosh/go-logger"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	regex := regexp.MustCompile(`mul\(\d*,\d*\)`)

	sum := 0
	for _, exp := range regex.FindAllString(in, -1) {
		var a, b int
		if _, err := fmt.Sscanf(exp, "mul(%d,%d)", &a, &b); err != nil {
			logger.Fatal("Failed to read expression %q: %s", exp, err)
		}

		sum += a * b
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) Part2(in string) string {
	regex := regexp.MustCompile(`mul\(\d*,\d*\)`)
	ignoreRegex := regexp.MustCompile(`don't()`)
	doRegex := regexp.MustCompile(`do()`)

	dos := doRegex.FindAllStringIndex(in, -1)
	ignores := ignoreRegex.FindAllStringIndex(in, -1)
	matchesIndex := regex.FindAllStringIndex(in, -1)

	sum := 0
	for i, exp := range regex.FindAllString(in, -1) {
		var a, b int
		if _, err := fmt.Sscanf(exp, "mul(%d,%d)", &a, &b); err != nil {
			logger.Fatal("Failed to read expression %q: %s", exp, err)
		}

		// Check if it's enabled
		startIndex := matchesIndex[i][0]
		lastDo := 0
		lastIgnore := -1

		for _, do := range dos {
			if do[0] > startIndex {
				break
			} else if do[0] > lastDo {
				lastDo = do[0]
			}
		}
		for _, ignore := range ignores {
			if ignore[0] > startIndex {
				break
			} else if ignore[0] > lastIgnore {
				lastIgnore = ignore[0]
			}
		}

		if lastDo > lastIgnore {
			sum += a * b
		} else {
			logger.Debug("Ignoring %q", exp)
		}
	}

	return fmt.Sprintf("%d", sum)
}
