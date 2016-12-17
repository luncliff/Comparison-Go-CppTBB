// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//  File 	: matrix.go
//  Author 	: Park  Dong Ha ( luncliff@gmail.com )
//  Updated : 2016/12/17
//
//  Note 	:
//		Simple matrix with go slice
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package matrix

// Int2D ...
// 		Allocate large `int` array
//      - Note
//          overflow handling required
func Int2D(row int, col int) (mat [][]int) {
	size := row * col
	arr := make([]int, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}

// Float642D ...
// 		Allocate large `float64 array
func Float642D(row int, col int) (mat [][]float64) {
	size := row * col
	arr := make([]float64, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}

// ChanInt2D ...
// 		Allocate large `chan int`
func ChanInt2D(row int, col int) (mat [][]chan int) {
	size := row * col
	arr := make([]chan int, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}
