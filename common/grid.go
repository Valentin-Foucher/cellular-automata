package common

type Grid[E comparable] struct {
	Rows [][]E
	M    int
	N    int
}

func NewGrid[E comparable](m, n int, startValue E) *Grid[E] {
	g := Grid[E]{}
	g.Rows = make([][]E, m)
	for i := range m {
		g.Rows[i] = make([]E, n)
		for j := range n {
			g.Rows[i][j] = startValue
		}
	}
	return &g
}
