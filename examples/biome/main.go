package main

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/tomowarkar/biome"
)

func main() {
	w, h := 8, 5
	data := make([]int, w*h)
	palette := []color.Color{
		color.RGBA{255, 0, 255, 255},
		color.RGBA{0, 255, 255, 255},
		color.RGBA{255, 255, 0, 255},
	}
	img := biome.Slice2Image(biome.Biome{W: w, H: h, Data: data}, 10, palette)
	biome.ToPng("assets/tmp/fuck", img)

	var images []*image.Paletted
	for i := 0; i < 10; i++ {
		data := make([]int, w*h)
		for j := 0; j < w*h; j++ {
			data[j] = rand.Intn(len(palette))
		}
		images = append(images,
			biome.Slice2Paletted(biome.Biome{W: w, H: h, Data: data}, 10, palette))
	}
	biome.ToGif("assets/tmp/hogeg", images, 100)
}
