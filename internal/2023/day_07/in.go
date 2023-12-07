package day_07

import (
	"fmt"
	"sort"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
	hands map[byte]int
}

func (d *Day) Part1(in string) string {
	inSplit := strings.Split(in, "\n")
	//inSplit := []string{"23332 246"}

	hands := map[byte]int{
		'A': 0,
		'K': 1,
		'Q': 2,
		'J': 3,
		'T': 4,
		'9': 5,
		'8': 6,
		'7': 7,
		'6': 8,
		'5': 9,
		'4': 10,
		'3': 11,
		'2': 12,
	}

	sort.SliceStable(inSplit, func(a, b int) bool {
		aVal := getTypeOrders(strings.Split(inSplit[a], " ")[0])
		bVal := getTypeOrders(strings.Split(inSplit[b], " ")[0])

		if aVal > bVal {
			return true
		} else if aVal == bVal {
			// Get first number
			aS := strings.Split(inSplit[a], " ")[0]
			bS := strings.Split(inSplit[b], " ")[0]

			for i := 0; i < 5; i++ {
				if hands[aS[i]] > hands[bS[i]] {
					return true
				} else if hands[aS[i]] == hands[bS[i]] {
					continue
				} else {
					return false
				}
			}
			logger.Error("Two values are the same. What to do?")
			return false
		} else {
			return false
		}
	})

	rtc := 0
	for i, val := range inSplit {
		if val == "" {
			continue
		}
		logger.Debug("%s == %d", val, getTypeOrders(strings.Split(val, " ")[0]))
		rtc += (i + 1) * utils.ToInt(strings.Split(val, " ")[1])
	}

	return fmt.Sprintf("%d", rtc)
}

type Values struct {
	val       rune
	count     int
	positions string
}

func getTypeOrders(numbers string) int {

	if len(numbers) < 4 {
		logger.Error("Received invalid input: %q", numbers)
		return 0
	}

	// Five in a row
	if strings.Count(numbers, string(numbers[0])) == 5 {
		return 1
	}

	// Four in a row
	if strings.Count(numbers, string(numbers[0])) == 4 || strings.Count(numbers, string(numbers[1])) == 4 {
		return 2
	}
	// Prepare values
	numbersSorted := utils.SortRunes(numbers)
	runeMap := map[rune]int{}

	// Three of a kind or full house
	vals := make([]Values, 3)
	for i, val := range numbers {
		if vals[0].val == val {
			vals[0].count += 1
			vals[0].positions += fmt.Sprintf("%d", i)
		} else if vals[1].val == val {
			vals[1].count += 1
			vals[1].positions += fmt.Sprintf("%d", i)
		} else if vals[2].val == val {
			vals[2].count += 1
			vals[2].positions += fmt.Sprintf("%d", i)
		} else if vals[0].count == 0 {
			vals[0].val = val
			vals[0].count = 1
			vals[0].positions += fmt.Sprintf("%d", i)
		} else if vals[1].count == 0 {
			vals[1].val = val
			vals[1].count = 1
			vals[1].positions += fmt.Sprintf("%d", i)
		} else if vals[2].count == 0 {
			vals[2].val = val
			vals[2].count = 1
			vals[2].positions += fmt.Sprintf("%d", i)
		}
	}
	// We need three same kinds
	threeValid := false
	for i := 0; i < 3; i++ {
		if vals[i].count == 3 {
			threeValid = true
			/*
				//logger.Debug("==== Having three")
				// In order
				lastVal := int(vals[i].positions[0])
				for _, r := range vals[i].positions[1:] {
					//logger.Debug("order: %d %d", lastVal, r-1)
					if lastVal == int(r)-1 {
						threeValid = true
						lastVal = int(r)
					} else {
						threeValid = false
					}
				}
			*/
		}
	}

	if threeValid && ((vals[0].count == 3 && vals[1].count == 1) || (vals[1].count == 3 && vals[0].count == 1) || vals[2].count == 3 && vals[0].count == 1) {
		return 4
	} else if threeValid && (vals[1].count == 2 || vals[0].count == 2 || vals[3].count == 2) {
		return 3
	}

	// Two pair
	runeMap = map[rune]int{}
	for _, val := range numbersSorted {
		runeMap[val] += 1
	}
	pairs := 0
	for _, val := range runeMap {
		if val == 2 {
			pairs += 1
		}
	}

	// Two pairs
	if pairs == 2 {
		return 5
	}

	// One pair
	if pairs == 1 {
		return 6
	}

	// Do we need to distinct?
	areDistinct := true
	for _, val := range runeMap {
		if val > 1 {
			logger.Fatal("Values are not distinct. That's a logic error: %q", numbers)
			areDistinct = false
		}
	}

	if areDistinct {
		return 7
	}

	return 10
}

func (d *Day) GetBestJokerValue(numbers string) string {
	// Nothing to do here
	if !strings.Contains(numbers, "J") {
		return numbers
	}

	// Loop through every 'J'. A 'J' can occur twice -> Brute force with recursion :)
	bestNumbers := ""
	lastValue := 20
	for k := range d.hands {
		// Not allowed - endless loop!
		if k == 'J' {
			continue
		}

		lastJ := strings.LastIndex(numbers, "J")
		numbersNew := utils.ReplaceCharacterInString(numbers, string(k), lastJ)

		// Replace any other 'J' value
		numbersNew = d.GetBestJokerValue(numbersNew)

		// Test the value
		result := getTypeOrders(numbersNew)

		// Better result
		if result < lastValue {
			lastValue = result
			bestNumbers = numbersNew
		}
	}

	return bestNumbers
}

func (d *Day) Part2(in string) string {
	inSplit := strings.Split(in, "\n")
	//inSplit := []string{"23332 246"}

	d.hands = map[byte]int{
		'A': 0,
		'K': 1,
		'Q': 2,
		'T': 4,
		'9': 5,
		'8': 6,
		'7': 7,
		'6': 8,
		'5': 9,
		'4': 10,
		'3': 11,
		'2': 12,
		'J': 13,
	}

	sort.SliceStable(inSplit, func(a, b int) bool {
		aVal := getTypeOrders(d.GetBestJokerValue(strings.Split(inSplit[a], " ")[0]))
		bVal := getTypeOrders(d.GetBestJokerValue(strings.Split(inSplit[b], " ")[0]))

		if aVal > bVal {
			return true
		} else if aVal == bVal {
			// Get first number
			aS := strings.Split(inSplit[a], " ")[0]
			bS := strings.Split(inSplit[b], " ")[0]

			for i := 0; i < 5; i++ {
				if d.hands[aS[i]] > d.hands[bS[i]] {
					return true
				} else if d.hands[aS[i]] == d.hands[bS[i]] {
					continue
				} else {
					return false
				}
			}
			logger.Error("Two values are the same. What to do?")
			return false
		} else {
			return false
		}
	})

	rtc := 0
	for i, val := range inSplit {
		if val == "" {
			continue
		}
		logger.Debug("%s == %d", val, getTypeOrders(strings.Split(val, " ")[0]))
		rtc += (i + 1) * utils.ToInt(strings.Split(val, " ")[1])
	}

	return fmt.Sprintf("%d", rtc)
}
