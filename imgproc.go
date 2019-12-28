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
func Shade(num int, imgSrc image.Image) image.Image {
	w, h := imgSrc.Bounds().Dx(), imgSrc.Bounds().Dy()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	done := make(chan bool, h)
	for y := 0; y < h; y++ {
		go func(line int) {
			for x := 0; x < w; x++ {
				r, g, b := calc(x, line, num, w, h, imgSrc)
				img.Set(x, line, color.RGBA{r, g, b, 255})
			}
			done <- true
		}(y)
	}
	for i := 0; i < h; i++ {
		<-done
	}
	return img
}

func calc(x, y, num, w, h int, imgSrc image.Image) (rr, gg, bb uint8) {
	if num < 2 {
		log.Fatal("1より大きい奇数を入れてね")
	}
	n := num / 2
	var count float64
	var r1, g1, b1 uint32
	for i := MaxInt(0, x-n); i < MinInt(w-1, x+n); i++ {
		for j := MaxInt(0, y-n); j < MinInt(h-1, y+n); j++ {
			r, g, b, _ := imgSrc.At(i, j).RGBA()
			r1 += r
			g1 += g
			b1 += b
			count++
		}
	}
	rr = F2uint8(float64(r1) / count)
	gg = F2uint8(float64(g1) / count)
	bb = F2uint8(float64(b1) / count)
	return
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
