package biome

import "strconv"

// ToStr : convert int to string
func ToStr(n int) string {
	return strconv.Itoa(n)
}

// MaxInt :
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat :
func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// MinInt :
func MinInt(a, b int) int {
	if b > a {
		return a
	}
	return b
}

// F2uint8 :
func F2uint8(m float64) uint8 {
	return uint8(uint16(m+0.5) >> 8)
}
