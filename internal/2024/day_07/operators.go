package day_07

import (
	"fmt"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type operator interface {
	calculate(a, b int) int
}

type plus struct{}

func (p plus) calculate(a, b int) int {
	return a + b
}

type multiply struct{}

func (m multiply) calculate(a, b int) int {
	return a * b
}

type concat struct{}

func (c concat) calculate(a, b int) int {
	return utils.ToInt(fmt.Sprintf("%d%d", a, b))
}
