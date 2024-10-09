package automata2d

import (
	grid2d "cellular-automation/2d/grid"
	"errors"
)

type Automaton2D[E comparable] interface {
	NextState(grid2d.Grid) grid2d.Grid
}

type BaseAutomaton = Automaton2D[bool]

func nextStateForAllCells(grid grid2d.Grid, nextStateForCell func(int, int, int, int, func(k, l int) bool) bool) grid2d.Grid {
	m, n := grid.Width(), grid.Height()
	newGrid := grid2d.NewBooleanGrid(m, n)

	for i := range m {
		for j := range n {
			newGrid.Content[i][j] = nextStateForCell(
				i,
				j,
				m,
				n,
				func(k, l int) bool { return grid.Get(k, l).(bool) },
			)
		}
	}
	return newGrid
}

func Get(automatonType string) (BaseAutomaton, error) {
	switch automatonType {
	case "gol":
		return &GameOfLifeAutomaton{}, nil
	case "seeds":
		return &SeedsAutomaton{}, nil
	case "langton_ants":
		return &LangtonAntsAutomaton{}, nil
	default:
		return nil, errors.New("invalid automaton configuraton")
	}
}
