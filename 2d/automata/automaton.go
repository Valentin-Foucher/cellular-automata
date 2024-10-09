package automata2d

import (
	"cellular-automation/common"
	"errors"
)

type Automaton2D[E comparable] interface {
	NextState(*common.TwoDimensionalGrid[E]) *common.TwoDimensionalGrid[E]
}

type BaseAutomaton = Automaton2D[bool]

func Get(automatonType string) (BaseAutomaton, error) {
	switch automatonType {
	case "gol":
		return &GameOfLifeAutomaton{}, nil
	case "seeds":
		return &SeedsAutomaton{}, nil
	default:
		return nil, errors.New("invalid automaton configuraton")
	}
}
