package day_04

import (
	"fmt"
	"regexp"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
)

type Day struct{}

func (d *Day) Part1(in string) string {

	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\w*`)

	sum := 0
	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		// Extract card numbers
		numbers := strings.Split(val, ":")[1]
		winning := strings.Split(numbers, "|")[0]
		elfs := strings.Split(numbers, "|")[1]

		// Loop through all numbers
		points := 0
		for _, number := range numberRegex.FindAllString(elfs, -1) {
			if number == "" {
				continue
			}

			if strings.Contains(winning, " "+number+" ") {
				logger.Debug("Found matching number: %q (%s)", number, winning)

				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) Part2(in string) string {
	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\w*`)

	cards := 0
	numberOfCards := make(map[int]int, 0)
	for i, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		// Increment cards count
		cards++
		replicas := 0
		if val, exists := numberOfCards[i]; exists {
			logger.Info("Cards of game %d: %d", i+1, val)
			cards += val
			replicas = val
		}

		// Extract card numbers
		numbers := strings.Split(val, ":")[1]
		winning := strings.Split(numbers, "|")[0]
		elfs := strings.Split(numbers, "|")[1]

		// Loop through all numbers
		points := 0
		for _, number := range numberRegex.FindAllString(elfs, -1) {
			if number == "" {
				continue
			}

			if strings.Contains(winning, " "+number+" ") {
				logger.Debug("Found matching number: %q (%s)", number, winning)

				// Save number of cards
				points++
				numberOfCards[i+points] += 1
				// Each card wins that also
				numberOfCards[i+points] += replicas
			}
		}
	}

	return fmt.Sprintf("%d", cards)
}
