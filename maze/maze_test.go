package maze

import (
	"testing"
	"time"
)

func TestMaze(t *testing.T) {
	m := NewMaze(7, 9)
	m.StickDown(time.Now().UnixNano())
	m.Solve()
	m.Digging(time.Now().UnixNano())
	m.Solve()
}

// func TestMaze2(t *testing.T) {
// 	m := NewMaze(15, 24)
// 	m.StickDown(time.Now().UnixNano())
// 	m.Solve()
// 	m.Digging(time.Now().UnixNano())
// 	m.Solve()
// }
