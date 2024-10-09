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
)

type GifDisplay struct {
	frames []*image.Paletted
	delay  int
}

func (d *GifDisplay) Print(g grid2d.Grid) {
	frame := image.NewPaletted(image.Rect(0, 0, g.Width(), g.Height()), []color.Color{black, white})
	for i := range g.Width() {
		for j := range g.Height() {
			var c color.Color
			if g.Get(i, j).(bool) {
				c = white
			} else {
				c = black
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
