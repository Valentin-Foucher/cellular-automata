package grid2d

import (
	"cellular-automation/common"
	"time"

	"golang.org/x/exp/rand"
)

type TwoDimensionalGrid[E comparable] common.World[[][]E]

type BaseGrid = TwoDimensionalGrid[bool]

func NewTwoDimensionalGrid[E comparable](m, n int) *TwoDimensionalGrid[E] {
	g := TwoDimensionalGrid[E]{
		M:       m,
		N:       n,
		Content: make([][]E, m),
	}
	for i := range m {
		g.Content[i] = make([]E, n)
	}

	return &g
}

func NewBaseGrid(m, n int) *BaseGrid {
	return NewTwoDimensionalGrid[bool](m, n)
}

func Initialize(m, n int, distribution float32) *BaseGrid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := NewBaseGrid(m, n)
	for i := range m {
		for j := range n {
			g.Content[i][j] = rand.Float32() < distribution
		}
	}

	return g
}
