package gameoflife

import "time"

type Game struct {
	Grid      *Grid
	Steps     int
	display   Display
	sleepTime float32
}

func NewGame(conf *Configuration) (*Game, error) {
	display, err := getDisplay(conf)
	if err != nil {
		return nil, err
	}
	return &Game{
		Grid:      initializeGrid(conf.M, conf.N, conf.Distribution),
		Steps:     conf.Steps,
		display:   display,
		sleepTime: conf.SleepTime,
	}, nil
}

func (g *Game) Start() {
	g.display.ShowGrid(g.Grid)
	for range g.Steps {
		g.nextState()
		g.display.EraseGrid(g.Grid)
		g.display.ShowGrid(g.Grid)
		time.Sleep(time.Duration(g.sleepTime * 1_000_000_000))
	}
}

func (g *Game) nextState() {
	newGrid := newGrid(g.Grid.M, g.Grid.N)

	for i := range g.Grid.M {
		for j := range g.Grid.N {
			newGrid.Rows[i][j] = nextStateForCell(i, j, g.Grid.M, g.Grid.N, g.Grid.Rows)
		}
	}
	g.Grid = newGrid
}

func nextStateForCell(i, j, m, n int, rows [][]bool) bool {
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
