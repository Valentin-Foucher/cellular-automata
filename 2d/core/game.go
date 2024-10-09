package core2d

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
	Grid      *grid2d.BaseGrid
	Steps     int
	display   display2d.Display
	automaton automata2d.BaseAutomaton
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
		Grid:      grid2d.Initialize(conf.M, conf.N, conf.Distribution),
		Steps:     conf.Steps,
		display:   display,
		automaton: automaton,
		sleepTime: conf.SleepTime,
	}, nil
}

func (g *Game) Start() {
	g.display.Show(g.Grid)
	for range g.Steps {
		g.Grid = g.automaton.NextState(g.Grid)
		g.display.Erase()
		g.display.Show(g.Grid)
		time.Sleep(time.Duration(g.sleepTime * nanosecondsInOneSecond))
	}
}
