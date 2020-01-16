package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type idxFileFormat struct {
	// structure containing info on the idx file to parse
	dataType   int32
	dataDim    int32
	dims       []int32
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

	dataType := int32(bytes[2]) // Parse the data type and dimension
	dim := int32(bytes[3])

	dims := make([]int32, dim) // make dims array and parse them
	for i := int32(0); i < dim; i++ {
		dims[i] = int32(binary.BigEndian.Uint32(bytes[4+4*i : 8+4*i]))
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

	for n := int32(0); n < header.dims[0]/1000; n++ {
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

func main() {
	data, err := os.Open(os.Args[1]) // open the mnist file given in argument
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}

	imageFile := idxFileFormat{} // create struct and put info in it
	parseHeader(data, &imageFile)
	printHeader(imageFile)
	printData(imageFile)
}
