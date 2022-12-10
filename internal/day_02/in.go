package day_02

import (
	"fmt"
	"strings"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	sum := 0
	for _, val := range strings.Split(in, "\n") {
		playersVal := strings.Split(val, " ")

		if len(playersVal) <= 1 {
			continue
		}

		me := d.getValue(playersVal[1])
		you := d.getValue(playersVal[0])

		// 1 = stein
		// 2 = papier
		// 3 = schere

		// papier - stein
		if me == 1 && you == 3 {
			me += 6
		} else if me == 2 && you == 1 {
			me += 6
		} else if me == 3 && you == 2 {
			me += 6
		} else if me == you {
			me += 3
		}

		sum += me
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) Part2(in string) string {
	sum := 0
	for _, val := range strings.Split(in, "\n") {
		playersVal := strings.Split(val, " ")

		if len(playersVal) <= 1 {
			continue
		}

		me := d.getValue(playersVal[1])
		you := d.getValue(playersVal[0])

		// 1 = stein   - verlieren
		// 2 = papier  - unentschieden
		// 3 = schere  - gewinnen

		if me == 3 {
			if you == 1 {
				me = 2
			} else if you == 2 {
				me = 3
			} else {
				me = 1
			}
		} else if me == 2 {
			me = you
		} else {
			if you == 1 {
				me = 3
			} else if you == 2 {
				me = 1
			} else {
				me = 2
			}
		}

		// papier - stein
		if me == 1 && you == 3 {
			me += 6
		} else if me == 2 && you == 1 {
			me += 6
		} else if me == 3 && you == 2 {
			me += 6
		} else if me == you {
			me += 3
		}

		sum += me
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) getValue(val string) int {
	if val == "X" || val == "A" {
		return 1
	} else if val == "Y" || val == "B" {
		return 2
	} else {
		return 3
	}
}
