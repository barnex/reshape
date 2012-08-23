/*
 Converts between 1D arrays (like []float32) and
 mulit-dimensional arrays (like [][][]float32),
 while sharing underlying storage.

 This is often needed when passing multi-dimensional
 arrays as contigous lists to, e.g., C code.
*/
package reshape
