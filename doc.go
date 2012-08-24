/*
 This package converts between 1D arrays (like []float32) and
 mulit-dimensional arrays (like [][][]float32),
 while sharing underlying storage.

 Also allows to re-interpret arrays.
 E.g., cast a []complex128 as a []float64 of twice the size. 

 This is often needed when passing multi-dimensional
 arrays as contigous lists to, e.g., C code.
*/
package reshape
