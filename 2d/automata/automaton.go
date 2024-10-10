package automata2d

import (
	config2d "cellular-automation/2d/configs"
	grid2d "cellular-automation/2d/grid"
	"errors"
)

type Automaton2D interface {
	NextState(grid2d.Grid) grid2d.Grid
}

func Get(conf *config2d.Configuration) (Automaton2D, error) {
	switch conf.Automaton {
	case "gol":
		return &GameOfLifeAutomaton{}, nil
	case "seeds":
		return &SeedsAutomaton{}, nil
	case "langton_ants":
		return &LangtonAntsAutomaton{}, nil
	case "day_and_night":
		return &DayAndNightAutomaton{}, nil
	case "brians_brain":
		return &BriansBrainAutomaton{}, nil
	default:
		return nil, errors.New("invalid automaton configuraton")
	}
}
