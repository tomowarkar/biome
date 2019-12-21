package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/tomowarkar/biome"
)

const (
	wall  = 0
	path  = 1
	route = 2
)

// Maze :
type Maze interface {
	StickDown(seed int64)
	Digging(seed int64)
	Solve()
	ToPng(scale int, dirPath string) error
	Data() []int
	Set(data []int)
}

type maze struct {
	w    int
	h    int
	data []int
}

type p struct{ x, y int }

var (
	up    = p{0, -1}
	down  = p{0, 1}
	right = p{1, 0}
	left  = p{-1, 0}
)

// NewMaze :
func NewMaze(row, column int) Maze {
	if row < 7 {
		row = 7
	}
	if column < 7 {
		column = 7
	}
	w, h := (column/2)*2+1, (row/2)*2+1
	return &maze{
		w:    w,
		h:    h,
		data: make([]int, w*h),
	}
}

func (m maze) plt() {
	var text []string
	var strReplacer *strings.Replacer = strings.NewReplacer(
		"0", "#",
		"1", " ",
		"2", ".",
	)
	for _, v := range m.data {
		text = append(text, strReplacer.Replace(strconv.Itoa(v)))
	}

	for y := 0; y < m.h; y++ {
		fmt.Println(strings.Join(text[y*m.w:(y+1)*m.w], " "))
	}
	fmt.Println()
}

func (m maze) at(x, y int) int {
	return m.data[y*m.w+x]
}

func (m *maze) setObj(x, y, obj int) {
	m.data[y*m.w+x] = obj
}

func (m *maze) StickDown(seed int64) {
	rand.Seed(seed)
	// 初期データ挿入
	for y := 0; y < m.h; y++ {
		for x := 0; x < m.w; x++ {
			if x == 0 || x == m.w-1 || y == 0 || y == m.h-1 || (x%2 == 0 && y%2 == 0) {
				m.setObj(x, y, wall)
			} else {
				m.setObj(x, y, path)
			}
		}
	}
	// 迷路生成
	for y := 2; y < m.h-2; y += 2 {
		for x := 2; x < m.w-2; x += 2 {
			if y == 2 {
				dirs := []p{up, down, right, left}
				n := rand.Intn(len(dirs))
				m.setObj((x + dirs[n].x), (y + dirs[n].y), wall)
			} else {
				dirs := []p{down, right, left}
				for len(dirs) > 0 {
					n := rand.Intn(len(dirs))
					dir := dirs[n]
					dirs = append(dirs[:n], dirs[n+1:]...)
					if m.at(x+dir.x, y+dir.y) == path {
						m.setObj((x + dir.x), (y + dir.y), wall)
						break
					}
				}
			}
		}
	}
	m.plt()
}

func (m *maze) Solve() {
	start, goal := p{0, 1}, p{m.w - 2, m.h - 2}
	visited := make([]bool, len(m.data))
	costs := make([]int, len(m.data))
	for i := 0; i < len(costs); i++ {
		costs[i] = len(m.data)
	}
	queue := []p{start}
	costs[start.y*m.w+start.x] = 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		visited[pos.y*m.w+pos.x] = true

		for _, v := range []p{up, down, right, left} {
			x, y := pos.x+v.x, pos.y+v.y
			currP := p{x, y}
			if x < 0 || y < 0 || x > m.w-1 || y > m.h-1 {
				continue
			}
			if goal == currP {
				queue = make([]p, 0)
				costs[y*m.w+x] = costs[pos.y*m.w+pos.x] + 1
				break
			}
			if m.at(x, y) == path && visited[y*m.w+x] == false {
				queue = append(queue, currP)
				costs[y*m.w+x] = costs[pos.y*m.w+pos.x] + 1
			}
		}
	}

	point := goal
	cost := costs[goal.y*m.w+goal.x]
	if cost == len(m.data) {
		return
	}
	m.data[goal.y*m.w+goal.x] = route
	for cost > 1 {
		for _, v := range []p{up, down, right, left} {
			x, y := point.x+v.x, point.y+v.y
			if x < 0 || y < 0 || x > m.w-1 || y > m.h-1 {
				continue
			}
			if costs[y*m.w+x] == cost-1 {
				cost = costs[y*m.w+x]
				point = p{x, y}
				m.data[y*m.w+x] = route
				break
			}
		}
	}
	m.plt()
}

func (m *maze) Digging(seed int64) {
	rand.Seed(seed)
	for y := 0; y < m.h; y++ {
		for x := 0; x < m.w; x++ {
			if x == 0 || y == 0 || x == m.w-1 || y == m.h-1 {
				m.setObj(x, y, path)
			} else {
				m.setObj(x, y, wall)
			}
		}
	}
	m.dig(1, 1)
	for y := 0; y < m.h; y++ {
		for x := 0; x < m.w; x++ {
			if x == 0 || y == 0 || x == m.w-1 || y == m.h-1 {
				m.setObj(x, y, wall)
			}
		}
	}
	m.plt()
}

func (m *maze) dig(x, y int) {
	cands := []p{{x, y}}
	for len(cands) > 0 {
		n := rand.Intn(len(cands))
		cand := cands[n]
		for {
			var dirs []p
			var dirs2 []p
			if m.at(cand.x, cand.y-1) == wall && m.at(cand.x, cand.y-2) == wall {
				dirs = append(dirs, up)
				dirs2 = append(dirs2, p{0, -2})
			}
			if m.at(cand.x+1, cand.y) == wall && m.at(cand.x+2, cand.y) == wall {
				dirs = append(dirs, right)
				dirs2 = append(dirs2, p{2, 0})
			}
			if m.at(cand.x, cand.y+1) == wall && m.at(cand.x, cand.y+2) == wall {
				dirs = append(dirs, down)
				dirs2 = append(dirs2, p{0, 2})
			}
			if m.at(cand.x-1, cand.y) == wall && m.at(cand.x-2, cand.y) == wall {
				dirs = append(dirs, left)
				dirs2 = append(dirs2, p{-2, 0})
			}
			if len(dirs) == 0 {
				for i, v := range cands {
					if v == cand {
						cands = append(cands[:i], cands[i+1:]...)
						break
					}
				}
				break
			}
			m.setObj(cand.x, cand.y, path)
			nn := rand.Intn(len(dirs))
			m.setObj(cand.x+dirs[nn].x, cand.y+dirs[nn].y, path)
			m.setObj(cand.x+dirs2[nn].x, cand.y+dirs2[nn].y, path)
			cands = append(cands, p{cand.x + dirs2[nn].x, cand.y + dirs2[nn].y})
			cand = p{cand.x + dirs2[nn].x, cand.y + dirs2[nn].y}
		}
	}
}

func (m *maze) ToPng(scale int, dirPath string) error {
	d := biome.NewDicts()
	d.Set(wall, biome.Colors["darkgoldenrod"])
	d.Set(path, biome.Colors["khaki"])
	d.Set(route, biome.Colors["magenta"])

	img := image.NewRGBA(image.Rect(0, 0, m.w*scale, m.h*scale))
	for x := 0; x < m.w*scale; x++ {
		for y := 0; y < m.h*scale; y++ {
			if val, ok := d[m.data[y/scale*m.w+x/scale]]; ok {
				img.Set(x, y, val)
			}
		}
	}
	return encodePng(img, dirPath)
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

func (m *maze) Set(data []int) {
	if len(data) != len(m.data) {
		panic("err")
	}
	m.data = data
}

func (m maze) Data() []int {
	return m.copy()
}

func (m *maze) copy() []int {
	copy := make([]int, len(m.data))
	for i := range copy {
		copy[i] = m.data[i]
	}
	return copy
}

// ToGif :
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
