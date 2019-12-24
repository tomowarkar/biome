package main

import (
	"image/color"

	"github.com/tomowarkar/biome/paint"
)

const (
	white = 0
	black = 1
	blue  = 2
	green = 3
	red   = 4
)

var palette = []color.Color{
	color.RGBA{255, 255, 255, 255},
	color.RGBA{0, 0, 0, 255},
	color.RGBA{100, 100, 225, 255},
	color.RGBA{100, 225, 100, 255},
	color.RGBA{225, 100, 100, 255},
}

func main() {
	width, height := 1200, 800
	cnv := paint.NewCanv(width, height)
	cnv.Dot(600, 400, 300, blue)
	cnv.Dot(750, 300, 240, red)
	cnv.Dot(500, 400, 180, green)
	cnv.Dot(600, 400, 40, red)
	cnv.Fill(0, 0, green)
	cnv.Line(827, 427, 1100, 527, 150, red)
	cnv.Square(300, 100, 300, 30, white)
	cnv.Fill(300, 100, red)
	cnv.Square(300, 700, 300, 80, white)
	cnv.Triangle(1000, 300, 300, 10, white)
	cnv.Fill(1000, 300, blue)
	cnv.ToPng("paint", palette)
}
