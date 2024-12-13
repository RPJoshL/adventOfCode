package day_13

import (
	"fmt"
	"strings"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/pkg/utils"
)

type Day struct {
}

type cordinate struct {
	x int
	y int
}

func (d *Day) Part1(in string) string {
	blocks := strings.Split(in, "\n\n")

	sum := 0
	for _, block := range blocks {
		lines := utils.RemoveEmptyLines(strings.Split(block, "\n"))
		if len(lines) <= 2 {
			logger.Debug("Skipping line: %s", block)
			continue
		}

		btnA := cordinate{}
		btnB := cordinate{}
		price := cordinate{}

		utils.PanicTwo(fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &btnA.x, &btnA.y))
		utils.PanicTwo(fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &btnB.x, &btnB.y))
		utils.PanicTwo(fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &price.x, &price.y))

		sum += d.calculateCombination(btnA, btnB, price, 100)
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) calculateCombination(btnA, btnB, price cordinate, max int) (tokens int) {

	// There is either EXACTLY one or no solution. Using cramers rule here
	a := (btnB.y*price.x - btnB.x*price.y) / (btnA.x*btnB.y - btnA.y*btnB.x)
	b := (btnA.y*price.x - btnA.x*price.y) / (btnA.y*btnB.x - btnA.x*btnB.y)

	isValid := a*btnA.x+b*btnB.x == price.x && a*btnA.y+b*btnB.y == price.y
	if ((a <= max && b <= max) || max == -1) && isValid {
		return a*3 + b
	} else {
		return 0
	}
}

func (d *Day) Part2(in string) string {
	blocks := strings.Split(in, "\n\n")

	sum := 0
	for _, block := range blocks {
		lines := utils.RemoveEmptyLines(strings.Split(block, "\n"))
		if len(lines) <= 2 {
			logger.Debug("Skipping line: %s", block)
			continue
		}

		btnA := cordinate{}
		btnB := cordinate{}
		price := cordinate{}

		utils.PanicTwo(fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &btnA.x, &btnA.y))
		utils.PanicTwo(fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &btnB.x, &btnB.y))
		utils.PanicTwo(fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &price.x, &price.y))

		priceAdd := 10000000000000
		price.x += priceAdd
		price.y += priceAdd

		sum += d.calculateCombination(btnA, btnB, price, -1)
	}

	return fmt.Sprintf("%d", sum)
}
