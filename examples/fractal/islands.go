package main

import (
	"image/color"
	"time"

	"github.com/tomowarkar/biome/fractal"
)

var (
	sea    = color.RGBA{65, 105, 225, 255}
	beach  = color.RGBA{240, 230, 140, 255}
	island = color.RGBA{184, 134, 11, 255}
	forest = color.RGBA{46, 139, 87, 255}
	snow   = color.RGBA{255, 255, 255, 255}
)

func main() {
	seed := time.Now().UnixNano()
	world := fractal.Flactal(40, 60, seed)

	var palette []color.Color

	for i := 0; i < 256; i++ {
		if i < 150 {
			palette = append(palette, sea)
		} else if i < 160 {
			palette = append(palette, beach)
		} else if i < 190 {
			palette = append(palette, island)
		} else if i < 240 {
			palette = append(palette, forest)
		} else {
			palette = append(palette, snow)
		}
	}
	world.ToPng("assets/tmp/image", 10, palette)
}
