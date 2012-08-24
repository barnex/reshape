package reshape

// Assuming the input has a block layout,
// return its size.
func SizeofR3(block [][][]float32) [3]int {
	return [3]int{len(block), len(block[0]), len(block[0][0])}
}

// Assuming the input has a block layout,
// return its size.
func SizeofD3(block [][][]float64) [3]int {
	return [3]int{len(block), len(block[0]), len(block[0][0])}
}

// Assuming the input has a block layout,
// return its size.
func SizeofR4(block [][][][]float32) [4]int {
	return [4]int{len(block), len(block[0]), len(block[0][0]), len(block[0][0][0])}
}

// Assuming the input has a block layout,
// return its size.
func SizeofD4(block [][][][]float64) [4]int {
	return [4]int{len(block), len(block[0]), len(block[0][0]), len(block[0][0][0])}
}
