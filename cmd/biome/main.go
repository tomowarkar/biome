package main

import (
	"time"

	"github.com/tomowarkar/biome"
)

func main() {
	d := biome.DefaultDicts()
	for i := 0; i < 30; i++ {
		world := biome.Flactal(64, 96, time.Now().UnixNano())
		world.ToPng(5, d, "./assets/tmp/image")
		time.Sleep(1000 * time.Millisecond)
	}
}
