package day_11

import (
	"fmt"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct{}

type MonkeyData struct {
	operation      MonkeyOperation
	test           MonkeyTest
	items          []int
	inspectedItems int
}

type MonkeyOperation struct {
	multiply  bool
	sameValue bool
	value     int
}

type MonkeyTest struct {
	divValue     int
	trueThrowTo  int
	falseThrowTo int
}

func (d *Day) Part1(in string) string {
	return d.Perform(in, false, 20)
}

func (d *Day) Part2(in string) string {
	return d.Perform(in, true, 10000)
}

func (d *Day) Perform(in string, part2 bool, iterations int) string {
	var monkeys []MonkeyData

	part2DivValues := 1
	for i, val := range strings.Split(in, "\n") {
		mongo := i / 7

		// Parse monkey data
		switch i % 7 {
		case 0:
			{
				monkeys = append(monkeys, MonkeyData{})
			}
		case 1:
			{
				// Starting items
				items := strings.Split(val[18:], ", ")
				//monkeys[mongo].items = make([]int, 10)
				for _, item := range items {
					monkeys[mongo].items = append(monkeys[mongo].items, utils.ToInt(item))
				}
			}
		case 2:
			{
				// Operation >19<old
				op := val[23:]
				monkeys[mongo].operation.multiply = op[0:1] == "*"
				monkeys[mongo].operation.sameValue = op[2:] == "old"

				if !monkeys[mongo].operation.sameValue {
					monkeys[mongo].operation.value = utils.ToInt(op[2:])
				}
			}
		case 3:
			{
				// Test
				monkeys[mongo].test.divValue = utils.ToInt(val[21:])

				// That the values are getting not to big,
				part2DivValues *= utils.ToInt(val[21:])
			}
		case 4:
			{
				// test true
				monkeys[mongo].test.trueThrowTo = utils.ToInt(val[29:])
			}
		case 5:
			{
				//test false
				monkeys[mongo].test.falseThrowTo = utils.ToInt(val[30:])
			}
		}
	}

	for i := 0; i < iterations; i++ {
		for mongo := range monkeys {

			monk := &monkeys[mongo]
			// Inspect items
			for item := range monk.items {

				// Execute operation
				operationValue := monk.operation.value
				if monk.operation.sameValue {
					operationValue = monk.items[item]
				}
				if monk.operation.multiply {
					monk.items[item] = monk.items[item] * operationValue
				} else {
					monk.items[item] = monk.items[item] + operationValue
				}

				// Monkey gets bored
				if !part2 {
					monk.items[item] /= 3
				} else {
					monk.items[item] = monk.items[item] % part2DivValues
					//fmt.Printf("Number: %d\n", monk.items[item])
				}

				// Is it diviable?
				var throwTo int
				if monk.items[item]%monk.test.divValue == 0 {
					throwTo = monk.test.trueThrowTo
				} else {
					throwTo = monk.test.falseThrowTo
				}
				monkeys[throwTo].items = append(monkeys[throwTo].items, monk.items[item])

				monk.inspectedItems++
			}

			// clear items
			monk.items = nil
		}

		//fmt.Printf("\nRound %d\n", i)
		//for i, monkey := range monkeys {
		//	fmt.Printf("Monkey %d: %s\n", i, monkey.items)
		//}
	}

	maxOne := 0
	maxTwo := 0

	for _, monkey := range monkeys {
		if monkey.inspectedItems > maxOne {
			maxTwo = maxOne
			maxOne = monkey.inspectedItems
		} else if monkey.inspectedItems > maxTwo {
			maxTwo = monkey.inspectedItems
		}
	}

	return fmt.Sprintf("%d", maxOne*maxTwo)
}
