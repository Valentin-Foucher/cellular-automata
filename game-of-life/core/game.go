package gameoflife

import "time"

const (
	nanosecondsInOneSecond = 1_000_000_000
)

type Game struct {
	Grid      *Grid
	Steps     int
	display   Display
	automaton Automaton
	sleepTime float32
}

func NewGame(conf *Configuration) (*Game, error) {
	display, err := getDisplay(conf)
	if err != nil {
		return nil, err
	}
	automaton, err := getAutomaton(conf)
	if err != nil {
		return nil, err
	}

	return &Game{
		Grid:      initializeGrid(conf.M, conf.N, conf.Distribution),
		Steps:     conf.Steps,
		display:   display,
		automaton: automaton,
		sleepTime: conf.SleepTime,
	}, nil
}

func (g *Game) Start() {
	g.display.ShowGrid(g.Grid)
	for range g.Steps {
		g.Grid = g.automaton.NextState(g.Grid)
		g.display.EraseGrid()
		g.display.ShowGrid(g.Grid)
		time.Sleep(time.Duration(g.sleepTime * nanosecondsInOneSecond))
	}
}
