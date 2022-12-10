package day_08

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	data [][]int
}

func (d *Day) Part1(in string) string {
	for _, row := range strings.Split(in, "\n") {
		columns := make([]int, len(row))
		if row == "" {
			continue
		}
		for o := 0; o < len(row); o++ {
			val, _ := strconv.Atoi(row[o : o+1])
			columns[o] = val
		}
		d.data = append(d.data, columns)
	}

	sum := 0
	for i := range d.data {
		for o, column := range d.data[i] {
			if d.isEdge(i, o) || d.checkRow(i, o, column) || d.checkColumn(o, i, column) {
				//if !d.isEdge(i, o) {
				//	fmt.Printf("For %d-%d | %t %t %t\n", i, o, d.isEdge(i, o), d.checkRow(i, o, column), d.checkColumn(o, i, column))
				//}
				sum += 1
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day) Part2(in string) string {
	for _, row := range strings.Split(in, "\n") {
		columns := make([]int, len(row))
		if row == "" {
			continue
		}
		for o := 0; o < len(row); o++ {
			val, _ := strconv.Atoi(row[o : o+1])
			columns[o] = val
		}
		d.data = append(d.data, columns)
	}

	max := 0
	// loop through every tree
	for i := range d.data {
		for o, column := range d.data[i] {

			sum := d.lookColumn(o, i, column) * d.lookRow(i, o, column)
			if sum > max {
				max = sum
			}
		}
	}

	return fmt.Sprintf("%d", max)
}

// Tree is standing on the edge
func (d *Day) isEdge(row int, col int) bool {
	return row == 0 || row == len(d.data)-1 ||
		col == 0 || col == len(d.data[row])-1
}

// Checks if for top or bottom is visible
func (d *Day) checkColumn(column int, treeRow int, val int) bool {

	// Top
	successfull := true
	for i := 0; i < treeRow; i++ {
		current := d.data[i][column]
		if current >= val {
			successfull = false
			break
		}
	}

	if successfull {
		return true
	}

	// Bottom
	successfull = true
	for i := len(d.data) - 1; i > treeRow; i-- {
		current := d.data[i][column]
		if current >= val {
			successfull = false
			break
		}
	}

	return successfull
}

// Checks if left or right is visible
func (d *Day) checkRow(row int, treeColumn int, val int) bool {
	// Left
	successfull := true
	for i := 0; i < treeColumn; i++ {
		current := d.data[row][i]
		if current >= val {
			successfull = false
			break
		}
	}

	if successfull {
		return true
	}

	// Right
	successfull = true
	for i := len(d.data[row]) - 1; i > treeColumn; i-- {
		current := d.data[row][i]
		if current >= val {
			successfull = false
			break
		}
	}

	return successfull
}

// Checks if for top or bottom is visible
func (d *Day) lookColumn(column int, treeRow int, val int) int {

	// Top
	topCount := 0
	for i := treeRow - 1; i >= 0; i-- {
		if i == -1 {
			break
		}

		current := d.data[i][column]

		topCount++
		if current >= val {
			break
		}
	}

	// Bottom
	bottomCount := 0
	for i := treeRow + 1; i < len(d.data); i++ {
		current := d.data[i][column]

		bottomCount++
		if current >= val {
			break
		}
	}

	//fmt.Printf("Bottom %d Top %d\n", bottomCount, topCount)
	return bottomCount * topCount
}

// Checks if left or right is visible
func (d *Day) lookRow(row int, treeColumn int, val int) int {
	// Left
	leftCount := 0
	for i := treeColumn - 1; i >= 0; i-- {
		current := d.data[row][i]

		leftCount++
		if current >= val {
			break
		}
	}

	// Right
	rightCount := 0
	for i := treeColumn + 1; i < len(d.data[row]); i++ {
		current := d.data[row][i]

		rightCount++
		if current >= val {
			break
		}
	}

	//fmt.Printf("Left %d Right %d\n", leftCount, rightCount)
	return leftCount * rightCount
}
