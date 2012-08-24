package reshape

import (
	"fmt"
	"unsafe"
)

// Re-interprets the slice of complex numbers 
// as a slice of real numbers, twice as long.
// Underlying storage is shared.
func CtoR(c []complex64) []float32 {
	return (*(*[1<<31 - 1]float32)(unsafe.Pointer(&c[0])))[:2*len(c)]
}

// Re-interprets the slice of real numbers 
// as a slice of complex numbers, half as long.
// Underlying storage is shared.
func RtoC(r []float32) []complex64 {
	if len(r)%2 != 0 {
		panic(fmt.Errorf("reshape: RtoC input len should be even, got: %v", len(r)))
	}
	return (*(*[1<<31 - 1]complex64)(unsafe.Pointer(&r[0])))[:len(r)/2]
}

// Re-interprets the slice of complex numbers 
// as a slice of real numbers, twice as long.
// Underlying storage is shared.
func ZtoD(c []complex128) []float64 {
	return (*(*[1<<31 - 1]float64)(unsafe.Pointer(&c[0])))[:2*len(c)]
}

// Re-interprets the slice of real numbers 
// as a slice of complex numbers, half as long.
// Underlying storage is shared.
func DtoZ(r []float64) []complex128 {
	if len(r)%2 != 0 {
		panic(fmt.Errorf("reshape: DtoZ input len should be even, got: %v", len(r)))
	}
	return (*(*[1<<31 - 1]complex128)(unsafe.Pointer(&r[0])))[:len(r)/2]
}
