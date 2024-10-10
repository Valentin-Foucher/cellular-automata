package automata2d

import (
	grid2d "cellular-automation/2d/grid"
)

type GameOfLifeAutomaton struct{}

func (a *GameOfLifeAutomaton) NextState(grid grid2d.Grid) grid2d.Grid {
	return nextStateForAllBoolCells(grid, a.nextStateForCell)
}

func (a *GameOfLifeAutomaton) nextStateForCell(i, j, m, n int, get func(k, l int) bool) bool {
	alive := get(i, j)
	adjacentAliveCells := 0

	for _, k := range []int{-1, 0, 1} {
		for _, l := range []int{-1, 0, 1} {
			if i+k < 0 || i+k > m-1 ||
				j+l < 0 || j+l > n-1 ||
				(k == 0 && l == 0) {
				continue
			}

			if get(i+k, j+l) {
				adjacentAliveCells++
			}
		}
	}

	switch adjacentAliveCells {
	case 2:
		return alive
	case 3:
		return true
	default:
		return false
	}
}
