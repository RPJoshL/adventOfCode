package day_06

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Day struct {
	input     []string
	xPosition int
	yPosition int

	direction int

	// distinct positions
	positionCount int
}

func (d *Day) Part1(in string) string {
	d.input = utils.RemoveEmptyLines(strings.Split(in, "\n"))

	// Get starting position
	d.applyStartingPosition(true)

	for {
		nextX, nextY := d.getNextPosition()

		// Out of bounds
		if nextX < 0 || nextY < 0 || nextX >= len(d.input[0]) || nextY >= len(d.input) {
			break
		}

		// Check if next position is a blockade
		//logger.Debug("Moving to x = %d, y = %d, dir = %d", nextX, nextY, d.direction)
		if d.input[nextY][nextX] == '#' {
			logger.Debug("Found blockade at x = %d, y = %d", nextX, nextY)
			if d.direction == LEFT {
				d.direction = 0
			} else {
				d.direction += 1
			}
		} else {
			d.xPosition = nextX
			d.yPosition = nextY

			if d.input[nextY][nextX] != 'X' {
				d.positionCount++
				d.input[nextY] = utils.ReplaceCharacterInString(d.input[nextY], "X", nextX)
			}
		}
	}

	return fmt.Sprintf("%d", d.positionCount)
}

func (d *Day) applyStartingPosition(replaceStart bool) {
	for l, line := range d.input {
		for r, c := range line {
			if c == '^' {
				d.xPosition = r
				d.yPosition = l

				if replaceStart {
					d.input[l] = utils.ReplaceCharacterInString(d.input[l], "X", r)
				}
				break
			}
		}
	}
}

func (d *Day) getNextPosition() (x, y int) {
	x = d.xPosition
	y = d.yPosition

	if d.direction == UP {
		y -= 1
	} else if d.direction == DOWN {
		y += 1
	} else if d.direction == LEFT {
		x -= 1
	} else if d.direction == RIGHT {
		x += 1
	}

	return
}

func (d *Day) Part2(in string) string {
	d.input = utils.RemoveEmptyLines(strings.Split(in, "\n"))

	// Get starting position
	d.applyStartingPosition(false)
	origX := d.xPosition
	origY := d.yPosition

	loopCounter := 0
	for i := 0; i < len(d.input); i++ {
		for r := 0; r < len(d.input[i]); r++ {
			if i == origY && r == origX {
				logger.Debug("Found orig position: x = %d, y = %d", origX, origY)
				continue
			} else if d.input[i][r] == '#' {
				// Already a blockade
				continue
			}

			// Reset
			d.applyStartingPosition(false)
			d.direction = UP

			// Set blockade
			d.input[i] = utils.ReplaceCharacterInString(d.input[i], "#", r)

			// Positions which where already entered with direction
			visitedMap := make(map[string][]int, 0)

			isLoop := false
			for {
				nextX, nextY := d.getNextPosition()

				// Out of bounds
				if nextX < 0 || nextY < 0 || nextX >= len(d.input[0]) || nextY >= len(d.input) {
					break
				}

				// Check if next position is a blockade
				//logger.Debug("Moving to x = %d, y = %d, dir = %d", nextX, nextY, d.direction)
				if d.input[nextY][nextX] == '#' {
					//logger.Debug("Found blockade at x = %d, y = %d", nextX, nextY)
					if d.direction == LEFT {
						d.direction = 0
					} else {
						d.direction += 1
					}
				} else {
					d.xPosition = nextX
					d.yPosition = nextY
				}

				// Check if already visited
				key := fmt.Sprintf("%d-%d", d.xPosition, d.yPosition)
				if directions, exists := visitedMap[key]; exists {
					for _, dir := range directions {
						if dir == d.direction {
							isLoop = true
							logger.Debug("Found loop for x = %d, y = %d", r, i)
							break
						} else {
							visitedMap[key] = append(visitedMap[key], d.direction)
						}
					}

					if isLoop {
						break
					}
				} else {
					visitedMap[key] = []int{d.direction}
				}
			}

			if isLoop {
				loopCounter++
			}

			// Undo added blockade
			d.input[i] = utils.ReplaceCharacterInString(d.input[i], ".", r)
		}
	}

	return fmt.Sprintf("%d", loopCounter)
}
