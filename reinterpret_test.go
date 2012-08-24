package reshape

import "fmt"

func ExampleReinterpret() {
	cmplx := []complex64{complex(1, 2), complex(3, 4)}
	real := CtoR(cmplx)
	fmt.Println("CtoR(", cmplx, ") =", real)
	fmt.Println("RtoC(", real, ") =", RtoC(real))

	// Output:
	// CtoR( [(1+2i) (+3+4i)] ) = [1 2 3 4]
	// RtoC( [1 2 3 4] ) = [(1+2i) (+3+4i)]
}
