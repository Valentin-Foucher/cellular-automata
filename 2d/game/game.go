package game2d

import (
	automata2d "cellular-automation/2d/automata"
	config2d "cellular-automation/2d/configs"
	display2d "cellular-automation/2d/display"
	grid2d "cellular-automation/2d/grid"
	"time"
)

const (
	nanosecondsInOneSecond = 1_000_000_000
)

type Game struct {
	grid      grid2d.Grid
	display   display2d.Display
	automaton automata2d.Automaton2D
	maxSteps  int
	sleepTime float32
}

func NewGame(conf *config2d.Configuration) (*Game, error) {
	display, err := display2d.Get(conf)
	if err != nil {
		return nil, err
	}
	automaton, err := automata2d.Get(conf)
	if err != nil {
		return nil, err
	}
	grid, err := grid2d.Initialize(conf)
	if err != nil {
		return nil, err
	}

	return &Game{
		grid:      grid,
		maxSteps:  conf.Steps,
		display:   display,
		automaton: automaton,
		sleepTime: conf.SleepTime,
	}, nil
}

func (g *Game) Start() {
	if _, ok := g.display.(*display2d.ConsoleDisplay); ok {
		g.liveStrategy()
	} else {
		g.deferredStrategy()
	}
}

func (g *Game) liveStrategy() {
	g.display.Print(g.grid)
	for range g.maxSteps {
		g.grid = g.automaton.NextState(g.grid)
		g.display.Flush()
		g.display.Print(g.grid)
		time.Sleep(time.Duration(g.sleepTime * nanosecondsInOneSecond))
	}
}

func (g *Game) deferredStrategy() {
	for range g.maxSteps {
		g.display.Print(g.grid)
		g.grid = g.automaton.NextState(g.grid)
	}
	g.display.Flush()
}
