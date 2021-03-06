package L2_test

import (
	"github.com/ZombieHippie/Go-Concepts/MIT-Intro-to-Algorithms/L2"
	"github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func createDense(size int, fn func(int, int) float64) mat64.Matrix {
	rows := int(math.Pow(2.0, float64(size)))
	// Generate a 8×8 matrix of random values.
	data := make([]float64, rows*rows)
	for i := range data {
		data[i] = fn(i%rows, int(i-i%rows)/rows)
	}
	a := mat64.NewDense(rows, rows, data)
	return a
}

func randomFn(i, j int) float64 {
	return math.Floor(rand.NormFloat64() * 1e2)
}

func testStrassenMultiply(t *testing.T, size int) {
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
func TestApplyToDense(t *testing.T) {
	test := mat64.NewDense(4, 4, nil)
	applying := createDense(2, randomFn)
	t.Log("applying", applying)

	L2.ApplyToDense(applying, test, 2, 3, 2, 3)

	t.Log("Applied result", test)
}
func TestPartition(t *testing.T) {
	parting := createDense(2, randomFn)
	t.Logf("parting:\n %v", parting)
	n := 4
	w, x, y, z := 0, n/2-1, n/2, n-1

	// partition 12
	dense := L2.Partition(parting, w, x, y, z)

	t.Logf("Partition(parting, %d, %d, %d, %d):\n %v", w, x, y, z, dense)
	for i := 0; i < y; i++ {
		for j := 0; j < y; j++ {
			t.Logf("Parted result At(%d,%d): %f", i, j, dense.At(i, j))
		}
	}
}

func TestStrassenMultiply2(t *testing.T) {
	testStrassenMultiply(t, 2)
}
func TestStrassenMultiply1(t *testing.T) {
	testStrassenMultiply(t, 1)
}
func testSquareMatrixMultiply(t *testing.T, size int) {
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
func TestSquareMatrixMultiply2(t *testing.T) {
	testSquareMatrixMultiply(t, 2)
}
func TestSquareMatrixMultiply1(t *testing.T) {
	testSquareMatrixMultiply(t, 1)
}
func TestSquareMatrixMultiply0(t *testing.T) {
	testSquareMatrixMultiply(t, 0)
}
