package reshape

// Reshape the block to one contiguous list.
// Assumes the block's storage is contiguous.
func ContiguousR3(block [][][]float32) []float32 {
	N := prod3(SizeofR3(block))
	return block[0][0][:N]
}

// Reshape the block to one contiguous list.
// Assumes the block's storage is contiguous.
func ContiguousD3(block [][][]float64) []float64 {
	N := prod3(SizeofD3(block))
	return block[0][0][:N]
}

// Reshape the block to one contiguous list.
// Assumes the block's storage is contiguous.
func ContiguousR4(block [][][][]float32) []float32 {
	N := prod4(SizeofR4(block))
	return block[0][0][0][:N]
}

// Reshape the block to one contiguous list.
// Assumes the block's storage is contiguous.
func ContiguousD4(block [][][][]float64) []float64 {
	N := prod4(SizeofD4(block))
	return block[0][0][0][:N]
}
