package network

import (
	"fmt"
	"github.com/NathanHB/go_NeuralNet/matrices"
)

type Network struct {
	biases, weights []matrices.Matrix
}

func PrintNetwork(net Network) {
	fmt.Println("\nbiases:")
	for i := range net.biases {
		matrices.PrintMatrix(net.biases[i])
	}

	fmt.Println("weights:")
	for i := range net.weights {
		matrices.PrintMatrix(net.weights[i])
	}
	fmt.Println()
}

func InitNewNetwork(networkSig []uint) Network {
	// networkSig is an array containing the size of each layer starting
	// with input layer. len(networkSig) is the number of layer of the
	// network
	// return a net filled with random values

	biases := make([]matrices.Matrix, len(networkSig) - 1)

	for i := range biases {
		biases[i] = matrices.NewRandomMatrix(networkSig[i + 1], 1)
	}

	weights := make([]matrices.Matrix, len(networkSig) - 1)

	for i := range weights {
		weights[i] = matrices.NewRandomMatrix(networkSig[i + 1], networkSig[i])
	}

	net := Network{biases: biases, weights: weights}

	return net
}

func FeedForward(inputs matrices.Matrix) []float64 {

}
