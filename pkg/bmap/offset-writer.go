package bmap

import "io"

// OffsetWritder wraps an io.WriterAt and offsets all writes by the given amount
type OffsetWriter struct {
	writer io.WriterAt
	skip   int
}

func (sw *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error) {
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

// OffsetWriterNew returns a new OffsetWriter
func OffsetWriterNew(writer io.WriterAt, skipBytes int) *OffsetWriter {
	return &OffsetWriter{
		writer: writer,
		skip:   skipBytes,
	}
}
