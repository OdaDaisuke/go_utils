package sliceutil

func MakeMultiDimension(x, y uint) [][]uint {
	field := make([][]uint, x)
	for i, _ := range field {
		field[i] = make([]uint, y)
	}

	return field
}
