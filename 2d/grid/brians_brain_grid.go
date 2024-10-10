package grid2d

import (
	"time"

	"golang.org/x/exp/rand"
)

const (
	Off   = 0
	On    = 1
	Dying = 2
)

func initializeBriansBrainGrid(m, n int, distribution float32) Grid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := NewIntGrid(m, n)

	for i := range m {
		for j := range n {
			var value int
			if rand.Float32() >= distribution {
				value = Off
			} else if rand.Float32() >= distribution {
				value = Dying
			} else {
				value = On
			}
			g.Content[i][j] = value
		}
	}

	return g
}
