package matrices

import (
	"testing"
	"fmt"
)

var (
	M1 = Matrix{H: 2, W: 3, data: []float64{1, 2, 3, 4, 5, 6}}
	M2 = Matrix{H: 3, W: 2, data: []float64{7, 8, 9, 10, 11, 12}}
	M12 = Matrix{H: 2, W: 2, data: []float64{58, 64, 139, 154}}
	M4 = Matrix{H: 2, W: 2, data: []float64{2, 1, 2, 1}}
	M5 = Matrix{H: 2, W: 2, data: []float64{1, 0, 0, 1}}
	A45 = Matrix{H: 2, W: 2, data: []float64{3, 1, 2, 2}}
)

func TestNeWRandomMatrix(t *testing.T) {
	fmt.Println(NeWRandomMatrix(2, 2).data)
}

func TestMatrixMultiply(t *testing.T) {
	if got := MatrixMultiply(M1, M2) ; ! compareMatrices(got, M12) {
		printMatrix(got)
		printMatrix(M12)
		t.Errorf("Error Matrix Multiply")
	}
}

func TestMatrixAdd(t *testing.T) {
	if got := MatrixAdd(M4, M5) ; ! compareMatrices(got, A45) {
		printMatrix(got)
		printMatrix(A45)
		t.Errorf("Error matrix add")
	}
}
