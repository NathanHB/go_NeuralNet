package network

import (
	"testing"
	"github.com/NathanHB/go_NeuralNet/matrices"
)

var net = InitNewNetwork([]uint{2, 9, 120, 50, 10})

func TestInitNewNetwork(t *testing.T) {
	//PrintNetwork(net)
}

func TestFeedForward(t *testing.T) {
	input := matrices.Matrix{H: 2, W: 1, Data: []float64{1, 0}}
	matrices.PrintMatrix(FeedForward(input, net))
}

func TestTrainNet(t *testing.T) {
}

