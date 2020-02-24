package network

import (
	"fmt"
    "math/rand"
	"github.com/NathanHB/go_NeuralNet/matrices"
	"github.com/NathanHB/go_NeuralNet/actFunc"
)

type Network struct {
	biases, weights []matrices.Matrix
}

type Training_input struct {
    x []float64
    y float64
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

func costFunction(y, a matrices.Matrix) float64 {
	return 1
}

func FeedForward(inputs matrices.Matrix, net Network) matrices.Matrix {
	// takes the input matrix and returns the output after
	// putting it through the network.

	for i := range net.biases {
		inputs = matrices.MatrixMultiply(net.weights[i], inputs)
		inputs = matrices.MatrixAdd(net.biases[i], inputs)
		matrices.MatrixApply(&inputs, actFunc.Sigmoid)
	}

	return inputs
}

func update_batch(batch []Training_input, net Network, eta float64) {
    nablas_w := make([]matrices.Matrix, len(net.weights))
    nablas_b := make([]matrices.Matrix, len(net.weights))

    for i := range net.weights {
        nablas_w[i] = matrices.zeros(net.weights.h, net.weights.w)
        nablas_b[i] = matrices.zeros(net.biases.h, net.biases.w)
    }

    for i := range batch {
        delta_nablas_w, delta_nablas_b := backpropagation(batch)

        for j := range nablas_w {
            nablas_w[j] = matrices.MatrixAdd(delta_nablas_w[j], nablas_w[j])
            nablas_b[j] = matrices.MatrixAdd(delta_nablas_b[j], nablas_b[j])
        }
    }

    for i := range net.weights {
        matrices.MatrixApply(nablas_w[i],
        func(i float64) float64 { return i * eta / float64(len(batch)) })

        matrices.MatrixApply(nablas_b[i],
        func(i float64) float64 { return i * eta / float64(len(batch)) })

        net.weights[i] = matrices.MatrixSubb(net.weights[i], nablas_w[i])
        net.biases[i] = matrices.MatrixSubb(net.biases[i], nablas_b[i])
    }
}

func TrainNet(inputs [][]float64, desired_outputs []float64, epochs, batch_size int, net Network, eta float64) {
    // inputs is a list of the vectors of images to be processed and
    // desired_output is a list of corresponding output. len(inputs) == len(desired_ouptuts)
    // epochs is the number od epochs to go through and batch_size is the number
    // of elements to use for each epochs

    // creating the training_inputs array
    training_inputs := make([]Training_input, len(inputs))
    for i := range inputs {
        training_inputs[i].x = inputs[i]
        training_inputs[i].y = desired_outputs[i]
    }

    // shuffling the training array
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(training_inputs), func(i, j int) { a[i], a[j] = a[j], a[i] })

    for i := 0; i < epochs && i + batch_size < len(training_inputs); i+=batch_size {
        // update network weights and biases for this batch
        update_batch(training_inputs[i:i+batch_size], net, eta)
    }
}
