package actFunc

import (
	"math"
)

func Sigmoid(x float64)float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidPrime(x float64) float64 {
	// derivative of the sigmoid function
	return Sigmoid(x) * (1 - Sigmoid(x))

