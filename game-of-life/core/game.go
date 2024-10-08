package gameoflife

type Game struct {
	Grid *Grid
	Step int
}

func NewGame(c *Configuration) *Game {
	return &Game{
		Grid: NewGrid(c.M, c.N),
	}
}

func (g *Game) Start() {

}
