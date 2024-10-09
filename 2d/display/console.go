package display2d

import (
	grid2d "cellular-automation/2d/grid"
	"errors"
	"fmt"
)

type ConsoleDisplay struct{}

func (*ConsoleDisplay) Show(g *grid2d.BaseGrid) {
	for i := range g.M {
		for j := range g.N {
			if g.Content[i][j] {
				fmt.Print("⬜️")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (*ConsoleDisplay) Erase() {
	fmt.Print("\033[H\033[2J")
}

func Get(displayType string) (Display, error) {
	switch displayType {
	case "console":
		return &ConsoleDisplay{}, nil
	default:
		return nil, errors.New("invalid display configuraton")
	}
}
