package biome

func (f *field) s2m(x, y int) int {
	return f.width*y + x
}

func (f *field) d8(x, y int) (u, ur, r, dr, d, dl, l, ul int) {
	coverZerosn(1, f)
	u, ur, r, dr, d, dl, l, ul = f.field[f.s2m(x+1, y)],
		f.field[f.s2m(x+2, y)],
		f.field[f.s2m(x+2, y+1)],
		f.field[f.s2m(x+2, y+2)],
		f.field[f.s2m(x+1, y+2)],
		f.field[f.s2m(x, y+2)],
		f.field[f.s2m(x, y+1)],
		f.field[f.s2m(x, y)]
	removeCovern(1, f)
	return
}

func coverZerosn(n int, f *field) {
	var arr []int
	for i := 0; i < f.height; i++ {
		arr = append(arr, make([]int, n)...)
		arr = append(arr, f.field[i*f.width:(i+1)*f.width]...)
		arr = append(arr, make([]int, n)...)
	}
	arr = append(make([]int, (n*2+f.width)*n), arr...)
	arr = append(arr, make([]int, (n*2+f.width)*n)...)

	f.height += 2 * n
	f.width += 2 * n
	f.field = arr
}

func removeCovern(n int, f *field) {
	var arr []int
	f.field = f.field[n*f.width : len(f.field)-n*f.width]
	for i := 0; i < f.height-2*n; i++ {
		arr = append(arr, f.field[i*f.width+n:(i+1)*f.width-n]...)
	}
	f.height -= 2 * n
	f.width -= 2 * n
	f.field = arr
}
