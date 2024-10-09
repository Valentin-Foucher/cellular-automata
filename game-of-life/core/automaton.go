package gameoflife

import (
	"cellular-automation/common"
	"errors"
)

type Automaton = common.Automaton[bool]

type GameOfLifeAutomaton struct{}

func (a *GameOfLifeAutomaton) NextState(grid *Grid) *Grid {
	newGrid := newGrid(grid.M, grid.N)

	for i := range grid.M {
		for j := range grid.N {
			newGrid.Rows[i][j] = a.nextStateForCell(i, j, grid.M, grid.N, grid.Rows)
		}
	}
	return newGrid
}

func (a *GameOfLifeAutomaton) nextStateForCell(i, j, m, n int, rows [][]bool) bool {
	alive := rows[i][j]
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
		return alive
	case 3:
		return true
	default:
		return false
	}
}

func getAutomaton(conf *Configuration) (Automaton, error) {
	switch conf.Automaton {
	case "gol":
		return &GameOfLifeAutomaton{}, nil
	default:
		return nil, errors.New("invalid automaton configuraton")
	}
}
