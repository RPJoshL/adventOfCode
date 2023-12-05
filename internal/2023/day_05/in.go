package day_05

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {

	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\d+`)
	inSplit := strings.Split(in, "\n")

	// Contains the previous values of the last "section"
	lastValues := utils.ConvertArrayToInt(numberRegex.FindAllString(inSplit[0], -1))

	// Copy last values for each mapping.
	// It could be possible that an element is persent
	// in multiple mapping values
	lastValuesCurrMapp := make([]int, len(lastValues))
	copy(lastValuesCurrMapp, lastValues)

	for _, val := range inSplit[1:] {
		// Get numbers of input
		numbersString := numberRegex.FindAllString(val, -1)
		numbers := utils.ConvertArrayToInt(numbersString)

		// Re-Initialize values
		if val == "" {
			copy(lastValues, lastValuesCurrMapp)
		}

		// Don't do anything with invalid lines
		if len(numbers) <= 2 {
			continue
		}

		// Try to map the values
		startDest := numbers[0]
		startSource := numbers[1]
		length := numbers[2]
		logger.Debug("Found values %d %d %d", startDest, startSource, length)

		// Try to find the value in lastValues
		for v, val := range lastValues {
			if val >= startSource && val < startSource+length {
				diff := val - startSource
				lastValuesCurrMapp[v] = startDest + diff
			}
		}
	}

	return fmt.Sprintf("%d", utils.GetMinValue(lastValues))
}

func (d *Day) Part2(in string) string {
	// Regex to get all numbers
	numberRegex := regexp.MustCompile(`\d+`)
	inSplit := strings.Split(in, "\n")
	offset := 0

	// Multi thread
	minValue := math.MaxInt32
	var wg sync.WaitGroup
	var mtx sync.Mutex

	// Parse mapping once
	mapping := make([][]int, 0)
	for _, val := range inSplit[1:] {
		// Get numbers of input
		numbersString := numberRegex.FindAllString(val, -1)
		numbers := utils.ConvertArrayToInt(numbersString)

		if len(numbers) <= 2 && val != "" {
			continue
		}

		mapping = append(mapping, numbers)
	}

	// Unfold paris of seed values
	seeds := utils.ConvertArrayToInt(numberRegex.FindAllString(inSplit[offset], -1))
	for i := 0; i+1 < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			lastMin := seeds[i]
			for m := seeds[i]; m < seeds[i]+seeds[i+1]; m++ {
				if m%5000000 == 0 {
					logger.Debug("%d. Parsing %d", i, m)
				}

				// Bulk
				if m%100000 == 0 || m == seeds[i]+seeds[i+1]-1 {
					wg.Add(1)
					go func(min, max int) {
						rtc := d.ParseSingleFor2(min, max, &mapping)
						mtx.Lock()
						if rtc < minValue {
							minValue = rtc
						}
						mtx.Unlock()
						wg.Done()
					}(lastMin, m)
					lastMin = m
				}

			}
			logger.Info("Finished %d", i)
			wg.Done()
		}(i)

		// Limit max number of goroutines
		if i != 0 && i%10 == 0 {
			wg.Wait()
		}
	}
	wg.Wait()

	return fmt.Sprintf("%d", minValue)
}

func (d *Day) ParseSingleFor2(min int, max int, values *[][]int) int {

	lastValues := make([]int, max-min)
	for i := min; i < max; i++ {
		lastValues[i-min] = i
	}

	// Copy last values for each mapping.
	// It could be possible that an element is persent
	// in multiple mapping values
	lastValuesCurrMapp := make([]int, len(lastValues))
	copy(lastValuesCurrMapp, lastValues)

	for _, val := range *values {
		// Re-Initialize values
		if len(val) == 0 {
			copy(lastValues, lastValuesCurrMapp)
			continue
		}

		// Try to map the values
		startDest := val[0]
		startSource := val[1]
		length := val[2]

		// Try to find the value in lastValues
		for v, val := range lastValues {
			if val >= startSource && val < startSource+length {
				diff := val - startSource
				lastValuesCurrMapp[v] = startDest + diff
			}
		}
	}

	return utils.GetMinValue(lastValues)
}
