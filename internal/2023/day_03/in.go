package day_03

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	engine []string

	currentRow  int
	currentLine int
}

func (d *Day) Part1(in string) string {

	// Initialize variables
	sum := 0
	d.currentLine = -1
	d.engine = strings.Split(in, "\n")

	for _, val := range d.engine {
		if val == "" {
			continue
		}

		// Initialize  variables
		strVal := ""
		isSymbolNear := false
		d.currentRow = 0
		d.currentLine++

		// Loop through all characters and find ints
		for _, str := range val {
			if utils.IsInt(string(str)) {
				strVal += string(str)
				if !isSymbolNear {
					// Check if a symbol is near
					isSymbolNear = d.isSymbolNear()
				}
			} else {
				// It's a '.' or a symbole
				if strVal != "" && isSymbolNear {
					logger.Debug("Found value %q", strVal)
					sum += utils.ToInt(strVal)
				}

				// Reset variables
				strVal = ""
				isSymbolNear = false
			}

			d.currentRow += 1
		}

		// End of line
		if strVal != "" && isSymbolNear {
			logger.Debug("Found value %q", strVal)
			sum += utils.ToInt(strVal)
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) isSymbolNear() bool {

	// Left and right
	if d.isSymbole(d.currentRow-1, d.currentLine, ' ') || d.isSymbole(d.currentRow+1, d.currentLine, ' ') {
		return true
	}

	// Top and bottom
	if d.isSymbole(d.currentRow, d.currentLine-1, ' ') || d.isSymbole(d.currentRow, d.currentLine+1, ' ') {
		return true
	}

	// Diagonale
	if d.isSymbole(d.currentRow-1, d.currentLine-1, ' ') || d.isSymbole(d.currentRow-1, d.currentLine+1, ' ') {
		return true
	}
	if d.isSymbole(d.currentRow+1, d.currentLine-1, ' ') || d.isSymbole(d.currentRow+1, d.currentLine+1, ' ') {
		return true
	}

	return false
}

// isGearNear returns weather a * is near and the index of that gear
// (row,line)
func (d *Day) isGearNear() (bool, string) {

	// Left and right
	if d.isSymbole(d.currentRow-1, d.currentLine, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow-1, d.currentLine)
	}
	if d.isSymbole(d.currentRow+1, d.currentLine, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow+1, d.currentLine)
	}

	// Top and bottom
	if d.isSymbole(d.currentRow, d.currentLine-1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow, d.currentLine-1)
	}
	if d.isSymbole(d.currentRow, d.currentLine+1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow, d.currentLine+1)
	}

	// Diagonale
	if d.isSymbole(d.currentRow-1, d.currentLine-1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow-1, d.currentLine-1)
	}
	if d.isSymbole(d.currentRow-1, d.currentLine+1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow-1, d.currentLine+1)
	}
	if d.isSymbole(d.currentRow+1, d.currentLine-1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow+1, d.currentLine-1)
	}
	if d.isSymbole(d.currentRow+1, d.currentLine+1, '*') {
		return true, fmt.Sprintf("%d,%d", d.currentRow+1, d.currentLine+1)
	}

	return false, ""
}

// isSymbole checks if the characters at the given position is a rune.
// If an invalid line or row is provided, this function returns false
func (d *Day) isSymbole(row, line int, symbole rune) bool {
	// Check out of bounds
	if row < 0 || line < 0 || line >= len(d.engine) || row >= len(d.engine[line]) {
		return false
	}

	val := d.engine[line][row]
	return val != '.' && !utils.IsInt(string(val)) && (symbole == ' ' || symbole == rune(val))
}

func (d *Day) Part2(in string) string {
	// Initialize variables
	d.currentLine = -1
	d.engine = strings.Split(in, "\n")
	gears := make(map[string][]int, 0)

	for _, val := range d.engine {
		if val == "" {
			continue
		}

		// Initialize  variables
		strVal := ""
		isGearNear := false
		isGearNearPos := ""
		d.currentRow = 0
		d.currentLine++

		// Loop through all characters and find ints
		for _, str := range val {
			if utils.IsInt(string(str)) {
				strVal += string(str)
				// Check if a symbol is near
				if isNear, pos := d.isGearNear(); isNear {
					isGearNear = true
					isGearNearPos = pos
				}
			} else {
				// It's a '.' or a symbole
				if strVal != "" && isGearNear {
					logger.Debug("Found pos %q with value %q", isGearNearPos, strVal)
					gears[isGearNearPos] = append(gears[isGearNearPos], utils.ToInt(strVal))
				}

				// Reset variables
				strVal = ""
				isGearNear = false
				isGearNearPos = ""
			}

			d.currentRow += 1
		}

		// End of line
		// Check if a symbol is near
		if isGearNear {
			logger.Debug("Found pos %q with value %q", isGearNearPos, strVal)
			gears[isGearNearPos] = append(gears[isGearNearPos], utils.ToInt(strVal))
		}
	}

	sum := 0
	for _, values := range gears {
		if len(values) == 2 {
			sum += values[0] * values[1]
		}
	}

	return fmt.Sprintf("%d", sum)
}
