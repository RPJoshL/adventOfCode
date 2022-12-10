package day_09

import (
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	board [][]bool

	curHead Position
	curTail Position

	// Part two extras
	tails []Position
}

type Position struct {
	row    int
	column int
}

func (d *Day) Part1(in string) string {

	// Create initial board
	initialSize := 2000
	d.board = make([][]bool, initialSize)
	for i := range d.board {
		d.board[i] = make([]bool, initialSize)
	}

	// Initialize positions
	d.curHead.row = initialSize / 2
	d.curHead.column = initialSize / 2

	d.curTail.row = initialSize / 2
	d.curTail.column = initialSize / 2
	d.board[initialSize/2][initialSize/2] = true

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}
		pos := val[0:1]
		mov, _ := strconv.Atoi(val[2:])
		//fmt.Printf("Moving for %d in %s\n", mov, pos)
		for i := 0; i < mov; i++ {
			d.Move(pos)
		}
	}

	// Count positions
	sum := 0
	for _, val := range d.board {
		for _, wasOn := range val {
			if wasOn {
				sum++
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

// Left or right or Up and Down
func (d *Day) Move(direction string) {
	origHead := d.curHead
	if direction == "L" {
		d.curHead.column--
	} else if direction == "R" {
		d.curHead.column++
	} else if direction == "U" {
		d.curHead.row--
	} else if direction == "D" {
		d.curHead.row++
	}

	//fmt.Printf("Head: %s | Tail: %s\n", d.curHead, d.curTail)
	// Check if tail should move
	if d.isTailOnlyOneAway() || d.isDiagonally() || (d.curHead.row == d.curTail.row && d.curHead.column == d.curTail.column) {
		// The Head is already on the way :)
	} else {
		d.curTail = origHead
		// Update positions
		d.board[d.curTail.row][d.curTail.column] = true
	}
	//fmt.Printf("Head: %s | Tail: %s\n", d.curHead, d.curTail)

}

func (d *Day) isTailOnlyOneAway() bool {
	return (d.curHead.column == d.curTail.column && isOnlyOneAway(d.curHead.row, d.curTail.row)) ||
		(d.curHead.row == d.curTail.row && isOnlyOneAway(d.curHead.column, d.curTail.column))
}
func isOnlyOneAway(one int, two int) bool {
	//fmt.Printf("Comparing: %d %d %t %t\n", one, two, (one-1) == two, (one+1) == two)
	return (one-1) == two || (one+1) == two
}

func (d *Day) isDiagonally() bool {
	topRight := d.curHead.row-1 == d.curTail.row && d.curHead.column-1 == d.curTail.column
	topLeft := d.curHead.row-1 == d.curTail.row && d.curHead.column+1 == d.curTail.column
	bottomRight := d.curHead.row+1 == d.curTail.row && d.curHead.column-1 == d.curTail.column
	bottomLeft := d.curHead.row+1 == d.curTail.row && d.curHead.column+1 == d.curTail.column

	return topRight || topLeft || bottomLeft || bottomRight
}

func (d *Day) Part2(in string) string {
	// Create initial board
	initialSize := 2000
	d.board = make([][]bool, initialSize)
	for i := range d.board {
		d.board[i] = make([]bool, initialSize)
	}

	// Initialize positions
	d.curHead.row = initialSize / 2
	d.curHead.column = initialSize / 2

	d.tails = make([]Position, 10)
	for i := 0; i < 10; i++ {
		d.tails[i].row = initialSize / 2
		d.tails[i].column = initialSize / 2
	}
	d.board[initialSize/2][initialSize/2] = true

	for _, val := range strings.Split(in, "\n") {
		if val == "" {
			continue
		}
		pos := val[0:1]
		mov, _ := strconv.Atoi(val[2:])
		//fmt.Printf("Moving for %d in %s\n", mov, pos)
		for i := 0; i < mov; i++ {
			d.Move2(pos)
		}
	}

	// Count positions
	sum := 0
	for _, val := range d.board {
		for _, wasOn := range val {
			if wasOn {
				sum++
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

// Left or right or Up and Down
func (d *Day) Move2(direction string) {
	//origHead := d.curHead
	if direction == "L" {
		d.curHead.column--
	} else if direction == "R" {
		d.curHead.column++
	} else if direction == "U" {
		d.curHead.row++
	} else if direction == "D" {
		d.curHead.row--
	}

	for i := 0; i < 9; i++ {
		prevHead := d.curHead
		if i != 0 {
			prevHead = d.tails[i-1]
		}

		// Check if tail should moved
		if d.isTailOnlyOneAway2(d.tails[i], prevHead) || d.isDiagonally2(d.tails[i], prevHead) || (prevHead.row == d.tails[i].row && prevHead.column == d.tails[i].column) {
			// Not moving
		} else {
			// Up and Down
			if direction == "U" || direction == "D" {
				if d.tails[i].row < prevHead.row {
					d.tails[i].row++
				}
				if d.tails[i].row > prevHead.row {
					d.tails[i].row--
				}

				// left or right?
				if d.tails[i].column < prevHead.column {
					d.tails[i].column++
				} else if d.tails[i].column > prevHead.column {
					d.tails[i].column--
				}
				// Left and Right
			} else if direction == "L" || direction == "R" {
				if d.tails[i].column > prevHead.column {
					d.tails[i].column--
				}
				if d.tails[i].column < prevHead.column {
					d.tails[i].column++
				}

				// Up or down
				if d.tails[i].row < prevHead.row {
					d.tails[i].row++
				} else if d.tails[i].row > prevHead.row {
					d.tails[i].row--
				}
			}
		}

		if i == 8 {
			d.board[d.tails[i].row][d.tails[i].column] = true
		}

		//fmt.Printf("%d: Head: %s | Tail: %s\n", i+1, prevHead, d.tails[i])
	}

	//fmt.Printf("Head: %s | Tail: %s\n", d.curHead, d.curTail)
}
func (d *Day) isTailOnlyOneAway2(tail Position, prevTail Position) bool {
	return (prevTail.column == tail.column && isOnlyOneAway2(prevTail.row, tail.row)) ||
		(prevTail.row == tail.row && isOnlyOneAway2(prevTail.column, tail.column))
}
func isOnlyOneAway2(one int, two int) bool {
	//fmt.Printf("Comparing: %d %d %t %t\n", one, two, (one-1) == two, (one+1) == two)
	return (one-1) == two || (one+1) == two
}

func (d *Day) isDiagonally2(tail Position, prevTail Position) bool {
	topRight := prevTail.row-1 == tail.row && prevTail.column-1 == tail.column
	topLeft := prevTail.row-1 == tail.row && prevTail.column+1 == tail.column
	bottomRight := prevTail.row+1 == tail.row && prevTail.column-1 == tail.column
	bottomLeft := prevTail.row+1 == tail.row && prevTail.column+1 == tail.column

	return topRight || topLeft || bottomLeft || bottomRight
}
