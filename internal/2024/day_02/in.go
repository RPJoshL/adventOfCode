package day_02

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	safeReports := 0
	for _, line := range utils.RemoveEmptyLines(strings.Split(in, "\n")) {

		// Parse levels
		levels := make([]int, strings.Count(line, " ")+1)
		placeholder := make([]any, len(levels))
		format := ""
		for i := 0; i < len(levels); i++ {
			if i != 0 {
				format += " "
			}

			format += "%d"
			placeholder[i] = &levels[i]
		}
		if _, err := fmt.Sscanf(line, format, placeholder...); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}

		increasing := false
		for i := 1; i < len(levels); i++ {
			last := levels[i-1]
			diff := levels[i] - last

			if i == 1 {
				increasing = diff > 0

				if diff == 0 || diff < -3 || diff > 3 {
					break
				}
			} else if increasing && (diff <= 0 || diff > 3) {
				break
			} else if !increasing && (diff >= 0 || diff < -3) {
				break
			}

			if i == len(levels)-1 {
				//logger.Info("Level %q is safe", line)
				safeReports++
			}
		}
	}

	return fmt.Sprintf("%d", safeReports)
}

func (d *Day) Part2(in string) string {
	safeReports := 0
	for _, line := range utils.RemoveEmptyLines(strings.Split(in, "\n")) {

		// Parse levels
		levels := make([]int, strings.Count(line, " ")+1)
		placeholder := make([]any, len(levels))
		format := ""
		for i := 0; i < len(levels); i++ {
			if i != 0 {
				format += " "
			}

			format += "%d"
			placeholder[i] = &levels[i]
		}
		if _, err := fmt.Sscanf(line, format, placeholder...); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}

		second := []int{levels[0]}
		second = append(second, levels[2:]...)
		if d.IsReportSafe(levels, 0) || d.IsReportSafe(levels[1:], 1) || d.IsReportSafe(second, 1) {
			safeReports++
		}
	}

	return fmt.Sprintf("%d", safeReports)
}

func (d *Day) IsReportSafe(levels []int, unsafeReports int) bool {
	increasing := false
	wasLastUnsafe := false

	for i := 1; i < len(levels); i++ {
		last := levels[i-1]
		if wasLastUnsafe {
			last = levels[i-1-1]
		}
		diff := levels[i] - last

		if i == 1 {
			increasing = diff > 0

			if diff == 0 || diff < -3 || diff > 3 {
				unsafeReports++
				wasLastUnsafe = true
			}
		} else if increasing && (diff <= 0 || diff > 3) {
			unsafeReports++
			wasLastUnsafe = true
		} else if !increasing && (diff >= 0 || diff < -3) {
			unsafeReports++
			wasLastUnsafe = true
		}

		// Allow a single "unsafe" report
		if unsafeReports > 1 {
			break
		}

		if i == len(levels)-1 {
			//logger.Info("Level %q is safe", line)
			return true
		}
	}

	return false
}
