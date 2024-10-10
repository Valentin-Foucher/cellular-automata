package automata2d

import (
	grid2d "cellular-automation/2d/grid"
)

type BriansBrainAutomaton struct{}

func (a *BriansBrainAutomaton) NextState(grid grid2d.Grid) grid2d.Grid {
	m, n := grid.Width(), grid.Height()
	newGrid := grid2d.NewBriansBrainGrid(m, n)

	for i := range m {
		for j := range n {
			newGrid.Content[i][j] = a.nextStateForCell(
				i,
				j,
				m,
				n,
				func(k, l int) int { return grid.Get(k, l).(int) },
			)
		}
	}
	return newGrid
}

func (a *BriansBrainAutomaton) nextStateForCell(i, j, m, n int, get func(k, l int) int) int {
	state := get(i, j)
	if state == grid2d.On {
		return grid2d.Dying
	}
	adjacentAliveCells := 0

	for _, k := range []int{-1, 0, 1} {
		for _, l := range []int{-1, 0, 1} {
			if i+k < 0 || i+k > m-1 ||
				j+l < 0 || j+l > n-1 ||
				(k == 0 && l == 0) {
				continue
			}

			if get(i+k, j+l) == grid2d.On {
				adjacentAliveCells++
			}
		}
	}

	switch adjacentAliveCells {
	case 2:
		return grid2d.On
	default:
		return grid2d.Off
	}
}
