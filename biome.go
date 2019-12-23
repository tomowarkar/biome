package biome

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"strconv"
)

// ToPng : １次元Intスライスを元にpng画像を生成
func ToPng(filename string, width, height, scale int, data []int, palette []color.Color) {
	img := image.NewRGBA(image.Rect(0, 0, width*scale, height*scale))
	for x := 0; x < width*scale; x++ {
		for y := 0; y < height*scale; y++ {
			img.Set(x, y, palette[data[y/scale*width+x/scale]%len(palette)])
		}
	}
	encodePng(img, filename)
}

func encodePng(img *image.RGBA, path string) {
	f, err := os.Create(path + ".png")
	if err != nil {
		panic("encode failed")
	}
	defer f.Close()
	png.Encode(f, img)
	return
}

// ToGif : 2次元Intスライスを元にGIF画像を生成
func ToGif(filename string, w, h, scale, delay int, items [][]int, palette []color.Color) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < len(items); i++ {
		img := image.NewPaletted(image.Rect(0, 0, w*scale, h*scale), palette)
		images = append(images, img)
		delays = append(delays, delay)
		for x := 0; x < w*scale; x++ {
			for y := 0; y < h*scale; y++ {
				img.Set(x, y, palette[items[i][y/scale*w+x/scale]%len(palette)])
			}
		}
	}

	f, _ := os.OpenFile(filename+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

// ToStr : convert int to string
func ToStr(n int) string {
	return strconv.Itoa(n)
}
