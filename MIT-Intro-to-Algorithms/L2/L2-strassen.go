package L2

import "github.com/gonum/matrix/mat64"

// A and B, are the matrices
func StrassenMultiply(A, B mat64.Matrix) mat64.Matrix {
	rows, _ := A.Dims()
	dense := mat64.NewDense(rows, rows, nil)
	dense.Mul(A, B)
	return dense.T()
}

// A and B, are the matrices
func GonumMultiply(A, B mat64.Matrix) mat64.Matrix {
	rows, _ := A.Dims()
	dense := mat64.NewDense(rows, rows, nil)
	dense.Mul(A, B)
	return dense.T()
}

// M, is the matrix
func SquareMatrixMultiplyRecursive(A, B mat64.Matrix) mat64.Matrix {
	return A
}
