package automata2d

import (
	grid2d "cellular-automation/2d/grid"
	"errors"
)

type Automaton2D[E comparable] interface {
	NextState(grid2d.Grid) grid2d.Grid
}

type BaseAutomaton = Automaton2D[bool]

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
