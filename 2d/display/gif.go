package display2d

import (
	grid2d "cellular-automation/2d/grid"
	"fmt"
	"image"
	"os"
	"time"

	"image/color"
	"image/gif"
)

var (
	black = color.Black
	white = color.White
	red   = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 1,
	}
	yellow = color.RGBA{
		R: 255,
		G: 255,
		B: 128,
		A: 1,
	}
)

type GifDisplay struct {
	frames []*image.Paletted
	delay  int
}

func (d *GifDisplay) Print(g grid2d.Grid) {
	frame := image.NewPaletted(image.Rect(0, 0, g.Width(), g.Height()), []color.Color{black, white, red, yellow})
	for i := range g.Width() {
		for j := range g.Height() {
			var c color.Color

			switch v := g.Get(i, j).(type) {
			case bool:
				if v {
					c = white
				} else {
					c = black
				}
			case int:
				switch v {
				case grid2d.On:
					c = red
				case grid2d.Dying:
					c = yellow
				default:
					c = black
				}
			}
			frame.Set(i, j, c)
		}

		fmt.Print("\n")
	}
	fmt.Print("\n")
	d.frames = append(d.frames, frame)
}

func (d *GifDisplay) Flush() {
	f, err := os.OpenFile(fmt.Sprintf("compiled/%s.gif", time.Now().Format(time.RFC3339)), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := len(d.frames)
	delays := make([]int, n)
	for i := range n {
		delays[i] = d.delay
	}

	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: d.frames,
		Delay: delays,
	})
}
