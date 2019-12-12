package biome

import (
	"image/color"
	"testing"
)

func TestColors(t *testing.T) {
}
func TestPickColor(t *testing.T) {
	tests := []struct {
		name  string
		color string
		want  color.RGBA
	}{
		{"Pass", "khaki", color.RGBA{240, 230, 140, 255}},
		{"darkgoldenrod", "darkgoldenrod", color.RGBA{184, 134, 11, 255}},
		{"royalblue", "royalblue", color.RGBA{65, 105, 225, 255}},
		{"seagreen", "seagreen", color.RGBA{46, 139, 87, 255}},
		{"not registered", "seaeen", color.RGBA{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PickColor(tt.color)
			if got != tt.want {
				t.Fatalf("want %d but %d", tt.want, got)
			}
		})
	}
}
