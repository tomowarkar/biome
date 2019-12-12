package biome

import (
	"math/rand"
)

// Flactal 中点変位法...
func Flactal(worldY, worldX int, seed int64) (world *field) {
	world = &field{
		height: worldY,
		width:  worldX,
		field:  make([]int, worldX*worldY),
	}
	chunk := &field{
		height: 18,
		width:  18,
		field:  make([]int, 18*18),
	}

	for i := 0; i < worldX/16; i++ {
		for j := 0; j < worldY/16; j++ {
			//四角形の4点の高さを決定
			rand.Seed(seed + int64(i+j*10000+(i&j)*100))
			chunk.field[chunk.s2m(0, 0)] = rand.Intn(255)
			rand.Seed(seed + int64(i+1+j*10000+((i+1)&j)*100))
			chunk.field[chunk.s2m(16, 0)] = rand.Intn(255)
			rand.Seed(seed + int64(i+(j+1)*10000+(i&(j+1))*100))
			chunk.field[chunk.s2m(0, 16)] = rand.Intn(255)
			rand.Seed(seed + int64(i+1+(j+1)*10000+((i+1)&(j+1))*100))
			chunk.field[chunk.s2m(16, 16)] = rand.Intn(255)

			t1, t2, t3, t4 := chunk.field[chunk.s2m(0, 0)],
				chunk.field[chunk.s2m(16, 0)],
				chunk.field[chunk.s2m(0, 16)],
				chunk.field[chunk.s2m(16, 16)]

			mapMake(8, 8, 8, t1, t2, t3, t4, chunk)

			for i2 := 0; i2 < 16; i2++ {
				for j2 := 0; j2 < 16; j2++ {
					//生成したチャンクをワールドマップにコピペ
					world.field[world.s2m(i*16+i2, j*16+j2)] = chunk.field[chunk.s2m(i2, j2)]
				}
			}
		}
	}
	return
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
	f.field[f.s2m(x, y)] = mapPlus

	//四角形の2点同士の中点の高さを決定
	s1 := ((t1 + t2) >> 1)
	s2 := ((t1 + t3) >> 1)
	s3 := ((t2 + t4) >> 1)
	s4 := ((t3 + t4) >> 1)

	//4つの地点の座標を決める
	f.field[f.s2m(x+size, y)] = s3
	f.field[f.s2m(x-size, y)] = s3
	f.field[f.s2m(x, y+size)] = s4
	f.field[f.s2m(x, y-size)] = s1

	//分割サイズを半分にする
	size >>= 1
	mapMake(x-size, y-size, size, t1, s1, s2, mapPlus, f)
	mapMake(x+size, y-size, size, s1, t2, mapPlus, s3, f)
	mapMake(x-size, y+size, size, s2, mapPlus, t3, s4, f)
	mapMake(x+size, y+size, size, mapPlus, s3, s4, t4, f)
}
