package grid2d

import (
	"time"

	"golang.org/x/exp/rand"
)

type (
	LangtonAntGrid = BaseGrid[bool, *Ant]
	Ant            struct {
		X         int
		Y         int
		Direction int
	}
)

var (
	AntDirections = [4]string{
		"u",
		"r",
		"d",
		"l",
	}
)

func NewLangtonAntGrid(m, n, ants int) *LangtonAntGrid {
	g := NewBaseGrid[bool, *Ant](m, n)
	g.ParticularContent = make([]*Ant, ants)
	return g
}

func initializeLangtonAntGrid(m, n int, ants int) *LangtonAntGrid {
	rand.Seed(uint64(time.Now().UnixNano()))

	g := NewLangtonAntGrid(m, n, ants)
	for i := range ants {
		antX, antY := rand.Intn(m), rand.Intn(n)
		g.ParticularContent[i] = &Ant{
			X:         antX,
			Y:         antY,
			Direction: rand.Intn(len(AntDirections)),
		}
	}

	return g
}
