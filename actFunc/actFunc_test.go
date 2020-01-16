package actFunc

import (
	"testing"
	"fmt"
)

func TestSigmoid(t *testing.T) {
	fmt.Println(Sigmoid(float64(100)))
	fmt.Println(Sigmoid(float64(0)))
	fmt.Println(Sigmoid(float64(-100)))
}
