package day_09

import (
	"fmt"
	"math"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

const freeSpaceId = -1

type Day struct{}

func (d *Day) Part1(in string) string {
	in = strings.TrimRight(in, "\n")
	diskLayout := d.transform(in)
	// logger.Debug("Transformed to: %d", diskLayout)

	// Fill up free space
	for i := len(diskLayout) - 1; i >= 0; i-- {
		// Skip free space
		if diskLayout[i] == -1 {
			continue
		}

		// Find first free space
		found := false
		for c := range diskLayout {
			if c >= i {
				break
			} else if diskLayout[c] == freeSpaceId {
				found = true
				diskLayout[c] = diskLayout[i]
				diskLayout[i] = freeSpaceId
			}
		}

		// Finished (no more free space to fill)
		if !found {
			break
		}
	}

	// Calculate the sum
	sum := 0
	for i, id := range diskLayout {
		if id == freeSpaceId {
			break
		}

		sum += i * id
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) transform(input string) (rtc []int) {
	for i, r := range input {
		rInt := utils.ToInt(string(r))

		idToPrint := math.Floor((float64(i)) / 2.0)
		if i%2 == 1 {
			// Free space
			idToPrint = freeSpaceId
		}

		for c := 0; c < rInt; c++ {
			rtc = append(rtc, int(idToPrint))
		}
	}

	return
}

func (d *Day) Part2(in string) string {
	in = strings.TrimRight(in, "\n")
	diskLayout := d.transform(in)
	// logger.Debug("Transformed to: %d", diskLayout)

	// Fill up free space
	for i := len(diskLayout) - 1; i >= 0; i-- {
		// Skip free space
		if diskLayout[i] == -1 {
			continue
		}

		// How big the block is
		size := 0
		for ii := i; ii >= 0; ii-- {
			if diskLayout[ii] == diskLayout[i] {
				size++
			} else {
				break
			}
		}

		// Find first free space of described blocks
		found := false
		for c := range diskLayout {
			if c >= i {
				break
			} else if diskLayout[c] == freeSpaceId {
				// Check if block fits into free space
				foundBlocks := 0

				for ii := c; ii < len(diskLayout) && ii < c+size; ii++ {
					if diskLayout[ii] != freeSpaceId {
						break
					}
					foundBlocks++
				}

				if foundBlocks == size {
					found = true
					// Set all to free space / moved id
					for ii := c; ii < len(diskLayout) && ii < c+size; ii++ {
						diskLayout[ii] = diskLayout[i]
						diskLayout[i] = freeSpaceId
						i--

						if i < 0 {
							break
						}
					}
					i++

					break
				}
			}
		}

		// Decrease int manually
		if !found {
			i -= size - 1
		}
	}

	// Calculate the sum
	// logger.Debug("%d", diskLayout)
	sum := 0
	for i, id := range diskLayout {
		if id != freeSpaceId {
			sum += i * id
		}
	}

	return fmt.Sprintf("%d", sum)
}
