package day_01

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	// Read input
	for i, line := range lines {
		if _, err := fmt.Sscanf(line, "%d   %d", &left[i], &right[i]); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}
	}

	// Sort lists
	slices.Sort(left)
	slices.Sort(right)

	// Calc diff between lists
	diff := 0
	for i := range lines {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}

	return fmt.Sprintf("%d", diff)
}

func (d *Day) Part2(in string) string {
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	// Read input
	for i, line := range lines {
		if _, err := fmt.Sscanf(line, "%d   %d", &left[i], &right[i]); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}
	}

	// Calculate score
	score := 0
	for _, l := range left {
		occurrences := 0
		for _, r := range right {
			if l == r {
				occurrences++
			}
		}

		score += l * occurrences
	}

	return fmt.Sprintf("%d", score)
}
