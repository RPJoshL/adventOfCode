package day_02

import (
	"fmt"
	"regexp"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {

	// This map contains the maximum balls that are in the bag
	balls := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Regex to get the game number
	gameNumberRegex := regexp.MustCompile(`^Game (?P<gameID>\w*)`)
	ballsRegex := regexp.MustCompile(`(?P<balls>\w* (blue|red|green))(,|;)? ?`)

	rtc := 0
	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		// Get the game ID
		game := utils.ToInt(gameNumberRegex.FindStringSubmatch(val)[1])
		possible := true

		for _, ballAll := range ballsRegex.FindAllStringSubmatch(val, -1) {
			ballValues := strings.Split(ballAll[1], " ")

			// Extract count and color
			count := utils.ToInt(ballValues[0])
			color := ballValues[1]

			logger.Debug("%d. Found balls %q", game, ballAll[1])

			if count > balls[color] {
				possible = false
				break
			}
		}

		// Add the game to the counter
		if possible {
			rtc += game
		}
	}

	return fmt.Sprintf("%d", rtc)
}

func (d *Day) Part2(in string) string {
	// Regex to get the game number
	gameNumberRegex := regexp.MustCompile(`^Game (?P<gameID>\w*)`)
	ballsRegex := regexp.MustCompile(`(?P<balls>\w* (blue|red|green))(,|;)? ?`)

	rtc := 0
	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		// Get the game ID
		game := utils.ToInt(gameNumberRegex.FindStringSubmatch(val)[1])

		// Maximum balls count present in the game
		balls := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, ballAll := range ballsRegex.FindAllStringSubmatch(val, -1) {
			ballValues := strings.Split(ballAll[1], " ")

			// Extract count and color
			count := utils.ToInt(ballValues[0])
			color := ballValues[1]

			logger.Debug("%d. Found balls %q", game, ballAll[1])

			if count > balls[color] {
				balls[color] = count
			}
		}

		// Add the power to the result
		rtc += balls["red"] * balls["green"] * balls["blue"]
	}

	return fmt.Sprintf("%d", rtc)
}
