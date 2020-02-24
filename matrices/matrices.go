package matrices

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	H, W uint // height and width of the matrix
	Data []float64
}

func PrintMatrix(m Matrix) {
	fmt.Printf("h: %d | w: %d\n", m.H, m.W)
	fmt.Println(m.Data)
}

func CompareMatrices(a, b Matrix) bool {
	// return a == b

	if a.H != b.H || a.W != b.W {
		return false
	}

	for i := uint(0) ; i < a.H * a.W ; i++ {
		if  a.Data[i] != b.Data[i] {
			return false
		}
	}

	return true
}

func NewRandomMatrix(h, w uint) Matrix {
	// return a h × w matrix with random values

	rand.Seed(time.Now().UnixNano()) // seed the rand to get different output each run

	Data := make([]float64, h * w) // setup the.Data

	for i := uint(0) ; i < h * w ; i++ {
		Data[i] = rand.NormFloat64() // assign random value to.Data[i]
	}

	c := Matrix{h, w, Data}

	return c
}

func MatrixApply(m *Matrix, f func (float64) float64) {
	// take a matrix and a funciton and apply the function to all elements
	// in matrix

	for i := uint(0) ; i < m.H * m.W ; i++ {
		m.Data[i] = f(m.Data[i])
	}
}

func MatrixMultiply (a, b Matrix) (Matrix) {
	// return a × b

	if a.W != b.H {
		panic("Matrix Mult")
	}

	c := Matrix{H: a.H, W: b.W, Data: make([]float64, a.H * b.W)}

	for i := uint(0) ; i < a.H ; i++ {
		for j := uint(0) ; j < b.W ; j++ {
			for k := uint(0) ; k < a.W ; k++ {
				c.Data[i * c.W + j] += (a.Data[i * a.W + k] * b.Data[k * b.W + j])
			}
		}
	}

	return c
}

func MatrixAdd(a, b Matrix) (Matrix) {
	// return a + b

	if a.H != b.H || a.W != b.W {
		panic("matrix add")
	}

	c := Matrix{H: a.H, W: a.W, Data: make([]float64, a.H * a.W)}

	for i := uint(0) ; i < a.H ; i++ {
		for j := uint(0) ; j < a.W ; j++ {
			c.Data[i * c.W + j] = a.Data[i * a.W + j] + b.Data[i * b.W + j]
		}
	}

	return c
}
