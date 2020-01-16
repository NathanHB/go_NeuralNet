package matrices

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	h, w uint // height and width of the matrix
	data []float64
}

func printMatrix(m Matrix) {
	fmt.Printf("h: %d | w: %d\n", m.h, m.w)
	fmt.Println(m.data)
}

func compareMatrices(a, b Matrix) bool {
	// return a == b

	if a.h != b.h || a.w != b.w {
		return false
	}

	for i := uint(0) ; i < a.h * a.w ; i++ {
		if  a.data[i] != b.data[i] {
			return false
		}
	}

	return true
}

func NewRandomMatrix(h, w uint) Matrix {
	// return a h × w matrix with random values

	rand.Seed(time.Now().UnixNano()) // seed the rand to get different output each run

	data := make([]float64, h * w) // setup the data

	for i := uint(0) ; i < h * w ; i++ {
		data[i] = rand.NormFloat64() // assign random value to data[i]
	}

	c := Matrix{h, w, data}

	return c
}

func MatrixApply(m *Matrix, f func (float64) float64) {
	// take a matrix and a funciton and apply the function to all elements
	// in matrix

	for i := uint(0) ; i < m.h * m.w ; i++ {
		m.data[i] = f(m.data[i])
	}
}

func MatrixMultiply (a, b Matrix) (Matrix) {
	// return a × b

	if a.w != b.h {
		panic("Matrix Mult")
	}

	c := Matrix{h: a.h, w: b.w, data: make([]float64, a.h * b.w)}

	for i := uint(0) ; i < a.h ; i++ {
		for j := uint(0) ; j < b.w ; j++ {
			for k := uint(0) ; k < a.w ; k++ {
				c.data[i * c.w + j] += (a.data[i * a.w + k] * b.data[k * b.w + j])
			}
		}
	}

	return c
}

func MatrixAdd(a, b Matrix) (Matrix) {
	// return a + b

	if a.h != b.h || a.w != b.w {
		panic("matrix add")
	}

	c := Matrix{h: a.h, w: a.w, data: make([]float64, a.h * a.w)}

	for i := uint(0) ; i < a.h ; i++ {
		for j := uint(0) ; j < a.w ; j++ {
			c.data[i * c.w + j] = a.data[i * a.w + j] + b.data[i * b.w + j]
		}
	}

	return c
}
