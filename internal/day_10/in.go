package day_10

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	sumSignals := 0
	cycle := 0
	sumValue := 1

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		iterator := 1
		if val[0:4] == "addx" {
			iterator = 2
		}

		for i := 0; i < iterator; i++ {
			cycle++
			if (cycle-20)%40 == 0 {
				//fmt.Printf("Cycle: %d | Strength: %d\n", cycle, sumValue)
				sumSignals += cycle * sumValue
			}
		}

		if val[0:4] == "addx" {
			valToAdd, _ := strconv.Atoi(val[5:])
			sumValue += valToAdd
		}
	}

	return fmt.Sprintf("%d", sumSignals)
}

func (d *Day) Part2(in string) string {
	//var screen [][]string = make([]string, 200)
	//for i := range screen {
	//	screen [i] = append(screen, 39)
	//}

	cycle := 0
	sumValue := 1
	lastSpriteRow := 0

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}

		iterator := 1
		if val[0:4] == "addx" {
			iterator = 2
		}

		for i := 0; i < iterator; i++ {
			// The middle of the sprite
			spriteColumn := sumValue % 40

			cyclePos := cycle % 40
			if cyclePos-1 == spriteColumn || cyclePos == spriteColumn || cyclePos+1 == spriteColumn {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			cycle++
			// Print new line
			spriteRow := cycle / 40
			if lastSpriteRow != spriteRow {
				fmt.Print("\n")
				lastSpriteRow = spriteRow
			}
		}

		if val[0:4] == "addx" {
			valToAdd, _ := strconv.Atoi(val[5:])
			sumValue += valToAdd
		}
	}

	return ""
}
