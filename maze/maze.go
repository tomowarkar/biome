package maze

import (
	"fmt"
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
}

type maze struct {
	w    int
	h    int
	data []int
}

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
