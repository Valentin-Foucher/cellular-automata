package automata2d

import (
	grid2d "cellular-automation/2d/grid"
)

type DayAndNightAutomaton struct{}

func (a *DayAndNightAutomaton) NextState(grid grid2d.Grid) grid2d.Grid {
	return nextStateForAllCells(grid, a.nextStateForCell)
}

func (a *DayAndNightAutomaton) nextStateForCell(i, j, m, n int, get func(k, l int) bool) bool {
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
	if alive {
		switch adjacentAliveCells {
		case 3, 4, 6, 7, 8:
			return true
		default:
			return false
		}
	}

	switch adjacentAliveCells {
	case 3, 6, 7, 8:
		return true
	default:
		return false
	}
}
