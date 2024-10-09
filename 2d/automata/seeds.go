package automata2d

import grid2d "cellular-automation/2d/grid"

type SeedsAutomaton struct{}

func (a *SeedsAutomaton) NextState(grid *grid2d.BaseGrid) *grid2d.BaseGrid {
	newGrid := grid2d.NewBaseGrid(grid.M, grid.N)

	for i := range grid.M {
		for j := range grid.N {
			newGrid.Content[i][j] = a.nextStateForCell(i, j, grid.M, grid.N, grid.Content)
		}
	}
	return newGrid
}

func (a *SeedsAutomaton) nextStateForCell(i, j, m, n int, rows [][]bool) bool {
	alive := rows[i][j]
	if alive {
		return false
	}
	adjacentAliveCells := 0

	for _, k := range []int{-1, 0, 1} {
		for _, l := range []int{-1, 0, 1} {
			if i+k < 0 || i+k > m-1 ||
				j+l < 0 || j+l > n-1 ||
				(k == 0 && l == 0) {
				continue
			}

			if rows[i+k][j+l] {
				adjacentAliveCells++
			}
		}
	}

	switch adjacentAliveCells {
	case 2:
		return true
	default:
		return false
	}
}
