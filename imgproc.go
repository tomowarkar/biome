package biome

import (
	"image"
	"image/color"
	"log"
	"math"
)

// Gray :
func Gray(imgSrc image.Image) image.Image {
	w, h := imgSrc.Bounds().Dx(), imgSrc.Bounds().Dy()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rr, gg, bb, _ := imgSrc.At(x, y).RGBA()
			v := grayRGB(rr, gg, bb)
			img.Set(x, y, color.RGBA{v, v, v, 255})
		}
	}
	return img
}

// Shade :
func Shade(oddnum int, imgSrc image.Image) image.Image {
	if oddnum%2 == 0 || oddnum < 1 {
		log.Fatal("1より大きい奇数を入れてね")
	}
	n := oddnum / 2
	w, h := imgSrc.Bounds().Dx(), imgSrc.Bounds().Dy()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var rr, gg, bb uint32
			var count float64
			for i := MaxInt(0, x-n); i < MinInt(w-1, x+n); i++ {
				for j := MaxInt(0, y-n); j < MinInt(h-1, y+n); j++ {
					r, g, b, _ := imgSrc.At(i, j).RGBA()
					rr += r
					gg += g
					bb += b
					count++
				}
			}
			img.Set(x, y, color.RGBA{F2uint8(float64(rr) / count), F2uint8(float64(gg) / count), F2uint8(float64(bb) / count), 255})
		}
	}
	return img
}

// Sepia :
func Sepia(imgSrc image.Image) image.Image {
	w, h := imgSrc.Bounds().Dx(), imgSrc.Bounds().Dy()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rr, gg, bb, _ := imgSrc.At(x, y).RGBA()
			r, g, b := sepiaRGB(rr, gg, bb)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	return img
}

func grayRGB(rr, gg, bb uint32) uint8 {
	m := cieXYZ(rr, gg, bb)
	return F2uint8(m)
}

func cieXYZ(rr, gg, bb uint32) float64 {
	r := math.Pow(float64(rr), 2.2)
	g := math.Pow(float64(gg), 2.2)
	b := math.Pow(float64(bb), 2.2)
	return math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
}

func sepiaRGB(rr, gg, bb uint32) (r, g, b uint8) {
	m := cieXYZ(rr, gg, bb)
	r = F2uint8(m * 107 / 107)
	g = F2uint8(m * 74 / 107)
	b = F2uint8(m * 43 / 107)
	return
}
