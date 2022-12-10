package day_03

import (
	"fmt"
	"strings"
)

type Day struct{}

func (d *Day) Part1(in string) string {
	sum := 0

	for _, val := range strings.Split(in, "\n") {
		center := len(val) / 2
		fSegment := []rune(val[0:center])
		sSegment := []rune(val[center:])
		//fmt.Printf("%s | %s\n", fSegment, sSegment)

	outer:
		for _, fChar := range fSegment {
			for o, sChar := range sSegment {
				// Check if it is contained
				if fChar == sChar {
					sum += d.getPriority(fChar)
					sSegment[o] = 0
					break outer
				}
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) Part2(in string) string {
	sum := 0

	inArray := strings.Split(in, "\n")

outer:
	for i := 3; i <= len(inArray); i += 3 {

		// Loop for every character to find one every contains
		for chars := 65; chars <= 122; chars++ {
			if chars >= 91 && chars <= 96 {
				continue
			}

			numberFound := 0

			// Find number in three pair
			for ii := i - 3; ii < i; ii++ {
				found := false

				// Loop a single one
				for _, char := range []rune(inArray[ii]) {
					if char == rune(chars) {
						found = true
						break
					}
				}

				if !found {
					break
				} else {
					numberFound++
				}
			}

			if numberFound == 3 {
				sum += d.getPriority(rune(chars))
				continue outer
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) getPriority(in rune) int {
	//in := []byte(inString)[0]
	fmt.Printf("Got: %d %s\n", in, string(in))
	if in >= 97 {
		return int(in - 96)
	} else {
		return int(in - 38)
	}
}
