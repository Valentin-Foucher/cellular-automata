package automata2d

import "cellular-automation/common"

type SeedsAutomaton struct{}

func (a *SeedsAutomaton) NextState(grid *common.BaseGrid) *common.BaseGrid {
	newGrid := common.NewBaseGrid(grid.M, grid.N)

	for i := range grid.M {
		for j := range grid.N {
			newGrid.Rows[i][j] = a.nextStateForCell(i, j, grid.M, grid.N, grid.Rows)
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
