package fractal

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/tomowarkar/biome"
)

type field struct {
	b biome.Biome
}

// World ...
type World interface {
	Show()
	ToPng(filename string, scale int, palette []color.Color)
	Data() []int
}

func (f *field) Show() {
	for i := 0; i < f.b.H; i++ {
		fmt.Println(f.b.Data[i*f.b.W : (i+1)*f.b.W])
	}
}

func (f *field) Set(data []int) {
	if len(data) != len(f.b.Data) {
		panic("matrix size error")
	}
	f.b.Data = data
}

func (f *field) Data() []int {
	var data []int
	for i := range f.b.Data {
		data = append(data, f.b.Data[i])
	}
	return data
}

func (f *field) ToPng(filename string, scale int, palette []color.Color) {
	img := biome.Slice2Image(biome.Biome{W: f.b.W, H: f.b.H, Data: f.b.Data}, scale, palette)
	biome.ToPng(filename, img)
}

// Flactal 中点変位法...
func Flactal(worldY, worldX int, seed int64) *field {
	w, h := (worldX/16+1)*16, (worldY/16+1)*16
	world := &field{
		biome.Biome{
			H:    h,
			W:    w,
			Data: make([]int, h*w),
		},
	}
	chunk := &field{
		biome.Biome{
			H:    18,
			W:    18,
			Data: make([]int, 18*18),
		},
	}

	for i := 0; i < w/16; i++ {
		for j := 0; j < h/16; j++ {
			//四角形の4点の高さを決定
			rand.Seed(seed + int64(i+j*10000+(i&j)*100))
			chunk.b.Data[chunk.s2m(0, 0)] = rand.Intn(255)
			rand.Seed(seed + int64(i+1+j*10000+((i+1)&j)*100))
			chunk.b.Data[chunk.s2m(16, 0)] = rand.Intn(255)
			rand.Seed(seed + int64(i+(j+1)*10000+(i&(j+1))*100))
			chunk.b.Data[chunk.s2m(0, 16)] = rand.Intn(255)
			rand.Seed(seed + int64(i+1+(j+1)*10000+((i+1)&(j+1))*100))
			chunk.b.Data[chunk.s2m(16, 16)] = rand.Intn(255)

			t1, t2, t3, t4 := chunk.b.Data[chunk.s2m(0, 0)],
				chunk.b.Data[chunk.s2m(16, 0)],
				chunk.b.Data[chunk.s2m(0, 16)],
				chunk.b.Data[chunk.s2m(16, 16)]

			mapMake(8, 8, 8, t1, t2, t3, t4, chunk)

			for i2 := 0; i2 < 16; i2++ {
				for j2 := 0; j2 < 16; j2++ {
					//生成したチャンクをワールドマップにコピペ
					world.b.Data[world.s2m(i*16+i2, j*16+j2)] = chunk.b.Data[chunk.s2m(i2, j2)]
				}
			}
		}
	}
	var data []int
	for y := 0; y < worldY; y++ {
		data = append(data, world.b.Data[y*world.b.W:y*world.b.W+worldX]...)
	}
	return &field{
		biome.Biome{
			H:    worldY,
			W:    worldX,
			Data: data,
		},
	}
}

func mapMake(x, y, size, t1, t2, t3, t4 int, f *field) {
	//再起の終了処理
	if size < 1 {
		return
	}

	//頂点の高さを決める
	mapPlus := ((t1 + t2 + t3 + t4) >> 2) + rand.Intn(size)
	if mapPlus >= 255 {
		mapPlus = 255
	}
	f.b.Data[f.s2m(x, y)] = mapPlus

	//四角形の2点同士の中点の高さを決定
	s1 := ((t1 + t2) >> 1)
	s2 := ((t1 + t3) >> 1)
	s3 := ((t2 + t4) >> 1)
	s4 := ((t3 + t4) >> 1)

	//4つの地点の座標を決める
	f.b.Data[f.s2m(x+size, y)] = s3
	f.b.Data[f.s2m(x-size, y)] = s3
	f.b.Data[f.s2m(x, y+size)] = s4
	f.b.Data[f.s2m(x, y-size)] = s1

	//分割サイズを半分にする
	size >>= 1
	mapMake(x-size, y-size, size, t1, s1, s2, mapPlus, f)
	mapMake(x+size, y-size, size, s1, t2, mapPlus, s3, f)
	mapMake(x-size, y+size, size, s2, mapPlus, t3, s4, f)
	mapMake(x+size, y+size, size, mapPlus, s3, s4, t4, f)
}

func (f *field) s2m(x, y int) int {
	return f.b.W*y + x
}
