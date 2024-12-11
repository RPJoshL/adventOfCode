package day_11

import (
	"fmt"
	"strings"

	"rpjosh.de/adventOfCode/pkg/utils"
)

type rule interface {
	transform(value int) (transformed bool, replaceWith int, add []int)
}

type zeroReplace struct{}

func (z zeroReplace) transform(value int) (transformed bool, replaceWith int, add []int) {
	if value != 0 {
		return
	}

	replaceWith = 1
	transformed = true

	return
}

type evenNumber struct{}

func (e evenNumber) transform(value int) (transformed bool, replaceWith int, add []int) {
	valStr := fmt.Sprintf("%d", value)

	if len(valStr)%2 != 0 {
		return
	}

	// Replace with
	half := len(valStr) / 2
	replaceWith = utils.ToInt(valStr[0:half])

	// Add new values
	addStr := valStr[half:]
	addStr = strings.TrimLeft(addStr, "0")
	if addStr == "" {
		addStr = "0"
	}
	add = []int{utils.ToInt(addStr)}

	transformed = true
	return
}

type otherNumber struct{}

func (e otherNumber) transform(value int) (transformed bool, replaceWith int, add []int) {
	transformed = true
	replaceWith = value * 2024
	return
}
