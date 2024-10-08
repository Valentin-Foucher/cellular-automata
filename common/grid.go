package common

type Grid[E comparable] struct {
	Rows [][]E
	M    int
	N    int
}

func NewGrid[E comparable](m, n int) *Grid[E] {
	g := Grid[E]{
		M:    m,
		N:    n,
		Rows: make([][]E, m),
	}
	for i := range m {
		g.Rows[i] = make([]E, n)
	}

	return &g
}
