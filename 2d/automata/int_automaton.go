package automata2d

import grid2d "cellular-automation/2d/grid"

func nextStateForAllIntCells(grid grid2d.Grid, nextStateForCell func(int, int, int, int, func(k, l int) int) int) grid2d.Grid {
	m, n := grid.Width(), grid.Height()
	newGrid := grid2d.NewIntGrid(m, n)

	for i := range m {
		for j := range n {
			newGrid.Content[i][j] = nextStateForCell(
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
