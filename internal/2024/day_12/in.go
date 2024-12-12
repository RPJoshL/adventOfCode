package day_12

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	plants          []string
	processedPlants map[string]int
	regions         map[string]*region
}

type cordinat struct {
	x int
	y int
}

func (c cordinat) getKey() string {
	return fmt.Sprintf("%dx%d", c.x, c.y)
}

const (
	LEFT   = "LEFT"
	RIGHT  = "RIGHT"
	TOP    = "TOP"
	BOTTOM = "BOTTOM"
)

type direction struct {
	xOffset int
	yOffset int
	minX    int
	maxX    int
	minY    int
	maxY    int

	direction string
}

type region struct {
	fances         int
	plants         int
	fancesPosition map[string]bool
}

func (d *Day) Part1(in string) string {
	d.plants = utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.processedPlants = map[string]int{}
	d.regions = map[string]*region{}

	rtc := 0
	for y, row := range d.plants {
		for x := range row {
			cord := cordinat{x: x, y: y}

			// Init region
			regionKey := cord.getKey()
			region := &region{fancesPosition: map[string]bool{}}
			d.regions[regionKey] = region

			// Process and get result
			d.processSidePlant(cord, regionKey)
			rtc += region.fances * region.plants
			logger.Debug("For %s: fances = %d | plants = %d", regionKey, region.fances, region.plants)
		}
	}

	return fmt.Sprintf("%d", rtc)
}

func (d *Day) equals(source, dest cordinat) bool {
	s := d.plants[source.y][source.x]
	dd := d.plants[dest.y][dest.x]

	return s == dd
}

func (dd *Day) processSidePlant(c cordinat, regionKey string) {

	// Skip if already processed
	if _, exists := dd.processedPlants[c.getKey()]; exists {
		return
	} else {
		dd.processedPlants[c.getKey()] = 1
		dd.regions[regionKey].plants++
	}

	// All possible directions
	directions := []direction{
		{xOffset: -1, direction: LEFT},
		{xOffset: 1, maxX: len(dd.plants[0]), direction: RIGHT},
		{yOffset: -1, direction: TOP},
		{yOffset: 1, maxY: len(dd.plants), direction: BOTTOM},
	}

	// Check every direction
	for _, d := range directions {
		other := cordinat{x: c.x + d.xOffset, y: c.y + d.yOffset}

		// Out of bounds -> we need only a fence here
		if other.x < d.minX || other.y < d.minY || (other.x >= d.maxX && d.maxX != 0) || (other.y >= d.maxY && d.maxY != 0) {
			dd.regions[regionKey].fances++
			dd.regions[regionKey].fancesPosition[c.getKey()+d.direction] = true
		} else {
			// Process the next one for the same plant
			if dd.equals(c, other) {
				dd.processSidePlant(other, regionKey)
			} else {
				// We need a fence here
				dd.regions[regionKey].fances++
				dd.regions[regionKey].fancesPosition[c.getKey()+d.direction] = true
			}
		}
	}
}

func (d *Day) Part2(in string) string {
	d.plants = utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.processedPlants = map[string]int{}
	d.regions = map[string]*region{}

	rtc := 0
	for y, row := range d.plants {
		for x := range row {
			cord := cordinat{x: x, y: y}

			// Init region
			regionKey := cord.getKey()
			region := &region{fancesPosition: make(map[string]bool)}
			d.regions[regionKey] = region

			// Process and get result
			d.processSidePlant(cord, regionKey)
			// logger.Debug("For %s: fances = %d | plants = %d", regionKey, region.fancesPosition, region.plants)

			// Loop through every other one rough
			for key := range region.fancesPosition {
				if !region.fancesPosition[key] {
					continue
				}

				var x, y int
				var dir string
				if _, err := fmt.Sscanf(key, "%dx%d%s", &x, &y, &dir); err != nil {
					logger.Fatal("Failed to scan key %q: %s", key, err)
				}

				otherDirection := ""
				otherDirectionOffset := 0
				if strings.HasSuffix(key, TOP) {
					otherDirection = BOTTOM
					otherDirectionOffset = -1
				}
				if strings.HasSuffix(key, BOTTOM) {
					otherDirection = TOP
					otherDirectionOffset = 1
				}
				if strings.HasSuffix(key, LEFT) {
					otherDirection = RIGHT
					otherDirectionOffset = -1
				}
				if strings.HasSuffix(key, RIGHT) {
					otherDirection = LEFT
					otherDirectionOffset = 1
				}

				// Remove all fences (horizontal) that are in one line
				for xx := x + 1; xx < len(d.plants[0]); xx++ {
					keyThis := fmt.Sprintf("%dx%d%s", xx, y, dir)
					keyOther := fmt.Sprintf("%dx%d%s", xx+otherDirectionOffset, y, otherDirection)
					found := false

					if val, exists := region.fancesPosition[keyThis]; exists && val {
						region.fancesPosition[keyThis] = false
						found = true
					}
					if val, exists := region.fancesPosition[keyOther]; exists && val && 1 == 0 {
						region.fancesPosition[keyOther] = false
						found = true
					}

					if !found {
						break
					}
				}

				// Remove all fences (vertical) that are in one line
				for yy := y + 1; yy < len(d.plants); yy++ {
					keyThis := fmt.Sprintf("%dx%d%s", x, yy, dir)
					keyOther := fmt.Sprintf("%dx%d%s", x, yy+otherDirectionOffset, otherDirection)
					found := false

					if val, exists := region.fancesPosition[keyThis]; exists && val {
						region.fancesPosition[keyThis] = false
						found = true
					}
					if val, exists := region.fancesPosition[keyOther]; exists && val && 1 == 0 {
						region.fancesPosition[keyOther] = false
						found = true
					}

					if !found {
						break
					}
				}
			}

			// Get all distinct fences
			count := 0
			for _, val := range region.fancesPosition {
				if val {
					count++
				}
			}
			rtc += count * region.plants
			logger.Debug("%s has %d distinct fences", regionKey, count)
		}
	}

	return fmt.Sprintf("%d", rtc)
}
