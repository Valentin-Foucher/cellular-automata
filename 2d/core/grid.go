package core2d

import (
	"cellular-automation/common"
	"time"

	"golang.org/x/exp/rand"
)

func initializeGrid(m, n int, distribution float32) *common.BaseGrid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := common.NewBaseGrid(m, n)
	for i := range m {
		for j := range n {
			g.Rows[i][j] = rand.Float32() < distribution
		}
	}

	return g
}
