package main

import "github.com/tomowarkar/biome"

func main() {
	filename := "assets/tmp/imgproc.jpg"
	imgSrc := biome.ReadImage(filename)

	gray := biome.Gray(imgSrc)
	biome.ToPng("assets/tmp/imgproc_gray", gray)
	sepia := biome.Sepia(imgSrc)
	biome.ToPng("assets/tmp/imgproc_sepia", sepia)
	mosaic := biome.Shade(21, imgSrc)
	biome.ToPng("assets/tmp/imgproc_mosaic", mosaic)

}
