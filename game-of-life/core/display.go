package gameoflife

import (
	"cellular-automation/common"
	"errors"
	"fmt"
)

type Display = common.Display[bool]

type ConsoleDisplay struct{}

func (*ConsoleDisplay) ShowGrid(g *Grid) {
	for i := range g.M {
		for j := range g.N {
			if g.Rows[i][j] {
				fmt.Print("⬜️")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (*ConsoleDisplay) EraseGrid(g *Grid) {
	fmt.Print("\033[H\033[2J")
}

func getDisplay(conf *Configuration) (Display, error) {
	switch conf.Display {
	case "console":
		return &ConsoleDisplay{}, nil
	default:
		return nil, errors.New("invalid display configuraton")
	}
}
