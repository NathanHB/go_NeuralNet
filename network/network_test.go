package network

import (
	"testing"
    "fmt"
    "github.com/NathanHB/go_NeuralNet/matrices"
	p "github.com/NathanHB/go_NeuralNet/parser"
)

var net = InitNewNetwork([]uint{2, 3, 2})

func TestInitNewNetwork(t *testing.T) {
	//PrintNetwork(net)
}

func TestFeedForward(t *testing.T) {
	//input := matrices.Matrix{H: 2, W: 1, Data: []float64{1, 0}}
	//matrices.PrintMatrix(FeedForward(input, net))
}

func TestTrainNet(t *testing.T) {
    trainImagesPath := "../parser/train-images-idx3-ubyte"
    trainLabelsPath := "../parser/train-labels-idx1-ubyte"
    images, labels := p.MakeInputArray(trainImagesPath, trainLabelsPath)

    fmt.Println(net)
    input := matrices.Matrix{H: 2, W: 1, Data: []float64{1, 0}}
    matrices.PrintMatrix(FeedForward(input, net))
    TrainNet(images, labels, 10, 10, net, 0.5)
    fmt.Println(net)
    matrices.PrintMatrix(FeedForward(input, net))
}
