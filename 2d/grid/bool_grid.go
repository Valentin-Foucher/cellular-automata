package grid2d

import (
	"time"

	"golang.org/x/exp/rand"
)

type BooleanGrid = BaseGrid[bool, bool]

func NewBooleanGrid(m, n int) *BooleanGrid {
	return NewBaseGrid[bool, bool](m, n)
}

func initializeBooleanGrid(m, n int, distribution float32) Grid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := NewBooleanGrid(m, n)
	for i := range m {
		for j := range n {
			g.Content[i][j] = rand.Float32() < distribution
		}
	}

	return g
}
