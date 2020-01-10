package biome

import (
	"image"
	"image/color"
)

func max(arr []int) (res int) {
	res = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return
}

func min(arr []int) (res int) {
	res = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}
	return
}

func gradation(ratio float64) color.RGBA {
	if ratio > 100 {
		ratio = 100
	} else if ratio < 0 {
		ratio = 0
	}

	base := color.RGBA{255, 255, 85, 255}
	target := color.RGBA{70, 0, 70, 255}
	ar, ag, ab, _ := base.RGBA()
	br, bg, bb, _ := target.RGBA()
	r := F2uint8((float64(br)-float64(ar))*ratio/100 + float64(ar))
	g := F2uint8((float64(bg)-float64(ag))*ratio/100 + float64(ag))
	b := F2uint8((float64(bb)-float64(ab))*ratio/100 + float64(ab))
	return color.RGBA{r, g, b, 255}
}

// DebugData ...
func DebugData(width, height, scale int, data []int) {
	dataMax := max(data)
	dataMin := min(data)
	if dataMax-dataMin == 0 {
		dataMax++
	}

	img := image.NewRGBA(image.Rect(0, 0, width*scale, height*scale))
	for y := 0; y < height*scale; y++ {
		for x := 0; x < width*scale; x++ {
			var num float64
			ndata := data[y/scale*width+x/scale]
			num = 100 / float64(dataMax-dataMin) * float64(ndata)
			img.Set(x, y, gradation(num))
		}
	}
	ToPng("debug", img)
}
