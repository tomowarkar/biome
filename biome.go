package biome

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
)

// TODO: インターフェイスを使ったらSlice2ImageとSlice2Palettedをマージできるかも??

// Slice2Image : １次元Intスライスを元にimage.Imageを生成
func Slice2Image(width, height, scale int, data []int, palette []color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width*scale, height*scale))
	for x := 0; x < width*scale; x++ {
		for y := 0; y < height*scale; y++ {
			img.Set(x, y, palette[data[y/scale*width+x/scale]%len(palette)])
		}
	}
	return img
}

// Slice2Paletted : １次元Intスライスを元にimage.Palettedを生成
func Slice2Paletted(width, height, scale int, data []int, palette []color.Color) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, width*scale, height*scale), palette)
	for x := 0; x < width*scale; x++ {
		for y := 0; y < height*scale; y++ {
			img.Set(x, y, palette[data[y/scale*width+x/scale]%len(palette)])
		}
	}
	return img
}

// ToPng : １次元Intスライスを元にpng画像を生成
func ToPng(filename string, imgSrc image.Image) {
	wf, err := os.Create(filename + ".png")
	if err != nil {
		panic("encode failed")
	}
	defer wf.Close()
	png.Encode(wf, imgSrc)
}

// ToGif : 2次元Intスライスを元にGIF画像を生成
func ToGif(filename string, images []*image.Paletted, delay int) {
	var delays []int
	for i := 0; i < len(images); i++ {
		delays = append(delays, delay)
	}

	f, _ := os.OpenFile(filename+".gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
