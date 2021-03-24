package main

import "io"

type offsetWriter struct {
	writer io.WriterAt
	skip   int
}

func (sw *offsetWriter) WriteAt(p []byte, off int64) (n int, err error) {
	if int(off) >= sw.skip {
		off -= int64(sw.skip)
	} else if len(p) < sw.skip {
		return 0, nil
	} else {
		off = 0
		p = p[sw.skip:]
	}

	return sw.writer.WriteAt(p, off)
}

func offsetWriterNew(writer io.WriterAt, skipBytes int) *offsetWriter {
	return &offsetWriter{
		writer: writer,
		skip:   skipBytes,
	}
}
