package parser

import (
    "testing"
    "fmt"
)

func printImages(images [][]float64, labels []float64, k int) {

    for n := 0; n < k; n++ {
        for i := 0; i < 28; i++ {
            for j := 0; j < 28; j++ {
                if images[n][i * 28 + j] != 0 {
                    fmt.Printf("0 ")
                } else {
                    fmt.Printf(". ")
                }
            }
            fmt.Println()
        }
        fmt.Printf("value: %f\n", labels[n])
        fmt.Println()
    }
}

func TestMakeInputArray(t *testing.T) {
    testImagesPath := "t10k-images-idx3-ubyte"
    testLabelsPath := "t10k-labels-idx1-ubyte"
    testImages, testLabels := MakeInputArray(testImagesPath, testLabelsPath)
    printImages(testImages, testLabels, 10)
    trainImagesPath := "train-images-idx3-ubyte"
    trainLabelsPath := "train-labels-idx1-ubyte"
    images, labels := MakeInputArray(trainImagesPath, trainLabelsPath)
    printImages(images, labels, 10)
}

