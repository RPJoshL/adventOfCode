package day_06

import (
	"fmt"
	"regexp"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {

	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\d+`)
	inSplit := strings.Split(in, "\n")
	rtc := 1

	raceTimes := numberRegex.FindAllString(inSplit[0], -1)
	raceDistances := numberRegex.FindAllString(inSplit[1], -1)

	// Loop through every race and find possible button hold values
	for i := 0; i < len(raceTimes); i++ {
		time := utils.ToInt(raceTimes[i])
		distance := utils.ToInt(raceDistances[i])

		// Brute force every possibility and check if the distance could be bet
		possiblities := 0
		for holding := time; holding > 0; holding -= 1 {
			pDistance := (time - holding) * holding

			if pDistance > distance {
				possiblities += 1
			}
		}

		rtc *= possiblities
	}

	return fmt.Sprintf("%d", rtc)
}

func (d *Day) Part2(in string) string {
	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\d+`)
	inSplit := strings.Split(in, "\n")

	raceTimes := strings.Join(numberRegex.FindAllString(inSplit[0], -1), "")
	raceDistances := strings.Join(numberRegex.FindAllString(inSplit[1], -1), "")

	// Loop through every race and find possible button hold values
	time := utils.ToInt(raceTimes)
	distance := utils.ToInt(raceDistances)

	// Brute force every possibility and check if the distance could be bet
	possiblities := 0
	for holding := time; holding > 0; holding -= 1 {
		pDistance := (time - holding) * holding

		if pDistance > distance {
			possiblities += 1
		}
	}

	return fmt.Sprintf("%d", possiblities)
}
