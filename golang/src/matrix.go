// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author 	: Park  Dong Ha ( luncliff@gmail.com )
//
//  Note 	:
//		Simple matrix with go slice
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package main

type resType int

// template ...
//  	!!! Deprecated function !!!
// 		Matrix implementation with nested slice
func template(row int, col int) (mat [][]resType) {
	size := row * col
	arr := make([]resType, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}

// Int2D ...
// 		Allocate large `int`` array
type Int2D struct {
	row    uint
	column uint
	chunk  []int
}

func MakeInt2D(row int, col int) (mat Int2D) {
	mat.column = uint(col)
	mat.row = uint(row)
	mat.chunk = make([]int, mat.Size())
	return mat
}

func (mat *Int2D) Size() uint {
	return mat.row * mat.column
}

func (mat *Int2D) At(row int, col int) *int {
	index := row*int(mat.column) + col
	return &mat.chunk[index]
}

// Float2D ...
// 		Allocate large `float64`` array
type Float2D struct {
	row    uint
	column uint
	chunk  []float64
}

func MakeFloat2D(row int, col int) (mat Float2D) {
	mat.column = uint(col)
	mat.row = uint(row)
	mat.chunk = make([]float64, mat.Size())
	return mat
}

func (mat *Float2D) Size() uint {
	return mat.row * mat.column
}

func (mat *Float2D) At(row int, col int) *float64 {
	index := row*int(mat.column) + col
	return &mat.chunk[index]
}
