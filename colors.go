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
			d.Set(i, color.RGBA{255, 255, 255, 255})
		} else if 180 <= i {
			d.Set(i, PickColor("seagreen"))
		} else if 150 <= i {
			d.Set(i, PickColor("darkgoldenrod"))
		} else if 130 <= i {
			d.Set(i, PickColor("khaki"))
		} else {
			d.Set(i, PickColor("royalblue"))
		}
	}
	return d
}

// Set ...
func (ds Dicts) Set(number int, color color.RGBA) {
	ds[number] = color
}

// Colors ...
var Colors = map[string]color.RGBA{
	"khaki":         color.RGBA{240, 230, 140, 255},
	"darkgoldenrod": color.RGBA{184, 134, 11, 255},
	"royalblue":     color.RGBA{65, 105, 225, 255},
	"seagreen":      color.RGBA{46, 139, 87, 255},
}

// PickColor ...
func PickColor(color string) color.RGBA {
	return Colors[color]
}
