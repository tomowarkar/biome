package maze

import (
	"image/color"
	"testing"
	"time"
)

func TestMaze(t *testing.T) {

}

func ExampleMaze() {
	// mazeWidth and mazeHeight should odd number.
	mazeWidth, mazeHeight := 37, 27
	m := NewMaze(mazeHeight, mazeWidth)
	seed := time.Now().UnixNano()

	// GIF画像生成用素材の格納
	var debug [][]int

	// 棒倒し法で迷路を生成
	m.StickDown(seed)
	debug = append(debug, m.Data())
	// 迷路の最短経路となる解を導出
	m.Solve()
	debug = append(debug, m.Data())

	// 穴掘り法で迷路を生成
	m.Digging(seed)
	debug = append(debug, m.Data())
	// 迷路の最短経路となる解を導出
	m.Solve()
	debug = append(debug, m.Data())

	// 画像マッピング用の色を格納
	palette := []color.Color{
		color.RGBA{184, 134, 11, 255},
		color.RGBA{240, 230, 140, 255},
		color.RGBA{255, 0, 255, 255},
	}
	// 10倍スケール(270×370)のpng画像を生成(maze.png)
	m.ToPng("maze", 10, palette)
	// 10倍スケール(270×370)のGIF画像を生成(mazegif.gif)
	ToGif("mazegif", mazeWidth, mazeHeight, 10, 100, debug, palette)
}
