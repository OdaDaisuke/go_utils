package sliceutil

func MakeMultiDimension(x, y uint) [][]uint {
	field := make([][]uint, x)
	for i, _ := range field {
		field[i] = male([]uint, y)
	}

	return field
}
