package day_09

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	return d.paseHistories(in, d.extrapolite)
}

func (d *Day) paseHistories(in string, extrapolite func(vals [][]int) int) string {
	histories := strings.Split(in, "\n")

	sumExtrapolite := 0
	for _, history := range histories {
		if history == "" {
			continue
		}

		differences := make([][]int, 0)

		// Calculate differences until we only have zeros
		vals := utils.ConvertArrayToInt(strings.Split(history, " "))
		differences = append(differences, vals)
		for {
			//logger.Debug("%s", differences)
			// Do we only have zeros?
			if utils.AreAllElementsEqual(vals) && vals[0] == 0 {
				sumExtrapolite += extrapolite(differences)
				logger.Debug("%d", extrapolite(differences))
				break
			}

			// Get the new diffs
			vals = d.calculateDiffs(vals)
			differences = append(differences, vals)
		}
	}

	return fmt.Sprintf("%d", sumExtrapolite)
}

// calculateDiffs calculates the difference between two
// numbers in the provided array and returns a string
// containing the differences concatenated by a space
func (d *Day) calculateDiffs(vals []int) (rtc []int) {
	for i := 0; i < len(vals)-1; i += 1 {
		rtc = append(rtc, vals[i+1]-vals[i])
	}

	return rtc
}

func (d *Day) extrapolite(vals [][]int) int {
	lastVal := 0
	for i := len(vals) - 2; i >= 0; i-- {
		lastVal = lastVal + vals[i][len(vals[i])-1]
	}

	return lastVal
}

func (d *Day) extrapoliteFirst(vals [][]int) int {
	lastVal := 0
	for i := len(vals) - 2; i >= 0; i-- {
		lastVal = vals[i][0] - lastVal
	}

	return lastVal
}

func (d *Day) Part2(in string) string {
	return d.paseHistories(in, d.extrapoliteFirst)
}
