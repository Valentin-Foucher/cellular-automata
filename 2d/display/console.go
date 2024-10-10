package display2d

import (
	grid2d "cellular-automation/2d/grid"
	"fmt"
)

type ConsoleDisplay struct{}

func (*ConsoleDisplay) Print(g grid2d.Grid) {
	for i := range g.Width() {
		for j := range g.Height() {
			switch v := g.Get(i, j).(type) {
			case bool:
				if v {
					fmt.Print("⬜️")
				} else {
					fmt.Print("  ")
				}
			case int:
				switch v {
				case grid2d.On:
					fmt.Print("🟥")
				case grid2d.Dying:
					fmt.Print("🟨")
				default:
					fmt.Print("  ")
				}
			}
		}

		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (*ConsoleDisplay) Flush() {
	fmt.Print("\033[H\033[2J")
}
