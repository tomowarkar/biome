package biome

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
)

type field struct {
	height int
	width  int
	field  []int
}

// World ...
type World interface {
	Show()
	RandIntn(n int, seed int64)
	ToPng(scale int, d Dicts, path string) error
}

// NewWorld ...
func NewWorld(h, w int) World {
	return &field{
		height: h,
		width:  w,
		field:  make([]int, h*w),
	}
}

func (f *field) Show() {
	for i := 0; i < f.height; i++ {
		fmt.Println(f.field[i*f.width : (i+1)*f.width])
	}
}

func (f *field) RandIntn(n int, seed int64) {
	rand.Seed(seed)
	for i := 0; i < len(f.field); i++ {
		f.field[i] = rand.Intn(n)
	}
}

func (f *field) ToPng(scale int, d Dicts, path string) error {
	img := image.NewRGBA(image.Rect(0, 0, f.width*scale, f.height*scale))
	for x := 0; x < f.width*scale; x++ {
		for y := 0; y < f.height*scale; y++ {
			if val, ok := d[f.field[y/scale*f.width+x/scale]]; ok {
				img.Set(x, y, val)
			}
		}
	}
	return encodePng(img, path)
}

func encodePng(img *image.RGBA, path string) (err error) {
	f, err := os.Create(path + ".png")
	if err != nil {
		return
	}
	defer f.Close()
	png.Encode(f, img)
	return
}
