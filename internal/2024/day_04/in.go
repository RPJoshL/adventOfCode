package day_04

import (
	"fmt"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	input []string

	searchWord string
}

func (d *Day) Part1(in string) string {
	// Parse input
	d.input = utils.RemoveEmptyLines(strings.Split(in, "\n"))
	d.searchWord = "XMAS"

	rtc := 0
	rtc += d.horizontal()
	rtc += d.vertical()
	rtc += d.diagonal(false) // Left - Right
	rtc += d.diagonal(true)  // Right - Left

	return fmt.Sprintf("%d", rtc)
}

func (d *Day) horizontal() int {
	rtc := 0

	for _, line := range d.input {
		rtc += strings.Count(line, d.searchWord)
		rtc += strings.Count(utils.ReverseString(line), d.searchWord)
	}

	return rtc
}

func (d *Day) vertical() int {
	rtc := 0

	for i := 0; i < len(d.input[0]); i++ {
		column := ""

		for c := 0; c < len(d.input); c++ {
			column += string(d.input[c][i])
		}

		rtc += strings.Count(column, d.searchWord)
		rtc += strings.Count(utils.ReverseString(column), d.searchWord)
	}

	return rtc
}

func (d *Day) diagonal(reverseLine bool) int {
	rtc := 0

	for i := 0; i < len(d.input); i++ {
		diag := ""

		for cc := 0; cc < len(d.input[i]); cc++ {
			row := i
			diag = ""

			for c := cc; c < len(d.input[i]); c++ {
				if row >= len(d.input) || row >= i+4 {
					break
				}

				if reverseLine {
					diag += string(utils.ReverseString(d.input[row])[c])
				} else {
					diag += string(d.input[row][c])
				}

				row++
			}

			rtc += strings.Count(diag, d.searchWord)
			rtc += strings.Count(utils.ReverseString(diag), d.searchWord)
		}

	}

	return rtc
}

func (d *Day) Part2(in string) string {
	d.input = utils.RemoveEmptyLines(strings.Split(in, "\n"))

	rtc := 0
	for i, line := range d.input {
		for c := 0; c < len(d.input[i]); c++ {
			if line[c] == 'A' && i > 0 && c > 0 && i < len(d.input)-1 && c < len(d.input[i])-1 {
				// Characters at left top, left bottom, ...
				lt := d.input[i-1][c-1]
				lb := d.input[i+1][c-1]
				rt := d.input[i-1][c+1]
				rb := d.input[i+1][c+1]

				// If diagonal
				d1 := (lt == 'M' && rb == 'S') || (lt == 'S' && rb == 'M')
				d2 := (lb == 'M' && rt == 'S') || (lb == 'S' && rt == 'M')

				if d1 && d2 {
					rtc++
				}

				// logger.Debug("%dx%d: %s A %s | %s A %s = %t && %t", i, c, string(lt), string(rb), string(rt), string(lb), d1, d2)
			}
		}
	}

	return fmt.Sprintf("%d", rtc)
}
