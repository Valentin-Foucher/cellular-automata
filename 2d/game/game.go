package game2d

import (
	automata2d "cellular-automation/2d/automata"
	display2d "cellular-automation/2d/display"
	grid2d "cellular-automation/2d/grid"
	"time"
)

const (
	nanosecondsInOneSecond = 1_000_000_000
)

type Game struct {
	grid      *grid2d.BaseGrid
	display   display2d.Display
	automaton automata2d.BaseAutomaton
	maxSteps  int
	sleepTime float32
}

func NewGame(conf *Configuration) (*Game, error) {
	display, err := display2d.Get(conf.Display)
	if err != nil {
		return nil, err
	}
	automaton, err := automata2d.Get(conf.Automaton)
	if err != nil {
		return nil, err
	}

	return &Game{
		grid:      grid2d.Initialize(conf.M, conf.N, conf.Distribution),
		maxSteps:  conf.Steps,
		display:   display,
		automaton: automaton,
		sleepTime: conf.SleepTime,
	}, nil
}

func (g *Game) Start() {
	g.display.Show(g.grid)
	for range g.maxSteps {
		g.grid = g.automaton.NextState(g.grid)
		g.display.Erase()
		g.display.Show(g.grid)
		time.Sleep(time.Duration(g.sleepTime * nanosecondsInOneSecond))
	}
}
