package day_07

import (
	"fmt"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	operators []operator

	result int
	values []int
}

func (d *Day) Part1(in string) string {
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.operators = []operator{plus{}, multiply{}}

	sum := 0
	for _, line := range lines {
		d.result, d.values = d.parseLine(line)

		if d.checkResult(0, []operator{}) {
			sum += d.result
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) parseLine(line string) (result int, values []int) {
	for i, val := range strings.Split(line, " ") {
		if i == 0 {
			result = utils.ToInt(val[0 : len(val)-1])
		} else {
			values = append(values, utils.ToInt(val))
		}
	}

	return
}

func (d *Day) checkResult(startIndex int, allOperators []operator) bool {
	// No valid start index provided -> return
	if startIndex >= len(d.values) {
		return false
	}

	// Brute force calculation with every operator
	isValid := false
	for _, op := range d.operators {
		if isValid {
			return true
		}

		newOperators := make([]operator, len(allOperators)+1)
		copy(newOperators, allOperators)
		newOperators[startIndex] = op

		if d.calculateResult(newOperators, plus{}) == d.result {
			return true
		} else {
			isValid = d.checkResult(startIndex+1, newOperators)
		}
	}

	return isValid
}

func (d *Day) calculateResult(allOperators []operator, operatorForMissing operator) int {
	lastValue := d.values[0]

	for i := 1; i < len(d.values); i++ {
		oI := i - 1
		if i >= len(allOperators) {
			lastValue = operatorForMissing.calculate(lastValue, d.values[i])
		} else {
			lastValue = allOperators[oI].calculate(lastValue, d.values[i])
		}
	}

	//logger.Debug("Calculated result = %d for %s", lastValue, allOperators)
	return lastValue
}

func (d *Day) Part2(in string) string {
	lines := utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.operators = []operator{plus{}, multiply{}, concat{}}

	sum := 0
	for _, line := range lines {
		d.result, d.values = d.parseLine(line)

		if d.checkResult(0, []operator{}) {
			sum += d.result
		}
	}

	return fmt.Sprintf("%d", sum)
}
