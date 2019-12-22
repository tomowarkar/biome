package biome

import (
	"image/color"
)

// Dicts ...
type Dicts map[int]color.RGBA

// NewDicts ...
func NewDicts() Dicts {
	return map[int]color.RGBA{}
}

// DefaultDicts ...
func DefaultDicts() Dicts {
	d := NewDicts()
	for i := 0; i < 256; i++ {
		if 240 <= i {
			d.Set(i, White)
		} else if 180 <= i {
			d.Set(i, Seagreen)
		} else if 150 <= i {
			d.Set(i, Darkgoldenrod)
		} else if 130 <= i {
			d.Set(i, Khaki)
		} else {
			d.Set(i, Royalblue)
		}
	}
	return d
}

// Set ...
func (ds Dicts) Set(number int, color color.RGBA) {
	ds[number] = color
}

var (
	// White ...
	White = color.RGBA{255, 255, 255, 255}
	// Khaki ...
	Khaki = color.RGBA{240, 230, 140, 255}
	// Darkgoldenrod ...
	Darkgoldenrod = color.RGBA{184, 134, 11, 255}
	// Royalblue ...
	Royalblue = color.RGBA{65, 105, 225, 255}
	// Seagreen ...
	Seagreen = color.RGBA{46, 139, 87, 255}
	// Cyan ...
	Cyan = color.RGBA{0, 255, 255, 255}
	// Magenta ...
	Magenta = color.RGBA{255, 0, 255, 255}
	// Yellow ...
	Yellow = color.RGBA{255, 255, 0, 255}
)
