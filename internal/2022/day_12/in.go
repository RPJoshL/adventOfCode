package day_12

import (
	"fmt"
	"strings"
	"time"
)

type Day struct {
	heightMap [][]rune
	visited   [][]bool

	currentPosition     Position
	destinationPosition Position
}

type Position struct {
	row    int
	column int
}

const MaxInt = int(^uint(0) >> 1)

func (d *Day) Part1(in string) string {

	// Parse map
	for i, val := range strings.Split(in, "\n") {
		d.heightMap = append(d.heightMap, make([]rune, len(val)))
		d.visited = append(d.visited, make([]bool, len(val)+2))

		for c := 0; c < len(val); c++ {
			d.heightMap[i][c] = rune(val[c])
			d.visited[i][c] = false

			strVal := val[c : c+1]
			if strVal == "S" {
				d.currentPosition = Position{i, c}

				min := "a"
				d.heightMap[i][c] = rune(min[0])

				d.visited[i][c] = true
			} else if strVal == "E" {
				d.destinationPosition = Position{i, c}

				min := "z"
				d.heightMap[i][c] = rune(min[0])
			}
		}
	}
	d.visited = append(d.visited, make([]bool, 220))

	//fmt.Printf("\n%s\n", d.heightMap[0])
	i, _ := d.makeOneStep(d.visited, d.currentPosition, 0)
	return fmt.Sprintf("%d", i)
}

// Brute force the shortest way :)
// Positions already visited are ignored
func (d *Day) makeOneStep(visited [][]bool, pos Position, stepsMade int) (int, [][]bool) {
	currentVal := d.heightMap[pos.row][pos.column]
	visited[pos.row][pos.column] = true

	// We are arrived
	if pos.column == d.destinationPosition.column && pos.row == d.destinationPosition.row {
		fmt.Printf("We are here: %d\n", stepsMade)
		time.Sleep(50 * time.Millisecond)
		return stepsMade, visited
	}
	minSteps := MaxInt
	var minVisited [][]bool = clone(&visited)
	if stepsMade == 0 {
		//fmt.Printf("\nPossible to right: %t", isStepPossible(d.heightMap[pos.row][pos.column], d.heightMap[pos.row][pos.column+1]))
		//fmt.Printf("%d %d\n", d.heightMap[pos.row][pos.column], d.heightMap[pos.row][pos.column+1])
		//fmt.Printf("Posistion: %s\n", pos)
		//fmt.Printf("%s", visited2[1])
	}

	//visitedReal := clone(&visited)
	//if pos.row >= 18 && pos.row <= 24 && pos.column <= 184 && pos.column >= 12 {
	fmt.Printf("Position: %s (%d)\n", pos, stepsMade)
	//fmt.Printf("End:      %s\n", d.destinationPosition)
	//fmt.Printf("Visited: %s", visited[pos.row])
	//time.Sleep(50 * time.Millisecond)
	//}
	//fmt.Printf("Visided: %s\n", visited)

	// Make a step to each direction an return the last one
	// Because we could get trapped
	if d.isNotVisitedAndPossible(pos.column+1, pos.row, visited, currentVal) {
		//fmt.Printf("Going to the right :) \n")
		steps, vis := d.makeOneStep(visited, Position{column: pos.column + 1, row: pos.row}, stepsMade+1)

		if steps <= minSteps {
			minVisited = vis
			minSteps = steps
			//visited = clone(&visitedReal)
			//visited[pos.row][pos.column+1] = true

			//minVisited = vis
			//minVisited = realVisited
		} // else {
		visited[pos.row][pos.column+1] = true
		minVisited[pos.row][pos.column+1] = true
		//visited = clone(&visitedReal)
		//visited = vis
		//visited[pos.row][pos.column] = false

		//visited[pos.row][pos.column+1] = true
		//visited = realVisited
		//}
	}
	if d.isNotVisitedAndPossible(pos.column-1, pos.row, visited, currentVal) {
		//fmt.Printf("Going to the left :) \n")
		steps, vis := d.makeOneStep(visited, Position{column: pos.column - 1, row: pos.row}, stepsMade+1)

		if steps <= minSteps {
			minVisited = vis
			minSteps = steps
			// Remove visited flag of previous
			//visited[pos.row][pos.column-1] = true

			//visited[pos.row][pos.column+1] = false

			//visited = clone(&visitedReal)

			//visited = vis
			//minVisited = vis
			//minVisited = realVisited
		} // else {
		visited[pos.row][pos.column-1] = true
		minVisited[pos.row][pos.column-1] = true
		//visited = clone(&visitedReal)
		//visited = vis
		//visited[pos.row][pos.column] = false
		//visited[pos.row][pos.column-1] = false
		//visited = realVisited
		//}
	}
	if d.isNotVisitedAndPossible(pos.column, pos.row+1, visited, currentVal) {
		//fmt.Printf("Going to the top :) \n")
		steps, vis := d.makeOneStep(visited, Position{column: pos.column, row: pos.row + 1}, stepsMade+1)

		if steps <= minSteps {
			minVisited = vis
			minSteps = steps

			//visited = clone(&visitedReal)
			//visited[pos.row+1][pos.column] = true

			//visited[pos.row][pos.column+1] = false
			//if pos.column != 0 {
			//	visited[pos.row+1][pos.column] = false
			//}

			//minVisited = vis
			//minVisited = realVisited
		} // else {
		visited[pos.row+1][pos.column] = true
		minVisited[pos.row+1][pos.column] = true
		//visited = clone(&visitedReal)
		//visited[pos.row][pos.column] = false
		//visited = vis
		//visited = realVisited
		//}
	}
	if d.isNotVisitedAndPossible(pos.column, pos.row-1, visited, currentVal) {
		//fmt.Printf("Going to the bottom :) \n")
		steps, vis := d.makeOneStep(visited, Position{column: pos.column, row: pos.row - 1}, stepsMade+1)

		if steps <= minSteps {
			minVisited = vis
			minSteps = steps

			//vis[pos.row-1][pos.column] = false

			//visited[pos.row][pos.column+1] = false

			//visited[pos.row+1][pos.column] = false

			//visited = clone(&visitedReal)

			//visited = vis
			//minVisited = vis

			//minVisited = realVisited
		} // else {
		visited[pos.row-1][pos.column] = true
		minVisited[pos.row-1][pos.column] = true
		//visited = clone(&visitedReal)
		//visited[pos.row][pos.column] = false
		//visited = vis
		//}
	}

	//fmt.Printf("Steps made: %d\n", minSteps)
	// No step was was possible
	if minSteps == MaxInt {
		return minSteps, visited
	}

	//minVisited[pos.row-1][pos.column] = true
	//minVisited[pos.row+1][pos.column] = true
	//minVisited[pos.row][pos.column+1] = true
	//minVisited[pos.row][pos.column-1] = true
	return minSteps, minVisited
}

func (d *Day) isNotVisitedAndPossible(column int, row int, visited [][]bool, currentPos rune) bool {
	return d.isNotVisited(row, column, visited) && isStepPossible(currentPos, d.heightMap[row][column])
}

func (d *Day) isNotVisited(row int, column int, visited [][]bool) bool {
	if row < 0 || column < 0 || row >= len(d.heightMap) || column >= len(d.heightMap[row]) {
		// We are on standing on the edge
		return false
	}

	// Check if all four directions are already visited
	return !visited[row][column] ||
		(row+1 != len(visited) && !visited[row+1][column] && isStepPossible(d.heightMap[row][column], d.heightMap[row+1][column])) ||
		(column+1 != len(visited[row]) && !visited[row][column+1] && isStepPossible(d.heightMap[row][column], d.heightMap[row][column+1])) ||
		(row != 0 && !visited[row-1][column] && isStepPossible(d.heightMap[row][column], d.heightMap[row-1][column])) ||
		(column != 0 && !visited[row][column-1] && isStepPossible(d.heightMap[row][column], d.heightMap[row][column-1]))
}

// Can we master the height
func isStepPossible(current rune, next rune) bool {
	return next-1 <= current
}

func (d *Day) Part2(in string) string {
	return ""
}

func clone(visited *[][]bool) [][]bool {
	var s = make([][]bool, len(*visited))
	for i := range s {
		s[i] = make([]bool, len((*visited)[i]))
		copy(s[i], (*visited)[i])
	}
	//copy(s, *visited)
	return s
}
