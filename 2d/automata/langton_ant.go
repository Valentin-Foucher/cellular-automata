package automata2d

import grid2d "cellular-automation/2d/grid"

type LangtonAntsAutomaton struct{}

func (a *LangtonAntsAutomaton) NextState(grid grid2d.Grid) grid2d.Grid {
	m, n := grid.Width(), grid.Height()
	ants := grid.GetParticularContent().([]*grid2d.Ant)

	newGrid := grid2d.NewLangtonAntGrid(m, n, len(ants))
	newGrid.Content = grid.GetContent().([][]bool)

	for _, ant := range ants {
		isWhite := grid.Get(ant.X, ant.Y).(bool)
		newGrid.Content[ant.X][ant.Y] = !isWhite
		moveAnt(isWhite, ant, m, n)

	}
	newGrid.ParticularContent = ants
	return newGrid
}

func moveAnt(isWhite bool, ant *grid2d.Ant, m, n int) {
	if isWhite {
		ant.Direction = (ant.Direction + 1) % 4
	} else if ant.Direction == 0 {
		ant.Direction = 3
	} else {
		ant.Direction = (ant.Direction - 1) % 4
	}

	switch grid2d.AntDirections[ant.Direction] {
	case "u":
		if ant.Y > 0 {
			ant.Y--
			return
		}
		ant.Y = n - 1
	case "r":
		if ant.X != m-1 {
			ant.X++
			return
		}
		ant.X = 0
	case "d":
		if ant.Y != n-1 {
			ant.Y++
			return
		}
		ant.Y = 0
	case "l":
		if ant.X > 0 {
			ant.X--
			return
		}
		ant.X = m - 1
	default:
		panic("invalid direction")
	}
}
