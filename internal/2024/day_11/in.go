package day_11

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	stones []int

	// cache indexed by number and tries for count
	cacheMap map[int]map[int]int
}

func (d *Day) Part1(in string) string {
	d.parse(in)

	// Get all transformer
	transformers := []rule{zeroReplace{}, evenNumber{}, otherNumber{}}

	for i := 0; i < 25; i++ {
		for s := 0; s < len(d.stones); s++ {
			for _, trans := range transformers {
				transformed, replace, add := trans.transform(d.stones[s])

				if transformed {
					d.stones[s] = replace

					if len(add) > 0 {
						newSlice := make([]int, 0)
						newSlice = append(newSlice, d.stones[0:s+1]...)
						newSlice = append(newSlice, add...)
						newSlice = append(newSlice, d.stones[s+1:]...)
						d.stones = newSlice
						s += len(add)
					}

					break
				}
			}
		}

		logger.Debug("Transformed %d blinks", i+1)
	}

	// logger.Debug("Stones: %d", d.stones)
	return fmt.Sprintf("%d", len(d.stones))
}

func (d *Day) parse(in string) {
	rtc := utils.ParseSeperatedNumbers(
		utils.RemoveEmptyLines(strings.Split(in, "\n")),
		" ",
	)
	d.stones = rtc[0]
}

func (d *Day) Part2(in string) string {
	d.parse(in)

	d.cacheMap = map[int]map[int]int{}
	stones := len(d.stones)
	for _, stone := range d.stones {
		stones += d.stonesForNumber(stone, 1, func(index int, withVal int) {})
	}

	// logger.Debug("Stones: %d", d.stones)
	return fmt.Sprintf("%d", stones)
}

func (d *Day) stonesForNumber(value int, index int, increment func(index int, withVal int)) int {
	max := 75

	// Check if we do have a cache for this number
	if c, exists := d.cacheMap[value]; exists {
		// Check if we have data for full points
		if _, ex := c[max-index]; ex {
			sum := 0

			for i := max - index; i >= 0; i-- {
				increment(i+index, c[i])
				sum += c[i]
			}

			// logger.Debug("Used cache for %dx steps", max-index)
			// logger.Debug("%d", c)
			return sum
		}
	}

	// Wrapper function to write cache
	thisCacheValues := make(map[int]int, 0)
	incrementWrap := func(wrapIndex int, withVal int) {
		increment(wrapIndex, withVal)

		if _, ex := thisCacheValues[wrapIndex-index]; ex {
			thisCacheValues[wrapIndex-index] += withVal
		} else {
			thisCacheValues[wrapIndex-index] = withVal
		}
	}

	transformers := []rule{zeroReplace{}, evenNumber{}, otherNumber{}}

	for _, trans := range transformers {
		transformed, replace, add := trans.transform(value)

		if transformed {
			rtc := len(add)

			if index >= max {
				incrementWrap(index, rtc)
				return rtc
			} else {
				// logger.Debug("Processing %d %d", replace, add)
				rtc += d.stonesForNumber(replace, index+1, incrementWrap)

				// Make sure that map has this key
				incrementWrap(index, len(add))

				for _, valAdd := range add {
					rtc += d.stonesForNumber(valAdd, index+1, incrementWrap)
				}

				d.cacheMap[value] = thisCacheValues
				logger.Debug("%d should match: %d", rtc, thisCacheValues)
				return rtc
			}
		}
	}

	return 0
}
