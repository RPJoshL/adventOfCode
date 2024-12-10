package day_10

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	topoMap [][]int
}

type cordinat struct {
	x      int
	y      int
	height int
}

func (c cordinat) getKey() string {
	return fmt.Sprintf("%dx%d", c.x, c.y)
}

func (d *Day) Part1(in string) string {
	d.parseMap(utils.RemoveEmptyLines(strings.Split(in, "\n")))

	// Start processing of tracks with height 0
	sumScore := 0
	for y := range d.topoMap {
		for x := range d.topoMap[y] {
			co := cordinat{x: x, y: y, height: d.topoMap[y][x]}
			if co.height == 0 {
				sumScore += d.getScoreOfTrack(
					co,
					make(map[string]int),
					make(map[string]int),
					false,
				)
			}
		}
	}

	return fmt.Sprintf("%d", sumScore)
}

func (d *Day) getScoreOfTrack(current cordinat, visited map[string]int, allEndings map[string]int, forRating bool) (sum int) {

	// Mark current position as visited
	newVisted := utils.CopyMap(visited)
	newVisted[current.getKey()] = 1

	newPositions := []cordinat{
		{x: current.x - 1, y: current.y}, // Left
		{x: current.x + 1, y: current.y}, // Right
		{x: current.x, y: current.y + 1}, // Top
		{x: current.x, y: current.y - 1}, // Bottom
	}
	nextHeight := current.height + 1

	// Check for valid directions
	for i, p := range newPositions {

		// Invalid
		if p.y < 0 || p.y >= len(d.topoMap) || p.x < 0 || p.x >= len(d.topoMap[p.y]) {
			continue
		}

		// Check height
		newPositions[i].height = d.topoMap[p.y][p.x]
		if newPositions[i].height != nextHeight {
			continue
		}

		// Check if already visited
		if _, exists := newVisted[p.getKey()]; !exists || forRating {

			// Got to the end => add one
			if newPositions[i].height == 9 {
				if _, exists := allEndings[p.getKey()]; !exists {
					logger.Debug("Found a new accessable track on %s", newPositions[i].getKey())
					sum += 1
					allEndings[p.getKey()] = 1
				} else {
					logger.Debug("Found existing ending")
					allEndings[p.getKey()] += 1
				}
			} else {
				sum += d.getScoreOfTrack(newPositions[i], newVisted, allEndings, forRating)
			}
		}
	}

	return
}

func (d *Day) parseMap(input []string) {
	for _, line := range input {
		horizontal := []int{}

		for _, c := range line {
			if c == '.' {
				horizontal = append(horizontal, -1)
			} else {
				horizontal = append(horizontal, utils.ToInt(string(c)))
			}
		}

		d.topoMap = append(d.topoMap, horizontal)
	}
}

func (d *Day) Part2(in string) string {
	d.parseMap(utils.RemoveEmptyLines(strings.Split(in, "\n")))

	// Start processing of tracks with height 0
	sumScore := 0
	for y := range d.topoMap {
		for x := range d.topoMap[y] {
			ratings := make(map[string]int)

			co := cordinat{x: x, y: y, height: d.topoMap[y][x]}
			if co.height == 0 {
				d.getScoreOfTrack(
					co,
					make(map[string]int),
					ratings,
					true,
				)
			}

			// Count how
			for _, val := range ratings {
				sumScore += val
			}
		}
	}

	return fmt.Sprintf("%d", sumScore)
}
