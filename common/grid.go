package common

type TwoDimensionalGrid[E comparable] struct {
	Rows [][]E
	M    int
	N    int
}

func NewTwoDimensionalGrid[E comparable](m, n int) *TwoDimensionalGrid[E] {
	g := TwoDimensionalGrid[E]{
		M:    m,
		N:    n,
		Rows: make([][]E, m),
	}
	for i := range m {
		g.Rows[i] = make([]E, n)
	}

	return &g
}
