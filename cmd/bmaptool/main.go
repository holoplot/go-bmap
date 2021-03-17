package main

import (
	"compress/bzip2"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	bmap "github.com/holoplot/go-bmap/pkg/bmap"
)

func main() {
	bmapFileFlag := flag.String("bmap", "", "Path to bmap XML file (required)")
	inputFileFlag := flag.String("input", "", "Path to input data file (required)")
	outpuFileFlag := flag.String("output", "", "Path to output file or device(required)")
	flag.Parse()

	if len(*bmapFileFlag) == 0 || len(*inputFileFlag) == 0 || len(*outpuFileFlag) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	bmapXML, err := ioutil.ReadFile(*bmapFileFlag)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Open(*inputFileFlag)
	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	var input io.Reader

	if strings.HasSuffix(*inputFileFlag, ".bz2") {
		input = bzip2.NewReader(inputFile)
	} else {
		input = inputFile
	}

	outputFile, err := os.OpenFile(*outpuFileFlag, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	r, err := bmap.NewReader(bmapXML, input)
	if err != nil {
		panic(err)
	}

	err = r.WriteTo(outputFile)
	if err != nil {
		panic(err)
	}

	outputFile.Truncate(r.Size())

	fmt.Printf("Copied %d bytes.\n", r.Size())
}
