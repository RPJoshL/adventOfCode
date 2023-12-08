package day_08

import (
	"fmt"
	"regexp"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	inSplit := strings.Split(in, "\n")

	// Initialize variables
	instructions := inSplit[0]
	line := 0
	instructionCounter := 0
	instructionRegex := regexp.MustCompile(`[ \(](?P<primary>\w+)[,\)]`)
	steps := 0

	// Build indexed map for better performance
	instructionMap := make(map[string]int, 0)
	for i, val := range inSplit {
		if i > 1 && len(val) > 2 {
			v := val[0:3]
			instructionMap[v] = i

			// Found start
			if v == "AAA" {
				line = i
			}
		}
	}

	// Find destination
	for {
		// Start by first instruction on end
		if instructionCounter >= len(instructions) {
			instructionCounter = 0
		}

		// Get left or right
		instructionIndex := 0
		if instructions[instructionCounter] == 'R' {
			instructionIndex = 1
		}

		// Get the next instruction to move on
		next := instructionRegex.FindAllStringSubmatch(inSplit[line], -1)[instructionIndex]
		line = instructionMap[next[1]]
		//logger.Debug("Going to %q on line %d", next[1], line)

		steps++
		instructionCounter++

		// Did we reach the end?
		if strings.HasPrefix(next[1], "ZZZ") {
			break
		}
	}

	return fmt.Sprintf("%d", steps)
}

func (d *Day) Part2(in string) string {
	inSplit := strings.Split(in, "\n")

	// Initialize variables
	instructions := inSplit[0]
	starts := make([]int, 0)
	stepsByStarts := make([]int, 0)
	instructionRegex := regexp.MustCompile(`[ \(](?P<primary>\w+)[,\)]`)

	// Build map for better performance
	instructionMap := make(map[string]int, 0)
	for i, val := range inSplit {
		if i > 1 && len(val) > 2 {
			v := val[0:3]
			instructionMap[v] = i

			// Found start
			if v[2] == 'A' {
				starts = append(starts, i)
			}
		}
	}

	// Find for each starting point the amount of steps that we need to go to the
	// first 'Z'.
	// The input data makes sure that the way back from a Z Node to another Z Node
	// is always the same as from start to the first Z node.
	// How would you find out without looping the first time!
	for _, val := range starts {
		steps := 0
		line := val
		instructionCounter := 0
		for {
			// Start by first instruction on end
			if instructionCounter >= len(instructions) {
				instructionCounter = 0
			}

			// Get left or right
			instructionIndex := 0
			if instructions[instructionCounter] == 'R' {
				instructionIndex = 1
			}

			next := instructionRegex.FindAllStringSubmatch(inSplit[line], -1)[instructionIndex]
			line = instructionMap[next[1]]
			//logger.Debug("%d. Going to %q on line %d", i, next[1], line)

			steps++
			instructionCounter++

			if next[1][2] == 'Z' {
				break
			}
		}
		stepsByStarts = append(stepsByStarts, steps)
	}

	return fmt.Sprintf("%d", utils.CalculateLCM(stepsByStarts...))
}
