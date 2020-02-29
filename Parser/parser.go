package parser

import (
	"encoding/binary"
	"fmt"
	"os"
)

type idxFileFormat struct {
	// structure containing info on the idx file to parse
	dataType   int32 // unsigned int most of the time
	dataDim    int32 // dimension of the data
	dims       []int32 // nb of image; nb of rows; nb of columns
	pixelCount int32
	bytes      []byte
}

func printHeader(file idxFileFormat) {
	// print the info of the file
	fmt.Printf("Data type: %d\nData dim: %d\nPixel count: %d\n%d\n",
		file.dataType, file.dataDim, file.pixelCount, file.dims)
}

func parseHeader(data *os.File, idxFile *idxFileFormat) {
	// parse the idx file data and put info in idxFile struct

	stats, err := data.Stat() // gather the stats on this file
	if err != nil {
		fmt.Println(err)
	}
	var size int64 = stats.Size()

	bytes := make([]byte, size) // make data array and put data in it with data.Read()
	_, err = data.Read(bytes)

// Parse the data type and dimension
	dataType := int32(bytes[2]) // the third byte encode the type of data
    dim := int32(bytes[3]) // fourth byte encode the number of dimensions
	dims := make([]int32, dim) // make dims array and parse them
	for i := int32(0); i < dim; i++ {
        // start at fourth byte and go four by four
		dims[i] = int32(binary.BigEndian.Uint32(bytes[4 + (4 * i) : 8 + (4 * i)]))
	}

	// Put info into struct
	idxFile.dataType = dataType
	idxFile.dataDim = dim
	idxFile.dims = dims
	idxFile.bytes = bytes
	idxFile.pixelCount = dims[0]
	for i := int32(1); i < dim; i++ {
		idxFile.pixelCount *= dims[i]
	}

}

func printData(header idxFileFormat) {
	// Prints the pixels representing the image data in mnist

	if header.dataDim != 3 {
		fmt.Println("Can't print image with idx fil eother than image type (dim  = 3)")
		return
	}

	start := 4 * (header.dataDim) // define where the data starts

	for n := int32(0); n < header.dims[0]; n++ {
		for i := int32(0); i < header.dims[2]; i++ {
			for j := int32(0); j < header.dims[1]; j++ {
				if pixel := int32(header.bytes[n*header.dims[1]*header.dims[2]+
					i*header.dims[2]+j+start+5]); pixel != 0 {
					fmt.Printf("0 ")
				} else {
					fmt.Printf(". ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func MakeInputArray(imagesPath string, labelsPath string) ([][]float64, []float64) {

	dataImage, err := os.Open(imagesPath)
	defer dataImage.Close()
	if err != nil {
		fmt.Println(err)
	}

	dataLabels, err := os.Open(labelsPath)
	defer dataLabels.Close()
	if err != nil {
		fmt.Println(err)
	}

    images := idxFileFormat{}
    labels := idxFileFormat{}
    parseHeader(dataImage, &images)
    parseHeader(dataLabels, &labels)
    imagesArray := make([][]float64, images.dims[0])
    labelsArray := make([]float64, labels.dims[0])
	start := 4 * (images.dataDim)  + 4 // define where the data starts
    imageSize := images.dims[1] * images.dims[2] // imageSize := W * H

	for n := int32(0); n < images.dims[0]; n++ { // go through the images
        imagesArray[n] = make([]float64, imageSize)
        for i := int32(0); i < imageSize; i++ {
            imagesArray[n][i] = float64(images.bytes[n * imageSize + start + i])
        }
        labelsArray[n] = float64(labels.bytes[n + labels.dataDim * 4 + 4])
	}

    return imagesArray, labelsArray
}
