package day_08

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	// Map with characters for x, y positions
	charMap map[rune][][]int
}

func (d *Day) Part1(in string) string {
	// Parse input
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.parseAntenna(lines)

	// Build antinode map
	ctx := 0
	antinodeMap := make(map[string]int, 0)
	for _, pos := range d.charMap {
		for i, this := range pos {

			// Check anitnodes with each other
			for o, same := range pos {
				if i == o {
					continue
				}

				antiX, antiY := d.getOffsetTo(this, same)
				antiX = this[0] - antiX
				antiY = this[1] - antiY
				key := fmt.Sprintf("%dx%d", antiX, antiY)

				// Check bounds
				if antiX < 0 || antiY < 0 || antiX >= len(lines[0]) || antiY >= len(lines) {
					logger.Debug("Antinode out of bound: %s", key)
					continue
				}
				if _, exists := antinodeMap[key]; exists {
					logger.Debug("Antinode already exists: %s", key)
				} else {
					logger.Debug("Found new antinode at: %s", key)
					ctx++
					antinodeMap[key] = 0
				}

			}
		}
	}

	return fmt.Sprintf("%d", ctx)
}

func (d *Day) parseAntenna(lines []string) {
	d.charMap = make(map[rune][][]int)

	for i, line := range lines {
		for r, c := range line {
			// Not a valid antenna
			if c == '.' || c == '#' {
				continue
			}

			if _, exists := d.charMap[c]; exists {
				d.charMap[c] = append(d.charMap[c], []int{r, i})
			} else {
				d.charMap[c] = [][]int{{r, i}}
			}
		}
	}
}

func (d *Day) getOffsetTo(from []int, to []int) (x, y int) {
	return to[0] - from[0], to[1] - from[1]
}

func (d *Day) Part2(in string) string {
	// Parse input
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.parseAntenna(lines)

	// Build antinode map
	ctx := 0
	antinodeMap := make(map[string]int, 0)
	for _, pos := range d.charMap {
		posResonances := [][]int{}

		for i, this := range pos {

			// Check anitnodes with each other
			for o, same := range pos {
				if i == o {
					continue
				}

				antiX, antiY := d.getOffsetTo(this, same)
				newAntiX := this[0] - antiX
				newAntiY := this[1] - antiY

				key := fmt.Sprintf("%dx%d", newAntiX, newAntiY)

				// Check bounds
				boundsValid := true
				if newAntiX < 0 || newAntiY < 0 || newAntiX >= len(lines[0]) || newAntiY >= len(lines) {
					logger.Debug("Antinode out of bound: %s", key)
					boundsValid = false
				}
				if boundsValid {
					if _, exists := antinodeMap[key]; exists {
						logger.Debug("Antinode already exists: %s", key)
					} else {
						logger.Debug("Found new antinode at: %s", key)
						ctx++
						antinodeMap[key] = 0
					}
				}

				fmt.Println()

				ii := 1
				for {
					newAntiX := this[0] + (antiX * ii)
					newAntiY := this[1] + (antiY * ii)
					key := fmt.Sprintf("%dx%d", newAntiX, newAntiY)

					// Check bounds
					if newAntiX < 0 || newAntiY < 0 || newAntiX >= len(lines[0]) || newAntiY >= len(lines) {
						logger.Debug("[RES] Antinode out of bound: %s", key)
						break
					}
					if _, exists := antinodeMap[key]; exists {
						logger.Debug("[RES] Antinode already exists: %s", key)
					} else {
						logger.Debug("[RES] Found new antinode at: %s", key)
						ctx++
						antinodeMap[key] = 0
					}

					// Add to resonance map
					alreadyExists := false
					for _, ii := range posResonances {
						if ii[0] == newAntiX && ii[1] == newAntiY {
							alreadyExists = true
						}
					}
					if !alreadyExists {
						posResonances = append(posResonances, []int{newAntiX, newAntiY})
					}

					ii++
				}
				fmt.Println()

			}
		}
	}

	return fmt.Sprintf("%d", ctx)
}

// 746 to low
