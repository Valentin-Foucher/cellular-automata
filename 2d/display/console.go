package display2d

import (
	"cellular-automation/common"
	"errors"
	"fmt"
)

type ConsoleDisplay struct{}

func (*ConsoleDisplay) ShowGrid(g *common.BaseGrid) {
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

func (*ConsoleDisplay) EraseGrid() {
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
