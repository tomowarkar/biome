package biome

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
)

// Biome ...
type Biome struct {
	W    int
	H    int
	Data []int
}

// TODO: インターフェイスを使ったらSlice2ImageとSlice2Palettedをマージできるかも??

// Slice2Image : １次元Intスライスを元にimage.Imageを生成
func Slice2Image(biome Biome, scale int, palette []color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, biome.W*scale, biome.H*scale))
	for x := 0; x < biome.W*scale; x++ {
		for y := 0; y < biome.H*scale; y++ {
			img.Set(x, y, palette[biome.Data[y/scale*biome.W+x/scale]%len(palette)])
		}
	}
	return img
}

// Slice2Paletted : １次元Intスライスを元にimage.Palettedを生成
func Slice2Paletted(biome Biome, scale int, palette []color.Color) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, biome.W*scale, biome.H*scale), palette)
	for x := 0; x < biome.W*scale; x++ {
		for y := 0; y < biome.H*scale; y++ {
			img.Set(x, y, palette[biome.Data[y/scale*biome.W+x/scale]%len(palette)])
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

// ReadImage :
func ReadImage(path string) (imgSrc image.Image) {
	rf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer rf.Close()

	buffer := make([]byte, 512)
	_, err = rf.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	cType := http.DetectContentType(buffer)
	rf.Seek(0, 0)

	switch cType {
	case "image/png":
		imgSrc, err = png.Decode(rf)
		if err != nil {
			log.Fatal(err)
		}
	case "image/jpeg":
		imgSrc, _ = jpeg.Decode(rf)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("not defined")
	}
	return
}
