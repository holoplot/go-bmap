package bmap

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash"
	"io"
	"strconv"
	"strings"
)

// Reader is used to parse bmap xml files and write content to files or block devices
type Reader struct {
	xml         bmapXML
	inputReader io.Reader
}

// NewReader creates a new Bmap reader from the given XML. The inputReader will be used when the content is exported.
func NewReader(bmapXML []byte, inputReader io.Reader) (*Reader, error) {
	r := &Reader{
		inputReader: inputReader,
	}

	err := xml.Unmarshal(bmapXML, &r.xml)
	if err != nil {
		return nil, err
	}

	r.xml.ChecksumType = strings.Trim(r.xml.ChecksumType, " ")

	for _, r := range r.xml.BlockMap.Ranges {
		blocks := strings.Trim(r.Blocks, " ")
		a := strings.Split(blocks, "-")

		r.blockStart, _ = strconv.ParseInt(a[0], 0, 64)
		if len(a) > 1 {
			r.blockEnd, _ = strconv.ParseInt(a[1], 0, 64)
		} else {
			r.blockEnd = r.blockStart
		}
	}

	return r, nil
}

// Size returns the total size of the image, in bytes
func (r *Reader) Size() int64 {
	return r.xml.ImageSize
}

func readBlock(r io.Reader, size int64) ([]byte, error) {
	out := make([]byte, 0)
	pos := int64(0)

	for pos < size {
		b := make([]byte, size-pos)
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		out = append(out, b[:n]...)
		pos += int64(n)
	}

	return out, nil
}

// WriteTo exports the bmap to the given output writer
func (r *Reader) WriteTo(output io.WriterAt) error {
	var h hash.Hash

	switch r.xml.ChecksumType {
	case "sha256":
		h = sha256.New()
	default:
		return fmt.Errorf("Unsupported checksum type %s", r.xml.ChecksumType)
	}

	lastBlock := int64(0)

	for _, rng := range r.xml.BlockMap.Ranges {
		for lastBlock != rng.blockStart {
			_, err := readBlock(r.inputReader, r.xml.BlockSize)
			if err != nil {
				return err
			}

			lastBlock++
		}

		h.Reset()

		for i := rng.blockStart; i <= rng.blockEnd; i++ {
			block, err := readBlock(r.inputReader, r.xml.BlockSize)
			if err != nil {
				return err
			}

			_, err = output.WriteAt(block, i*r.xml.BlockSize)
			if err != nil {
				return err
			}

			h.Write(block)
			lastBlock++
		}

		sum := hex.EncodeToString(h.Sum(nil))
		if sum != rng.Chksum {
			return fmt.Errorf("Checkum mismatch for blocks %d-%d: got %s, expected %s", rng.blockStart, rng.blockEnd, sum, rng.Chksum)
		}
	}

	return nil
}
