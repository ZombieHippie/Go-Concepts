package L2_test

import (
	"github.com/ZombieHippie/Go-Concepts/MIT-Intro-to-Algorithms/L2"
	"github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
	"testing"
)

func createDense(size int, fn func(int, int) float64) mat64.Matrix {
	rows := int(math.Pow(2.0, float64(size)))
	// Generate a 8×8 matrix of random values.
	data := make([]float64, rows*rows)
	for i := range data {
		data[i] = fn(i%size, int(i-i%size)/size)
	}
	a := mat64.NewDense(rows, rows, data)
	return a
}

func randomFn(i, j int) float64 {
	return math.Floor(rand.NormFloat64() * 1e2)
}

func TestStrassenMultiply(t *testing.T) {
	size := 2

	A := createDense(size, randomFn)
	B := createDense(size, randomFn)
	t.Log("A", A)
	t.Log("B", B)

	matResultStrassen := L2.StrassenMultiply(A, B)
	matResultGonum := L2.GonumMultiply(A, B)

	t.Log("Gonum result", mat64.DenseCopyOf(matResultGonum))
	t.Log("Strassen result", mat64.DenseCopyOf(matResultStrassen))

	if !mat64.Equal(matResultGonum, matResultStrassen) {
		t.Fatal("Results were not equivalent")
		t.FailNow()
	}
}

func TestSquareMatrixMultiply(t *testing.T) {
	size := 2

	A := createDense(size, randomFn)
	B := createDense(size, randomFn)
	t.Log("A", A)
	t.Log("B", B)

	matResultStrassen := L2.SquareMatrixMultiplyRecursive(A, B)
	matResultGonum := L2.GonumMultiply(A, B)

	t.Log("Gonum result", mat64.DenseCopyOf(matResultGonum))
	t.Log("Square Matrix result", mat64.DenseCopyOf(matResultStrassen))

	if !mat64.Equal(matResultGonum, matResultStrassen) {
		t.Fatal("Results were not equivalent")
		t.FailNow()
	}
}
