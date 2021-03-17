package bmap

import (
	"compress/bzip2"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func testFiles(t *testing.T) {
	inputFile, err := os.Open("../../test/data/input.bz2")
	if err != nil {
		t.Fatalf("open: %v", err)
	}

	inputReader := bzip2.NewReader(inputFile)

	bmapXML, err := ioutil.ReadFile("../../test/data/input.bmap")
	if err != nil {
		t.Fatalf("open: %v", err)
	}

	outputReference, err := ioutil.ReadFile("../../test/data/output.bin")
	if err != nil {
		t.Fatalf("open: %v", err)
	}

	outputFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("tmpfile: %v", err)
	}

	outputName := outputFile.Name()
	defer os.Remove(outputName)

	r, err := NewReader(bmapXML, inputReader)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	err = r.WriteTo(outputFile)
	if err != nil {
		t.Fatalf("WriteTo: %v", err)
	}

	outputFile.Close()

	output, err := ioutil.ReadFile(outputName)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	if !reflect.DeepEqual(output, outputReference) {
		t.Fatalf("DeeepEqual of generated output failed")
	}
}

func TestNew(t *testing.T) {
	t.Run("files", testFiles)
}
