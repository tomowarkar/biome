package maze

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const (
	wall  = 0
	path  = 1
	route = 2
)

// Maze :
type Maze interface {
	StickDown(seed int64)
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
