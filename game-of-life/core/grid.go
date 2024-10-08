package gameoflife

import (
	"cellular-automation/common"
	"time"

	"golang.org/x/exp/rand"
)

type Grid = common.Grid[bool]

func newGrid(m, n int) *Grid {
	return common.NewGrid[bool](m, n)
}

func initializeGrid(m, n int, distribution float32) *Grid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := newGrid(m, n)
	for i := range m {
		for j := range n {
			g.Rows[i][j] = rand.Float32() < distribution
		}
	}

	return g
}
