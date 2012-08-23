package reshape

import "fmt"

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func R3(array []float32, size [3]int) [][][]float32 {
	Nx, Ny, Nz := size[0], size[1], size[2]
	if Nx*Ny*Nz != len(array) {
		panic(fmt.Errorf("reshape: size mismatch: %v*%v*%v != %v", Nx, Ny, Nz, len(array)))
	}
	sliced := make([][][]float32, Nx)
	for i := range sliced {
		sliced[i] = make([][]float32, Ny)
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = array[(i*Ny+j)*Nz+0 : (i*Ny+j)*Nz+Nz]
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func R4(list []float32, size [4]int) [][][][]float32 {
	sliced := make([][][][]float32, size[0])
	for i := range sliced {
		sliced[i] = make([][][]float32, size[1])
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = make([][]float32, size[2])
		}
	}

	for i := range sliced {
		for j := range sliced[i] {
			for k := range sliced[i][j] {
				sliced[i][j][k] = list[((i*size[1]+j)*size[2]+k)*size[3]+0 : ((i*size[1]+j)*size[2]+k)*size[3]+size[3]]
			}
		}
	}
	return sliced
}

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

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func D3(array []float64, size [3]int) [][][]float64 {
	Nx, Ny, Nz := size[0], size[1], size[2]
	if Nx*Ny*Nz != len(array) {
		panic(fmt.Errorf("reshape: size mismatch: %v*%v*%v != %v", Nx, Ny, Nz, len(array)))
	}
	sliced := make([][][]float64, Nx)
	for i := range sliced {
		sliced[i] = make([][]float64, Ny)
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = array[(i*Ny+j)*Nz+0 : (i*Ny+j)*Nz+Nz]
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func C3(array []complex64, size [3]int) [][][]complex64 {
	Nx, Ny, Nz := size[0], size[1], size[2]
	if Nx*Ny*Nz != len(array) {
		panic(fmt.Errorf("reshape: size mismatch: %v*%v*%v != %v", Nx, Ny, Nz, len(array)))
	}
	sliced := make([][][]complex64, Nx)
	for i := range sliced {
		sliced[i] = make([][]complex64, Ny)
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = array[(i*Ny+j)*Nz+0 : (i*Ny+j)*Nz+Nz]
		}
	}
	return sliced
}

// Re-interpret a contiguous array as a multi-dimensional array of given size.
// Underlying storage is shared.
func Z3(array []complex128, size [3]int) [][][]complex128 {
	Nx, Ny, Nz := size[0], size[1], size[2]
	if Nx*Ny*Nz != len(array) {
		panic(fmt.Errorf("reshape: size mismatch: %v*%v*%v != %v", Nx, Ny, Nz, len(array)))
	}
	sliced := make([][][]complex128, Nx)
	for i := range sliced {
		sliced[i] = make([][]complex128, Ny)
	}
	for i := range sliced {
		for j := range sliced[i] {
			sliced[i][j] = array[(i*Ny+j)*Nz+0 : (i*Ny+j)*Nz+Nz]
		}
	}
	return sliced
}

func prod3(size [3]int) int {
	return size[0] * size[1] * size[2]
}

func prod4(size [4]int) int {
	return size[0] * size[1] * size[2] * size[3]
}
