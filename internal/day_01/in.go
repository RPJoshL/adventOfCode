package day_01

import (
	"fmt"
	"strings"

	"rpjosh.de/adventOfCode2022/pkg/utils"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	current := 0
	max := 0

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			// Clear it
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		current += utils.ToInt(val)
	}

	return fmt.Sprintf("%d", max)
}

func (d *Day) Part2(in string) string {
	current := 0
	var totals []int

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			// Clear it
			totals = append(totals, current)
			current = 0
			continue
		}

		current += utils.ToInt(val)
	}
	totals = append(totals, current)

	// Sort array
	for i := 0; i < len(totals)-1; i++ {
		if totals[i+1] > totals[i] {
			tmp := totals[i]
			totals[i] = totals[i+1]
			totals[i+1] = tmp
			i = -1
		}
	}

	sum := 0
	for i := 0; i < 3; i++ {
		sum += totals[i]
	}

	return fmt.Sprintf("%d", sum)
}
