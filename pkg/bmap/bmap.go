package bmap

import (
	"encoding/xml"
)

type bmapXML struct {
	XMLName           xml.Name    `xml:"bmap"`
	Version           string      `xml:"version,attr"`
	ImageSize         int64       `xml:"ImageSize"`
	BlockSize         int64       `xml:"BlockSize"`
	BlockCount        int64       `xml:"BlockCount"`
	MappedBlocksCount int64       `xml:"MappedBlocksCount"`
	ChecksumType      string      `xml:"ChecksumType"`
	BmapFileChecksum  string      `xml:"BmapFileChecksum"`
	BlockMap          blockMapXML `xml:"BlockMap"`
}

type blockMapXML struct {
	XMLName xml.Name    `xml:"BlockMap"`
	Ranges  []*rangeXML `xml:"Range"`
}

type rangeXML struct {
	XMLName xml.Name `xml:"Range"`
	Chksum  string   `xml:"chksum,attr"`
	Blocks  string   `xml:",chardata"`

	// parsed version of the above
	blockStart int64
	blockEnd   int64
}
