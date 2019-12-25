package paint

import (
	"image/color"
	"math"

	"github.com/tomowarkar/biome"
)

// Canvas :
type Canvas interface {
	Fill(x, y int, obj int)
	Line(x1, y1, x2, y2 int, px float64, obj int)
	Square(x, y int, px, deg float64, obj int)
	Triangle(x, y int, px, deg float64, obj int)
	Dot(x, y int, px float64, obj int)
	Data() []int
	ToPng(filename string, palette []color.Color)
}

type canvas struct {
	height int
	width  int
	data   []int
}

type point struct {
	x, y, r float64
}

type line struct {
	begin, end point
}

// NewCanv :
func NewCanv(width, height int) Canvas {
	return &canvas{
		height: height,
		width:  width,
		data:   make([]int, height*width),
	}
}

func (c *canvas) Dot(x, y int, px float64, obj int) {
	dot := point{float64(x), float64(y), px}
	c.dot(dot, obj)
}

func (c *canvas) Triangle(x, y int, px, deg float64, obj int) {
	dot1 := point{
		float64(x) + px*math.Cos((90+deg)/180*math.Pi),
		float64(y) - px*math.Sin((90+deg)/180*math.Pi),
		1,
	}
	dot2 := point{
		float64(x) + px*math.Cos((210+deg)/180*math.Pi),
		float64(y) - px*math.Sin((210+deg)/180*math.Pi),
		1,
	}
	dot3 := point{
		float64(x) + px*math.Cos((330+deg)/180*math.Pi),
		float64(y) - px*math.Sin((330+deg)/180*math.Pi),
		1,
	}

	c.line(line{dot1, dot2}, obj)
	c.line(line{dot2, dot3}, obj)
	c.line(line{dot3, dot1}, obj)
}

func (c *canvas) Square(x, y int, px, deg float64, obj int) {
	arm := px * math.Sqrt2
	dot1 := point{
		float64(x) + arm*math.Cos((45+deg)/180*math.Pi),
		float64(y) - arm*math.Sin((45+deg)/180*math.Pi),
		1,
	}
	dot2 := point{
		float64(x) + arm*math.Cos((135+deg)/180*math.Pi),
		float64(y) - arm*math.Sin((135+deg)/180*math.Pi),
		1,
	}
	dot3 := point{
		float64(x) + arm*math.Cos((225+deg)/180*math.Pi),
		float64(y) - arm*math.Sin((225+deg)/180*math.Pi),
		1,
	}
	dot4 := point{
		float64(x) + arm*math.Cos((315+deg)/180*math.Pi),
		float64(y) - arm*math.Sin((315+deg)/180*math.Pi),
		1,
	}
	c.line(line{dot1, dot2}, obj)
	c.line(line{dot2, dot3}, obj)
	c.line(line{dot3, dot4}, obj)
	c.line(line{dot4, dot1}, obj)
}

func (c *canvas) Data() []int {
	var data []int
	for i := range c.data {
		data = append(data, c.data[i])
	}
	return data
}

func (c *canvas) Line(x1, y1, x2, y2 int, px float64, obj int) {
	line := line{
		point{float64(x1), float64(y1), px},
		point{float64(x2), float64(y2), px},
	}
	c.line(line, obj)
}

func (c *canvas) Fill(x, y int, obj int) {
	c.fill(x, y, obj)
}

func (c *canvas) ToPng(filename string, palette []color.Color) {
	img := biome.Slice2Image(c.width, c.height, 1, c.data, palette)
	biome.ToPng(filename, img)
}

func (c canvas) arc(x, y, r, sAngle, eAngle float64) {
	// TODO 曲線
}

func (c canvas) fill(x, y, obj int) {
	target := c.data[y*c.width+x]

	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}

	var now [2]int
	queue := [][2]int{[2]int{x, y}}
	for len(queue) > 0 {
		now, queue = queue[0], queue[1:]
		for i := 0; i < 4; i++ {
			xx, yy := now[0]+dx[i], now[1]+dy[i]
			if xx < 0 || yy < 0 || xx > c.width-1 || yy > c.height-1 {
				continue
			}
			if c.data[yy*c.width+xx] == target {
				queue = append(queue, [2]int{xx, yy})
				c.data[yy*c.width+xx] = obj
			}
		}
	}

}

func distL(x, y int, l line) bool {
	var xx, yy = float64(x), float64(y)
	var mx, my = (l.begin.x + l.end.x) / 2, (l.begin.y + l.end.y) / 2
	var dx, dy = l.end.x - l.begin.x, l.end.y - l.begin.y
	var b1, b2 = -dy / dx, dx / dy
	var c1, c2 = -(b1*mx + my), -(b2*mx + my)
	r1 := biome.MaxFloat(l.begin.r, l.end.r)
	r2 := math.Sqrt(dx*dx+dy*dy) / 2
	d1 := math.Abs(yy+b1*xx+c1) / math.Sqrt(1+b1*b1)
	d2 := math.Abs(yy+b2*xx+c2) / math.Sqrt(1+b2*b2)

	if dx == 0 && dy == 0 {
		return distP(x, y, l.begin)
	} else if dx == 0 {
		d1 = math.Abs(xx - mx)
		d2 = math.Abs(yy - my)
	} else if dy == 0 {
		d1 = math.Abs(yy - my)
		d2 = math.Abs(xx - mx)
	}

	if d1 < r1 && d2 < r2 {
		return true
	}
	return false
}

// TODO 全面参照ではなく、範囲を絞って描画
func (c canvas) line(l line, obj int) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if distL(x, y, l) {
				c.data[y*c.width+x] = obj
			}
		}
	}
}

func distP(x, y int, p point) bool {
	var dx, dy = p.x - float64(x), p.y - float64(y)
	d := math.Sqrt(dx*dx+dy*dy) / p.r
	if d < 1 {
		return true
	}
	return false
}

// TODO 全面参照ではなく、範囲を絞って描画
func (c canvas) dot(p point, obj int) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			if distP(x, y, p) {
				c.data[y*c.width+x] = obj
			}
		}
	}
}
