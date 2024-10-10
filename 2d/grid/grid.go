package grid2d

import (
	config2d "cellular-automation/2d/configs"
	"cellular-automation/common"
	"errors"
)

type Grid interface {
	Get(i, j int) any
	Width() int
	Height() int
	GetContent() any
	GetParticularContent() any
}

type BaseGrid[E comparable, F any] common.World[[][]E, F]

func (g *BaseGrid[E, F]) Get(i, j int) any {
	return g.Content[i][j]
}

func (g *BaseGrid[E, F]) Width() int {
	return g.M
}

func (g *BaseGrid[E, F]) Height() int {
	return g.N
}

func (g *BaseGrid[E, F]) GetContent() any {
	return g.Content
}

func (g *BaseGrid[E, F]) GetParticularContent() any {
	return g.ParticularContent
}

func NewBaseGrid[E comparable, F any](m, n int) *BaseGrid[E, F] {
	g := BaseGrid[E, F]{
		M:       m,
		N:       n,
		Content: make([][]E, m),
	}
	for i := range m {
		g.Content[i] = make([]E, n)
	}

	return &g
}

func Initialize(conf *config2d.Configuration) (Grid, error) {
	switch conf.Automaton {
	case "gol", "seeds", "day_and_night":
		return initializeBooleanGrid(conf.M, conf.N, conf.Distribution), nil
	case "langton_ants":
		return initializeLangtonAntGrid(conf.M, conf.N, conf.AntsCount), nil

	default:
		return nil, errors.New("invalid automaton configuration")
	}
}
