package day_05

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	orderRules [][]int
}

func (d *Day) Part1(in string) string {
	splittedInput := strings.Split(in, "\n\n")

	// Parse rules
	for _, line := range utils.RemoveEmptyLines(strings.Split(splittedInput[0], "\n")) {
		newRules := make([]int, 2)

		if _, err := fmt.Sscanf(line, "%d|%d", &newRules[0], &newRules[1]); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}
		d.orderRules = append(d.orderRules, newRules)
	}

	// Process updates
	sum := 0
	for _, line := range utils.RemoveEmptyLines(strings.Split(splittedInput[1], "\n")) {
		updates := strings.Split(line, ",")
		isValid, middle, _, _ := d.isUpdateValid(updates)

		//logger.Debug("Is Valid %q: %t => +%d\n", updates, isValid, middle)
		if isValid {
			sum += middle
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) isUpdateValid(updates []string) (isValid bool, middle int, invalidIndex int, invalidCount int) {
	isValid = true

	for i, update := range updates {

		// Check if middle page number
		updateNum := utils.ToInt(update)
		if i == (len(updates)-1)/2 {
			middle = updateNum
		}

		// Check if any rules are validated
		for _, rule := range d.orderRules {

			// Boot numbers have to be contained in the update
			if !d.containsNumber(rule[0], updates) || !d.containsNumber(rule[1], updates) {
				continue
			}

			// Check after rule (if there is at least one before it)
			isAfter := rule[1] == updateNum && !d.isNumberBefore(rule[0], i, updates)
			isBefore := rule[0] == updateNum && !d.isNumberAfter(rule[1], i, updates)
			if isAfter || isBefore {
				//logger.Debug("Rule for %d is invalid (rules for %d / %d)", updateNum, rule[0], rule[1])
				isValid = false
				if invalidCount == 0 {
					invalidIndex = i
				}
				invalidCount++
			}
		}
	}

	return
}

func (d *Day) containsNumber(number int, updates []string) bool {
	numberStr := fmt.Sprintf("%d", number)

	for i := 0; i < len(updates); i++ {
		if updates[i] == numberStr {
			return true
		}
	}

	return false
}

func (d *Day) isNumberBefore(number, numberIndex int, updates []string) bool {
	numberStr := fmt.Sprintf("%d", number)

	for i := 0; i < numberIndex; i++ {
		if updates[i] == numberStr {
			return true
		}
	}

	return false
}

func (d *Day) isNumberAfter(number, numberIndex int, updates []string) bool {
	numberStr := fmt.Sprintf("%d", number)

	for i := numberIndex; i < len(updates); i++ {
		if updates[i] == numberStr {
			return true
		}
	}

	return false
}

func (d *Day) Part2(in string) string {
	splittedInput := strings.Split(in, "\n\n")

	// Parse rules
	for _, line := range utils.RemoveEmptyLines(strings.Split(splittedInput[0], "\n")) {
		newRules := make([]int, 2)

		if _, err := fmt.Sscanf(line, "%d|%d", &newRules[0], &newRules[1]); err != nil {
			logger.Fatal("Failed to read line %q: %s", line, err)
		}
		d.orderRules = append(d.orderRules, newRules)
	}

	// Process updates
	sum := 0
	for _, line := range utils.RemoveEmptyLines(strings.Split(splittedInput[1], "\n")) {
		updates := strings.Split(line, ",")
		isValid, middle, _, invCount := d.isUpdateValid(updates)

		if !isValid {
			newUpdates := make([]string, len(updates))
			copy(newUpdates, updates)

			// Brute force the correct ordering
			fromIndex := 0
			toIndex := 0
			for {
				if fromIndex >= len(newUpdates)-1 {
					//logger.Fatal("Failed to correct line %q => %q", line, newUpdates)
					toIndex = 0
					fromIndex = 0
				} else if toIndex >= len(newUpdates) {
					toIndex = 0
					fromIndex++
				}

				tmpLastIndex := invCount
				tt := utils.MoveSliceElement(newUpdates, fromIndex, toIndex)
				isValid, middle, _, invCount = d.isUpdateValid(tt)

				if isValid {
					logger.Debug("Adding middle sum of %d", middle)
					sum += middle
					break
				} else if invCount < tmpLastIndex {
					newUpdates = tt
					//toIndex = 0
					//fromIndex = 0
				} else {
					//logger.Debug("Tries was not successfull (old = %d, new = %d, to = %d, from = %d): %q", tmpLastIndex, invCount, toIndex, fromIndex, tt)
					invCount = tmpLastIndex
					toIndex++
				}
			}

		}
	}

	return fmt.Sprintf("%d", sum)
}
