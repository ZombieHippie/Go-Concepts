package L2

import "github.com/gonum/matrix/mat64"

// A and B, are the matrices
// essentially our base case for testing
func GonumMultiply(A, B mat64.Matrix) mat64.Matrix {
	rows, _ := A.Dims()
	dense := mat64.NewDense(rows, rows, nil)
	dense.Mul(A, B)
	return dense.T()
}

// A and B, are the matrices
func StrassenMultiply(A, B mat64.Matrix) mat64.Matrix {
	rows, _ := A.Dims()
	dense := mat64.NewDense(rows, rows, nil)
	dense.Mul(A, B)
	return dense.T()
}

// M, is the matrix
func SquareMatrixMultiplyRecursive(A, B mat64.Matrix) mat64.Matrix {
	// aRows, aCols := A.Dims()
	// bRows, bCols := B.Dims()
	// C := mat64.NewDense(aRows, bCols, nil)
	// up to this point we are set up to not be a square, but if we are not a square, we will break.
	//squareMatrixMultiplyR2(A, B, 0, aRows, 0, aCols, 0, bRows, 0, bCols, C)
	result := squareMatrixMultiplyR1(A, B)
	return result.T()

}

func at(M mat64.Matrix, r0, r, c0, c int) float64 {
	return M.At(r0+r, c0+c)
}

func partition(M mat64.Matrix, r0, r1, c0, c1 int) *mat64.Dense {
	r := r1 - r0 + 1
	c := c1 - c0 + 1
	dense := mat64.NewDense(r, c, nil)

	// loop through, copying values to dense
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			v := at(M, r0, i, c0, j)
			dense.Set(i, j, v)
		}
	}

	return dense
}

func applyToDense(M mat64.Matrix, d *mat64.Dense, r0, r1, c0, c1 int) {
	r := r1 - r0 + 1
	c := c1 - c0 + 1

	// loop through, copying values to dense
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			v := M.At(i, j)
			d.Set(r0+i, c0+j, v)
		}
	}
}

func squareMatrixMultiplyR1(A, B mat64.Matrix) mat64.Matrix {
	n, _ := A.Dims() // n = A.rows
	C := mat64.NewDense(n, n, nil)
	if n == 1 {
		av := A.At(0, 0)
		bv := B.At(0, 0)
		C.Set(0, 0, av*bv) // C11 = A11 * B11
	} else {
		w, x, y, z := 0, n/2-1, n/2, n-1
		B11 := partition(B, w, x, w, x)
		B12 := partition(B, w, x, y, z)
		B21 := partition(B, y, z, w, x)
		B22 := partition(B, y, z, y, z)
		A11 := partition(A, w, x, w, x)
		A12 := partition(A, w, x, y, z)
		A21 := partition(A, y, z, w, x)
		A22 := partition(A, y, z, y, z)
		C11 := mat64.NewDense(n/2, n/2, nil)
		C12 := mat64.NewDense(n/2, n/2, nil)
		C21 := mat64.NewDense(n/2, n/2, nil)
		C22 := mat64.NewDense(n/2, n/2, nil)
		C11.Add(squareMatrixMultiplyR1(A11, B11),
			squareMatrixMultiplyR1(A12, B21))
		C12.Add(squareMatrixMultiplyR1(A11, B12),
			squareMatrixMultiplyR1(A12, B22))
		C21.Add(squareMatrixMultiplyR1(A21, B11),
			squareMatrixMultiplyR1(A22, B21))
		C22.Add(squareMatrixMultiplyR1(A21, B12),
			squareMatrixMultiplyR1(A22, B22))
		applyToDense(C11, C, w, x, w, x)
		applyToDense(C12, C, w, x, y, z)
		applyToDense(C21, C, y, z, w, x)
		applyToDense(C22, C, y, z, y, z)
	}
	return C.T()
}

func squareMatrixMultiplyR2(A, B mat64.Matrix,
	ar0, ar1, ac0, ac1,
	br0, br1, bc0, bc1 int,
	C *mat64.Dense) {
	n := ar1 - ar0 // n = A.rows
	if n == 0 {
		av := at(A, ar0, 0, ac0, 0)
		bv := at(B, br0, 0, bc0, 0)
		C.Set(ar0, bc0, av*bv) // C11 = A11 * B11
	}
}
